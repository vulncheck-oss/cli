package bill

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/format"
	"github.com/anchore/syft/syft/format/cyclonedxjson"
	"github.com/anchore/syft/syft/sbom"
	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/client"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeuri"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"github.com/vulncheck-oss/cli/pkg/db"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
)

type InputSbomRef struct {
	SbomRef string
	PURL    string
}

func GetSBOM(dir string) (*sbom.SBOM, error) {
	src, err := syft.GetSource(context.Background(), dir, nil)
	if err != nil {
		return nil, err
	}

	sbm, err := syft.CreateSBOM(context.Background(), src, nil)
	if err != nil {
		return nil, err
	}

	return sbm, nil
}

func SaveSBOM(sbm *sbom.SBOM, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("unable to create file %s: %w", file, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			_ = err
		}
	}()
	encoder, err := cyclonedxjson.NewFormatEncoderWithConfig(cyclonedxjson.DefaultEncoderConfig())
	if err != nil {
		return err
	}

	data, err := format.Encode(*sbm, encoder)
	if err != nil {
		return fmt.Errorf("unable to encode SBOM: %w", err)
	}

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write to file %s: %w", file, err)
	}

	return nil
}

func LoadSBOM(inputFile string) (*sbom.SBOM, []InputSbomRef, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open SBOM file %s: %w", inputFile, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_ = err
		}
	}()

	// Read the entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read SBOM file %s: %w", inputFile, err)
	}

	// Parse JSON to extract bom-ref and purl
	var rawSBOM map[string]interface{}
	err = json.Unmarshal(content, &rawSBOM)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to parse SBOM JSON from file %s: %w", inputFile, err)
	}

	var inputSbomRefs []InputSbomRef

	// Extract bom-ref and purl from components
	if components, ok := rawSBOM["components"].([]interface{}); ok {
		for _, comp := range components {
			if component, ok := comp.(map[string]interface{}); ok {
				bomRef, bomRefOk := component["bom-ref"].(string)
				purl, purlOk := component["purl"].(string)
				if bomRefOk && purlOk {
					inputSbomRefs = append(inputSbomRefs, InputSbomRef{
						SbomRef: bomRef,
						PURL:    purl,
					})
				}
			}
		}
	}

	// Reset file pointer to the beginning for Syft to read
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to reset file pointer: %w", err)
	}

	// Decode SBOM using Syft
	sbm, _, _, err := format.Decode(file)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to decode SBOM from file %s: %w", inputFile, err)
	}

	return sbm, inputSbomRefs, nil
}

func GetCPEDetail(sbm *sbom.SBOM) []string {
	if sbm == nil {
		return []string{}
	}

	var cpes []string
	seen := make(map[string]struct{})

	if sbm.Artifacts.LinuxDistribution != nil && sbm.Artifacts.LinuxDistribution.CPEName != "" {
		cpeStr := strings.TrimSpace(sbm.Artifacts.LinuxDistribution.CPEName)
		norm := cpeutils.NormalizeCPEString(cpeStr)
		if !strings.Contains(cpeStr, ".github/workflows") {
			if _, exists := seen[norm]; !exists {
				cpes = append(cpes, cpeStr)
				seen[norm] = struct{}{}
			}
		}
	}

	for p := range sbm.Artifacts.Packages.Enumerate() {
		if len(p.CPEs) > 0 {
			for _, cpe := range p.CPEs {
				cpeStr := strings.TrimSpace(cpe.Attributes.BindToFmtString())
				norm := cpeutils.NormalizeCPEString(cpeStr)
				if !strings.Contains(cpeStr, ".github/workflows") {
					if _, exists := seen[norm]; !exists {
						cpes = append(cpes, cpeStr)
						seen[norm] = struct{}{}
					}
				}
			}
		}
	}
	return cpes
}

func GetPURLDetail(sbm *sbom.SBOM, inputRefs []InputSbomRef) []models.PurlDetail {
	if sbm == nil {
		return []models.PurlDetail{}
	}

	var purls []models.PurlDetail

	for p := range sbm.Artifacts.Packages.Enumerate() {
		if p.PURL != "" && !strings.HasPrefix(p.PURL, "pkg:github") {
			locations := make([]string, len(p.Locations.ToSlice()))
			for i, l := range p.Locations.ToSlice() {
				locations[i] = l.RealPath
			}

			purlDetail := models.PurlDetail{
				Purl:        p.PURL,
				PackageType: string(p.Type),
				Cataloger:   p.FoundBy,
				Locations:   locations,
			}

			for _, ref := range inputRefs {
				if ref.PURL == p.PURL {
					purlDetail.SbomRef = ref.SbomRef
					break
				}
			}

			purls = append(purls, purlDetail)

		}
	}
	return purls
}

