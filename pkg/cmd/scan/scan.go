package scan

import (
	"context"
	"fmt"
	"github.com/anchore/syft/syft"
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

			src, err := syft.GetSource(context.Background(), args[0], nil)

			if err != nil {
				return err
			}

			sbom, err := syft.CreateSBOM(context.Background(), src, nil)

			if err != nil {
				return err
			}

			var purls []string

			for p := range sbom.Artifacts.Packages.Enumerate() {
				if !strings.HasPrefix(p.PURL, "pkg:github") {
					purls = append(purls, p.PURL)
				}
			}

			if !opts.Json && !opts.Annotate {
				ui.Info(fmt.Sprintf(i18n.C.ScanPackagesFound, len(purls)))
				ui.NewProgress(len(purls))
			}

			var vulns []models.ScanResultVulnerabilities

			for index, purl := range purls {
				response, err := session.Connect(config.Token()).GetPurl(purl)
				if err != nil {
					return fmt.Errorf("error fetching purl %s: %w", purl, err)
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
				if !opts.Json && !opts.Annotate {
					ui.UpdateProgress(index + 1)
				}
			}

			result := models.ScanResult{
				Vulnerabilities: []models.ScanResultVulnerabilities{},
			}

			for _, vuln := range vulns {
				nvd2Response, err := session.Connect(config.Token()).GetIndexVulncheckNvd2(sdk.IndexQueryParameters{Cve: vuln.CVE})
				result.Vulnerabilities = append(result.Vulnerabilities, models.ScanResultVulnerabilities{
					CVE:               vuln.CVE,
					Name:              vuln.Name,
					Version:           vuln.Version,
					CVSSBaseScore:     baseScore(nvd2Response.Data[0]),
					CVSSTemporalScore: temporalScore(nvd2Response.Data[0]),
					FixedVersions:     vuln.FixedVersions,
				})

				if err != nil {
					return err
				}
			}

			if opts.Json {
				ui.Json(result)
				return nil
			}

			if len(vulns) == 0 {
				ui.Info(fmt.Sprintf(i18n.C.ScanNoCvesFound, len(purls)))
				return nil
			}

			ui.Info(fmt.Sprintf(i18n.C.ScanCvesFound, len(vulns), len(purls)))

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
