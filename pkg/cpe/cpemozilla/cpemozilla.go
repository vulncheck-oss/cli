package cpemozilla

import (
	"fmt"
	hcversion "github.com/hashicorp/go-version"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
	"github.com/vulncheck-oss/cli/pkg/search"
	"regexp"
	"strconv"
	"strings"
)

func supportedProductsMap() map[string]string {
	return map[string]string{
		"firefox":             "firefox",
		"firefox esr":         "firefox esr",
		"firefox os":          "firefox os",
		"firefox adsense":     "firefox adsense",
		"firefox mobile":      "firefox mobile",
		"firefox for ios":     "firefox for ios",
		"firefox for android": "firefox for android",
		"firefox_esr":         "firefox esr",
		"firefox_os":          "firefox os",
		"firefox_adsense":     "firefox adsense",
		"firefox_mobile":      "firefox mobile",
		"firefox_for_ios":     "firefox for ios",
		"firefox_for_android": "firefox for android",
		"thunderbird":         "thunderbird",
		"thunderbird_esr":     "thunderbird esr",
		// XXX: need to support product=X os=Y, say firefox and android => Firefox for Android
	}
}

func Parse(cpe cpetypes.CPE) (*string, error) {
	if !isParseableVersion(cpetypes.Unquote(cpe.Version)) {
		return nil, fmt.Errorf("failed to parse requested version")
	}

	if product, err := supportedProductsMap()[cpetypes.Unquote(cpe.Product)]; err {
		return &product, nil
	}

	return nil, fmt.Errorf("unable to find mozilla product")
}

func Process(cpe cpetypes.CPE, results search.AdvisoryEntries) ([]string, error) {

	var CVEs []string

	for _, hit := range results {

		var err error
		if CVEs, err = processAdvisory(cpe, hit, CVEs); err != nil {
			return nil, fmt.Errorf("failed to process mozilla resutl: %w", err)
		}

	}

	return CVEs, nil
}
func processAdvisory(cpe cpetypes.CPE, advisory search.AdvisoryEntry, CVEs []string) (cves []string, err error) {

	isSameProduct := false
	var hitprod string

	for _, hitprod = range advisory.Products {
		hitprod = strings.ToLower(hitprod)
		if hitprod == cpe.Product {
			isSameProduct = true
			break
		}
	}

	if !isSameProduct {
		return CVEs, nil
	}

	ltFixedVersion := false

	var remainderCVEs []string

	for _, fiv := range advisory.FixedIn {

		fiv = strings.ToLower(fiv)
		if !strings.HasPrefix(fiv, cpe.Product) {
			continue
		}

		// handle firefox esr vs firefox, etc
		if (strings.Contains(fiv, " esr") && !strings.Contains(cpe.Product, " esr")) ||
			(strings.Contains(fiv, " os") && !strings.Contains(cpe.Product, " os")) ||
			(strings.Contains(fiv, " adsense") && !strings.Contains(cpe.Product, " adsense")) ||
			(strings.Contains(fiv, " mobile") && !strings.Contains(cpe.Product, " mobile")) ||
			(strings.Contains(fiv, " for ios") && !strings.Contains(cpe.Product, " for ios")) ||
			(strings.Contains(fiv, " for android") && !strings.Contains(cpe.Product, " for android")) {
			continue
		}

		fisplit := strings.Split(fiv, cpe.Product)
		if len(fisplit) == 2 {
			fiver := strings.TrimSpace(fisplit[1])

			var cmpr int
			if cmpr, err = CompareVersions(cpetypes.Unquote(cpe.Version), fiver); err == nil && cmpr < 0 {
				//fmt.Printf("product: %v cmpre: %v, versionUnquote: %v, fiver: %v, CVE: %v Title: %v\n", hitprod, cmpr, versionUnquote, fiver, mozillaAdvisory.CVE, mozillaAdvisory.Title)
				if len(advisory.AffectedComponents) == 0 {
					CVEs = append(CVEs, advisory.CVE...)
					remainderCVEs = remainderCVEs[0:0]
				}
				ltFixedVersion = true
				break
			} else if err != nil {
				fmt.Printf("GetAssessFromMozillaByProductAndVersionUnioned: version compare issue: input=%v mozilla-data=%v err=%v\n", cpetypes.Unquote(cpe.Version), fiver, err)
			}
		} else {
			fmt.Printf("GetAssessFromMozillaByProductAndVersionUnioned: fixed_in split != 2: fiv: %v, sprod: %v\n", fiv, cpetypes.Unquote(cpe.Version))
		}

	}

	if ltFixedVersion && len(remainderCVEs) > 0 {
		CVEs = append(CVEs, remainderCVEs...)
	}

	return CVEs, nil

}