func GetBatchVulns(purls []models.PurlDetail, iterator func(cur int, total int)) ([]models.ScanResultVulnerabilities, error) {
	const batchSize = 100

	var vulns []models.ScanResultVulnerabilities

	purlStrings := make([]string, 0, len(purls))
	for _, purl := range purls {
		purlStrings = append(purlStrings, purl.Purl)
	}

	total := len(purlStrings)

	for start := 0; start < total; start += batchSize {
		end := min(start+batchSize, total)

		batch := purlStrings[start:end]

		response, err := session.Connect(config.Token()).GetPurls(batch)
		if err != nil {
			return nil, fmt.Errorf("error fetching purls %v: %w", batch, err)
		}

		for _, purlResponse := range response.PurlData {
			for _, vuln := range purlResponse.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          purlResponse.PurlMeta.Name,
					Version:       purlResponse.PurlMeta.Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
				})
			}
		}
		iterator(start, total)
	}

	return vulns, nil
}

func GetVulns(purls []models.PurlDetail, iterator func(cur int, total int)) ([]models.ScanResultVulnerabilities, error) {
	var vulns []models.ScanResultVulnerabilities

	i := 0
	for _, purl := range purls {
		i++
		response, err := session.Connect(config.Token()).GetPurl(purl.Purl)
		if err != nil {
			return nil, fmt.Errorf("error fetching purl %s: %v", purl.Purl, err)
		}
		if len(response.Data.Vulnerabilities) > 0 {
			for _, vuln := range response.Data.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          response.PurlMeta().Name,
					Version:       response.PurlMeta().Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
					PurlDetail:    purl,
				})
			}
		}
		iterator(i, len(purls))
	}

	return vulns, nil
}

func GetOfflineCpeVulns(indices cache.InfoFile, cpes []string, iterator func(cur int, total int), warnOnly bool) ([]models.ScanResultVulnerabilities, error) {
	var vulns []models.ScanResultVulnerabilities
	i := 0
	seen := make(map[string]struct{})

	indexAvailable, err := sync.EnsureIndexSync(indices, "cpecve", true)
	if err != nil {
		if warnOnly {
			fmt.Printf("[WARNING]: %s\n", err.Error())
			return nil, nil
		} else {
			return nil, err
		}
	}

	if !indexAvailable {
		if warnOnly {
			fmt.Printf("[WARNING]: index cpecve is required to proceed\n")
			return nil, nil
		} else {
			return nil, fmt.Errorf("index cpecve is required to proceed")
		}
	}

	for _, cpestring := range cpes {
		i++
		cpe, err := cpeuri.ToStruct(cpestring)
		if err != nil {
			return nil, err
		}

		results, _, err := db.CPESearch("cpecve", *cpe)
		if err != nil {
			return nil, err
		}

		cves, err := cpeutils.Process(cpe, results)
		if err != nil {
			return nil, err
		}

		for _, cve := range cves {
			key := cpestring + "|" + cve
			if _, exists := seen[key]; !exists {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:    cpeuri.RemoveSlashes(cpe.Product),
					Version: cpeuri.RemoveSlashes(cpe.Version),
					CVE:     cve,
					CPE:     cpestring,
				})
				seen[key] = struct{}{}
			}
		}
		iterator(i, len(cpes))
	}

	return vulns, nil
}

func GetOfflineVulns(indices cache.InfoFile, purls []models.PurlDetail, iterator func(cur int, total int), warnOnly bool) ([]models.ScanResultVulnerabilities, error) {
	var vulns []models.ScanResultVulnerabilities

	i := 0
	for _, purl := range purls {
		i++
		instance, err := packageurl.FromString(purl.Purl)
		if err != nil {
			return nil, err
		}

		/*
			if packages.IsOS(instance) {
				return nil, fmt.Errorf("operating system package support coming soon")
			}
		*/

		indexName := packages.IndexFromInstance(instance)

		indexAvailable, err := sync.EnsureIndexSync(indices, indexName, true)
		if err != nil {
			if warnOnly {
				fmt.Printf("[WARNING]: %s\n", err.Error())
				continue
			} else {
				return nil, err
			}
		}

		if !indexAvailable {
			if warnOnly {
				fmt.Printf("[WARNING]: index %s is required to PURL %s \n", indexName, purl.Purl)
				continue
			} else {
				return nil, fmt.Errorf("index %s is required to proceed", instance.Type)
			}
		}

		index := indices.GetIndex(indexName)

		results, _, err := db.PURLSearch(index.Name, instance)
		if err != nil {
			return nil, err
		}

		// loop through results and add to vulns
		for _, purlEntry := range results {
			for _, vuln := range purlEntry.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          purlEntry.Name,
					Version:       purlEntry.Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
					PurlDetail:    purl,
				})
			}
		}

		iterator(i, len(purls))
	}

	return vulns, nil
}

