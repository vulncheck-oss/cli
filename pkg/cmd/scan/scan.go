package scan

import (
	"context"
	"fmt"
	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/sbom"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/fumeapp/taskin"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
	"github.com/vulncheck-oss/sdk/pkg/client"
	"strings"
)

type Options struct {
	Json     bool
	Annotate bool
}

func Command() *cobra.Command {
	opts := &Options{
		Json:     false,
		Annotate: false,
	}

	cmd := &cobra.Command{
		Use:     "scan <path>",
		Short:   i18n.C.ScanShort,
		Example: i18n.C.ScanExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.ScanErrorDirectoryRequired)
			}

			var sbm *sbom.SBOM
			var purls []string
			var vulns *[]models.ScanResultVulnerabilities

			tasks := taskin.New(taskin.Tasks{
				{
					Title: "Generating SBOM",
					Task: func(t *taskin.Task) error {
						result, err := getSbom(args[0])
						if err != nil {
							return err
						}
						t.Title = "SBOM created"
						sbm = result
						return nil
					},
				},
				{
					Title: "Extracting PURLs",
					Task: func(t *taskin.Task) error {
						purls = getPurls(sbm)
						t.Title = fmt.Sprintf("%d PURLs extracted", len(purls))
						return nil
					},
				},
				{
					Title: "Scanning PURLs for vulnerabilities",
					Task: func(t *taskin.Task) error {
						results, err := getVulns(purls, func(cur int, total int) {
							t.Progress(cur, total)
						})
						if err != nil {
							return err
						}
						vulns = results

						t.Title = fmt.Sprintf("%d vulnerabilities found", len(*vulns))
						return nil
					},
				},
				{
					Title: "Fetching vulnerability metadata",
					Task: func(t *taskin.Task) error {
						results, err := getMeta(*vulns)
						if err != nil {
							return err
						}
						*vulns = results
						t.Title = "Vulnerability metadata fetched"
						return nil
					},
				},
			}, taskin.Config{
				ProgressOptions: []progress.Option{
					progress.WithScaledGradient("#6667AB", "#34D399"),
					progress.WithWidth(20),
				},
			})

			if err := tasks.Run(); err != nil {
				return err
			}

			if len(*vulns) == 0 {
				ui.Info(fmt.Sprintf(i18n.C.ScanNoCvesFound, len(purls)))
				return nil
			}

			result := models.ScanResult{
				Vulnerabilities: *vulns,
			}

			if err := ui.ScanResults(result.Vulnerabilities); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	cmd.Flags().BoolVarP(&opts.Annotate, "annotate", "a", false, "Output as Github Annotations")

	return cmd

}

func getSbom(dir string) (*sbom.SBOM, error) {
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

func getPurls(sbm *sbom.SBOM) []string {

	var purls []string

	for p := range sbm.Artifacts.Packages.Enumerate() {
		if p.PURL != "" && !strings.HasPrefix(p.PURL, "pkg:github") {
			purls = append(purls, p.PURL)
		}
	}

	return purls
}

func getVulns(purls []string, iterator func(cur int, total int)) (*[]models.ScanResultVulnerabilities, error) {

	var vulns []models.ScanResultVulnerabilities

	i := 0
	for _, purl := range purls {
		i++
		response, err := session.Connect(config.Token()).GetPurl(purl)
		if err != nil {
			return nil, fmt.Errorf("error fetching purl %s: %v", purl, err)
		}
		if len(response.Data.Vulnerabilities) > 0 {
			for _, vuln := range response.Data.Vulnerabilities {
				vulns = append(vulns, models.ScanResultVulnerabilities{
					Name:          response.PurlMeta().Name,
					Version:       response.PurlMeta().Version,
					CVE:           vuln.Detection,
					FixedVersions: vuln.FixedVersion,
				})

			}
		}
		iterator(i, len(purls))
	}

	return &vulns, nil
}

func getMeta(vulns []models.ScanResultVulnerabilities) ([]models.ScanResultVulnerabilities, error) {
	for i, vuln := range vulns {
		nvd2Response, err := session.Connect(config.Token()).GetIndexVulncheckNvd2(sdk.IndexQueryParameters{Cve: vuln.CVE})

		if err != nil {
			return nil, err
		}

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
