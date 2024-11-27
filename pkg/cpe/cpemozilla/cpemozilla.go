package cpemozilla

import (
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpetypes"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
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
	if !cpeutils.IsParseableVersion(cpetypes.Unquote(cpe.Version)) {
		return nil, fmt.Errorf("failed to parse requested version")
	}

	if product, err := supportedProductsMap()[cpetypes.Unquote(cpe.Product)]; err {
		return &product, nil
	}

	return nil, fmt.Errorf("unable to find mozilla product")
}

func Process(cpe cpetypes.CPE, results []cpetypes.MozillaAdvisory) ([]string, error) {

	var CVEs []string

	for _, hit := range results {

		var err error
		if CVEs, err = processAdvisory(cpe, hit, CVEs); err != nil {
			return nil, fmt.Errorf("failed to process mozilla resutl: %w", err)
		}

	}

	return CVEs, nil
}
func processAdvisory(cpe cpetypes.CPE, advisory cpetypes.MozillaAdvisory, CVEs []string) (cves []string, err error) {

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
			if cmpr, err = cpeutils.CompareVersions(cpetypes.Unquote(cpe.Version), fiver); err == nil && cmpr < 0 {
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

	if ltFixedVersion && len(advisory.AffectedComponents) > 0 {
		for _, affcomp := range advisory.AffectedComponents {
			remainderCVEs, CVEs = processAffectedMozillaComponent(affcomp, "", remainderCVEs, CVEs)
		}
	}

	if ltFixedVersion && len(remainderCVEs) > 0 {
		CVEs = append(CVEs, remainderCVEs...)
	}

	return CVEs, nil

}

func processAffectedMozillaComponent(affcomp cpetypes.MozillaComponent, targetOS string, remainderCVEs, CVEs []string) (rCVEs, cves []string) {
	// remove all CVEs mentioned by affected components. if there
	// are CVEs missed by this component list, then we will add them
	// in after and assume they impact us.
	remainderCVEs = cpeutils.RemoveCVEEntries(remainderCVEs, affcomp.CVE)

	// determine if we have a firefox for windows impacted entry
	// and if so, add it's CVEs
	if targetOS == "windows" && !isFirefoxWindowsAffected(affcomp.Description) {
		return remainderCVEs, CVEs
	}

	return remainderCVEs, append(CVEs, affcomp.CVE...)
}

func isFirefoxWindowsAffected(acdesc string) bool {
	noWin := []string{
		"Windows is not affected",
		"Windows systems are not affected",
		"does not affect Windows",
		"is not present on Windows",
		"issue does not affect Linux or Windows installations",
		"only affects Thunderbird for ",
		"only affects Firefox for Linux",
		"only affects Firefox on Linux",
		"only affects Firefox for macOS and Linux",
		"only affects Firefox on MacOS",
		"only affects Firefox for Android",
		//		"only affects Thunderbird for",
		"only affects OS X and Linux",
		"only affects OS X systems",
		"only affects OS X.",
		"only affects OS X operating systems",
		"only affected Mac OS operating",
		"only affects the Linux oper",
		"only affects Linux insta",
		"only affects Linux users",
		"only affects systems running the Linux operating system",
		"only affected Firefox for Android",
		"only affected Linux and Android operating",
		"only affected Linux operating systems",
		"issue is specific to Linux in",
		"issue only occurs on Linux",
		"issue only affects OS X installations",
		"issue only affects macOS",
		"issue only occurs on Mac OSX.",
		"issue only affects OS X in",
		"issue was limited to a subset of Intel drivers on Linux",
		"acceleration on macOS",
		"problem with Mesa drivers on Linux",
		"reported that on Linux systems,",
		"by other users on Linux and OS X systems.",
		"protocol allowed directory traversal on Linux when",
		"On Linux machines with gnome-vfs",
		"URLs passed to Linux versions of ",
		"Apple has shipped macOS 10.14.5 with an option to disable hyperthreading",
		" reported an Apple issue",
		"ships with Firefox on Mac OS X",
		"passed to Linux versions of Firefox",
		"by other users on Linux and OS X systems",
		"allowed directory traversal on Linux wh",
	}
	yesWin := []string{
		"This issue is specific to Windows",
		"This issue only affects Windows operating systems",
		"ANGLE graphics library is only used on Windows",
		"issue only affects systems running Windows",
		"bugs only affects Windows",
		"issue is limited to the Windows platform",
	}

	// try to normalize some
	acdesc = strings.ReplaceAll(acdesc, "\n", " ")
	acdesc = strings.ReplaceAll(acdesc, "  ", " ")
	acdesc = strings.ReplaceAll(acdesc, " .", ".")

	for _, yw := range yesWin {
		if strings.Contains(acdesc, yw) {
			return true
		}
	}

	for _, nw := range noWin {
		if strings.Contains(acdesc, nw) {
			return false
		}
	}
	return true
}