func GetMeta(vulns []models.ScanResultVulnerabilities) ([]models.ScanResultVulnerabilities, error) {
	for i, vuln := range vulns {
		nvd2Response, err := session.Connect(config.Token()).GetIndexVulncheckNvd2(sdk.IndexQueryParameters{Cve: vuln.CVE})
		if err != nil {
			return nil, err
		}

		vulns[i].InKEV = nvd2Response.Data[0].VulncheckKEVExploitAdd != nil
		vulns[i].Published = *nvd2Response.Data[0].Published
		vulns[i].CVSSBaseScore = baseScore(nvd2Response.Data[0])
		vulns[i].CVSSTemporalScore = temporalScore(nvd2Response.Data[0])
		vulns[i].Metrics = nvd2Response.Data[0].Metrics
		vulns[i].Weaknesses = nvd2Response.Data[0].Weaknesses

	}
	return vulns, nil
}

func GetOfflineMeta(indices cache.InfoFile, vulns []models.ScanResultVulnerabilities, warnOnly bool) ([]models.ScanResultVulnerabilities, error) {
	indexAvailable, err := sync.EnsureIndexSync(indices, "vulncheck-nvd2", true)
	if err != nil {
		if warnOnly {
			fmt.Printf("[WARNING]: %s\n", err.Error())
			return nil, nil
		} else {
			return nil, err
		}
	}

	if !indexAvailable {
		if warnOnly {
			fmt.Printf("[WARNING]: index vulncheck-nvd2 is required to proceed\n")
			return vulns, nil
		} else {
			return nil, fmt.Errorf("index vulncheck-nvd2 is required to proceed")
		}
	}

	for i, vuln := range vulns {
		nvd2Response, err := db.MetaByCVE(vuln.CVE)
		if err != nil {
			continue
		}

		if len(nvd2Response.Data) > 0 {
			vulns[i].InKEV = nvd2Response.Data[0].VulncheckKEVExploitAdd != nil
			vulns[i].Published = *nvd2Response.Data[0].Published
			vulns[i].CVSSBaseScore = baseScore(nvd2Response.Data[0])
			vulns[i].CVSSTemporalScore = temporalScore(nvd2Response.Data[0])
			vulns[i].Metrics = nvd2Response.Data[0].Metrics
			vulns[i].Weaknesses = nvd2Response.Data[0].Weaknesses
			vulns[i].Description = nvd2Response.Description
		}
	}
	return vulns, nil
}

func baseScore(item client.ApiNVD20CVEExtended) string {
	if item.Metrics == nil {
		return "n/a"
	}
	var score *float32
	if (item.Metrics.CvssMetricV31 != nil) && (len(*item.Metrics.CvssMetricV31) > 0) {
		score = (*item.Metrics.CvssMetricV31)[0].CvssData.BaseScore
	}

	if score == nil && (item.Metrics.CvssMetricV30 != nil) && (len(*item.Metrics.CvssMetricV30) > 0) {
		score = (*item.Metrics.CvssMetricV30)[0].CvssData.BaseScore
	}

	if score == nil && (item.Metrics.CvssMetricV2 != nil) && (len(*item.Metrics.CvssMetricV2) > 0) {
		score = (*item.Metrics.CvssMetricV2)[0].CvssData.BaseScore
	}

	if score == nil {
		return "n/a"
	}

	return formatSingleDecimal(score)
}

func temporalScore(item client.ApiNVD20CVEExtended) string {
	if item.Metrics == nil {
		return "n/a"
	}
	var score *float32

	if item.Metrics.TemporalCVSSV31 != nil {
		score = item.Metrics.TemporalCVSSV31.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV31Secondary != nil && len(*item.Metrics.TemporalCVSSV31Secondary) > 0 {
		score = (*item.Metrics.TemporalCVSSV31Secondary)[0].TemporalScore
	}

	if score == nil && item.Metrics.CvssMetricV30 != nil && len(*item.Metrics.CvssMetricV30) > 0 {
		score = (*item.Metrics.CvssMetricV30)[0].CvssData.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV30Secondary != nil && len(*item.Metrics.TemporalCVSSV30Secondary) > 0 {
		score = (*item.Metrics.TemporalCVSSV30Secondary)[0].TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV30 != nil {
		score = item.Metrics.TemporalCVSSV30.TemporalScore
	}

	if score == nil && item.Metrics.CvssMetricV2 != nil && len(*item.Metrics.CvssMetricV2) > 0 {
		score = (*item.Metrics.CvssMetricV2)[0].CvssData.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV2 != nil {
		score = item.Metrics.TemporalCVSSV2.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV2Secondary != nil && len(*item.Metrics.TemporalCVSSV2Secondary) > 0 {
		score = (*item.Metrics.TemporalCVSSV2Secondary)[0].TemporalScore
	}

	if score == nil {
		return "n/a"
	}

	return formatSingleDecimal(score)
}

func formatSingleDecimal(value *float32) string {
	return fmt.Sprintf("%.1f", *value)
}
