package scan

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/format"
	"github.com/anchore/syft/syft/format/cyclonedxjson"
	"github.com/anchore/syft/syft/sbom"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/fumeapp/taskin"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk-go"
	"github.com/vulncheck-oss/sdk-go/pkg/client"
)

type Options struct {
	File      bool
	FileName  string
	SbomFile  string
	SbomInput string
}

func Command() *cobra.Command {
	opts := &Options{
		File:     false,
		FileName: "output.json",
		SbomFile: "",
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
			var purls []models.PurlDetail
			var vulns []models.ScanResultVulnerabilities

			var output models.ScanResult

			startTime := time.Now()

			tasks := taskin.Tasks{
				{
					Title: i18n.C.ScanSbomStart,
					Task: func(t *taskin.Task) error {
						result, err := getSbom(args[0])
						if err != nil {
							return err
						}
						t.Title = i18n.C.ScanSbomEnd
						sbm = result
						return nil
					},
				},
				{
					Title: i18n.C.ScanExtractPurlStart,
					Task: func(t *taskin.Task) error {
						purls = getPurls(sbm)
						t.Title = fmt.Sprintf(i18n.C.ScanExtractPurlEnd, len(purls))
						return nil
					},
				},
				{
					Title: i18n.C.ScanScanPurlStart,
					Task: func(t *taskin.Task) error {
						vulns = []models.ScanResultVulnerabilities{}
						results, err := getVulns(purls, func(cur int, total int) {
							t.Title = fmt.Sprintf(i18n.C.ScanScanPurlProgress, cur, total)
							t.Progress(cur, total)
						})
						if err != nil {
							return err
						}
						vulns = results
						t.Title = fmt.Sprintf(i18n.C.ScanScanPurlEnd, len(vulns), len(purls))
						return nil
					},
				},
				{
					Title: i18n.C.ScanVulnMetaStart,
					Task: func(t *taskin.Task) error {
						results, err := getMeta(vulns)
						if err != nil {
							return err
						}
						vulns = results
						t.Title = i18n.C.ScanVulnMetaEnd
						output = models.ScanResult{
							Vulnerabilities: vulns,
						}
						return nil
					},
				},
			}

			if opts.SbomFile != "" {
				tasks = append(tasks, taskin.Task{
					Title: fmt.Sprintf("Saving SBOM to %s", opts.SbomFile),
					Task: func(t *taskin.Task) error {
						if err := saveSbom(sbm, opts.SbomFile); err != nil {
							return err
						}
						t.Title = fmt.Sprintf("SBOM saved to %s", opts.SbomFile)
						return nil
					},
				})
			}

			if opts.File {
				tasks = append(tasks, taskin.Task{
					Title: fmt.Sprintf("Saving results to %s", opts.FileName),
					Task: func(t *taskin.Task) error {
						if err := ui.JsonFile(output, opts.FileName); err != nil {
							return err
						}
						t.Title = fmt.Sprintf("Results saved to %s", opts.FileName)
						return nil
					},
				})
			}

			runners := taskin.New(tasks, taskin.Config{
				ProgressOptions: []progress.Option{
					progress.WithScaledGradient("#6667AB", "#34D399"),
					progress.WithWidth(20),
					progress.WithoutPercentage(),
				},
			})

			if err := runners.Run(); err != nil {
				return err
			}

			if vulns != nil {
				if len(vulns) == 0 {
					ui.Info(fmt.Sprintf(i18n.C.ScanNoCvesFound, len(purls)))
				}
				if len(vulns) > 0 {
					if err := ui.ScanResults(output.Vulnerabilities); err != nil {
						return err
					}
				}
			} else {
				ui.Info(fmt.Sprintf(i18n.C.ScanNoCvesFound, len(purls)))
			}

			elapsedTime := time.Since(startTime)

			ui.Info(fmt.Sprintf(i18n.C.ScanBenchmark, elapsedTime))

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.File, "file", "f", false, i18n.C.FlagSaveResults)
	cmd.Flags().StringVarP(&opts.FileName, "file-name", "n", "output.json", i18n.C.FlagSpecifyFile)
	cmd.Flags().StringVarP(&opts.SbomFile, "sbom-output-file", "so", "", i18n.C.FlagSpecifySbomFile)
	cmd.Flags().StringVarP(&opts.SbomInput, "sbom-input-file", "si", "", i18n.C.FlagSpecifySbomFile)

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

// saveSbom saves the SBOM to a specified file
func saveSbom(sbm *sbom.SBOM, file string) error {

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

func getPurls(sbm *sbom.SBOM) []models.PurlDetail {

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

func getVulns(purls []models.PurlDetail, iterator func(cur int, total int)) ([]models.ScanResultVulnerabilities, error) {

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

func getMeta(vulns []models.ScanResultVulnerabilities) ([]models.ScanResultVulnerabilities, error) {
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
