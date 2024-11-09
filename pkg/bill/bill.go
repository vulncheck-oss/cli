package bill

import (
	"context"
	"fmt"
	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/format"
	"github.com/anchore/syft/syft/format/cyclonedxjson"
	"github.com/anchore/syft/syft/sbom"
	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/search"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/sdk-go"
	"github.com/vulncheck-oss/sdk-go/pkg/client"
	"os"
	"strings"
)

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
	defer f.Close()
	encoder, err := cyclonedxjson.NewFormatEncoderWithConfig(cyclonedxjson.DefaultEncoderConfig())
	if err != nil {
		return err
	}

	data, err := format.Encode(*sbm, encoder)
	if err != nil {
		return fmt.Errorf("unable to encode SBOM: %w", err)
	}

	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write to file %s: %w", file, err)
	}

	return nil
}

func LoadSBOM(inputFile string) (*sbom.SBOM, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open SBOM file %s: %w", inputFile, err)
	}
	defer file.Close()

	sbm, _, _, err := format.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("unable to decode SBOM from file %s: %w", inputFile, err)
	}

	return sbm, nil
}

func GetPURLDetail(sbm *sbom.SBOM) []models.PurlDetail {

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
			purls = append(purls, models.PurlDetail{
				Purl:        p.PURL,
				PackageType: string(p.Type),
				Cataloger:   p.FoundBy,
				Locations:   locations,
			})
		}
	}
	return purls
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

func GetOfflineVulns(indices cache.InfoFile, purls []models.PurlDetail, iterator func(cur int, total int)) ([]models.ScanResultVulnerabilities, error) {

	var vulns []models.ScanResultVulnerabilities

	i := 0
	for _, purl := range purls {
		i++
		instance, err := packageurl.FromString(purl.Purl)

		if err != nil {
			return nil, err
		}

		if packages.IsOS(instance) {
			return nil, fmt.Errorf("operating system package support coming soon")
		}

		indexName := packages.IndexFromInstance(instance)

		indexAvailable, err := sync.EnsureIndexSync(indices, indexName, true)
		if err != nil {
			return nil, err
		}

		if !indexAvailable {
			return nil, fmt.Errorf("index %s is required to proceed", instance.Type)
		}

		index := indices.GetIndex(indexName)

		query := search.QueryPURL(instance)

		results, _, err := search.IndexPurl(index.Name, query)
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
		vulns[i].CVSSBaseScore = baseScore(nvd2Response.Data[0])
		vulns[i].CVSSTemporalScore = temporalScore(nvd2Response.Data[0])

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

	if item.Metrics.CvssMetricV31 != nil && len(*item.Metrics.CvssMetricV31) > 0 {
		score = (*item.Metrics.CvssMetricV31)[0].CvssData.TemporalScore
	}

	if score == nil && item.Metrics.TemporalCVSSV31 != nil {
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