func isParseableVersion(a string) bool {
	if _, err := GetUint64FromVersion(a); err == nil {
		return true
	}
	_, err := hcversion.NewVersion(a)
	return err == nil
}

// GetUint64FromVersion attempts to map period separated numeric version
// string to a uint64. Needs testing.
func GetUint64FromVersion(version string) (uint64, error) {
	iv, err := versionToUint64Slice(version)
	if err != nil {
		return 0, err
	}
	return uint64SliceToUint64(iv)
}

func uint64SliceToUint64(isl []uint64) (uint64, error) {
	if len(isl) == 0 {
		return 0, fmt.Errorf("no fields detected in version slice")
	}
	if len(isl) > 8 {
		return 0, fmt.Errorf("too many fields detected in version slice")
	}
	uver := uint64(0)
	mask := uint64(0xff00000000000000)
	for i, donor := range isl {
		uver |= (donor << (56 - (i * 8))) & mask
		mask = mask >> 8
	}
	return uver, nil
}

func versionToUint64Slice(version string) ([]uint64, error) {
	s := strings.Split(version, ".")
	if len(s) > 8 {
		return nil, fmt.Errorf("version %s contains greater than 8 fields", version)
	}

	ivals := make([]uint64, len(s))
	for i, v := range s {
		ival, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("version %s contains unsupported field: %v", version, err)
		}
		ivals[i] = ival
	}
	return ivals, nil
}

// Attempt to compare versions
func CompareVersions(a string, b string) (int, error) {
	var c int
	var err error

	// try to handle the e.g., 5.x case
	m, err := regexp.Match(`^\d{1,3}\.x$`, []byte(b))
	if err == nil && m {
		amaj := strings.Split(a, ".")[0]
		amajv, errA := strconv.Atoi(amaj)
		bmaj := strings.Split(b, ".")[0]
		bmajv, errB := strconv.Atoi(bmaj)
		if errA == nil && errB == nil {
			if amajv < bmajv {
				return -1, nil
			} else if amajv == bmajv {
				return 0, nil
			} else {
				return 1, nil
			}
		}
	}

	c, err = CompareVersionsByUint64(a, b)
	if err == nil {
		return c, nil
	}

	hcvA, err := hcversion.NewVersion(a)
	if err != nil {
		return 0, fmt.Errorf("unable to parse version: %s", a)
	}
	hcvB, err := hcversion.NewVersion(b)
	if err != nil {
		return 0, fmt.Errorf("unable to parse version: %s", b)
	}

	if hcvA.LessThan(hcvB) {
		return -1, nil
	} else if hcvA.Equal(hcvB) {
		return 0, nil
	}
	return 1, nil
}

func CompareVersionsByUint64(a, b string) (int, error) {
	au64, err1 := GetUint64FromVersion(a)
	bu64, err2 := GetUint64FromVersion(b)

	if err1 != nil {
		return 0, err1
	}
	if err2 != nil {
		return 0, err2
	}
	if au64 < bu64 {
		return -1, nil
	}
	if au64 == bu64 {
		return 0, nil
	}
	return 1, nil
}
