package scan

import (
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/bill"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"time"

	"github.com/anchore/syft/syft/sbom"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/fumeapp/taskin"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type Options struct {
	Json      bool
	File      bool
	FileName  string
	SbomFile  string
	SbomInput string
	Offline   bool
}

func Command() *cobra.Command {
	opts := &Options{
		Json:      false,
		File:      false,
		FileName:  "output.json",
		SbomFile:  "",
		SbomInput: "",
	}

	cmd := &cobra.Command{
		Use:     "scan <path>",
		Short:   i18n.C.ScanShort,
		Example: i18n.C.ScanExample,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.SbomInput == "" && len(args) < 1 {
				return ui.Error(i18n.C.ScanErrorDirectoryRequired)
			}

			var sbm *sbom.SBOM
			var inputRefs []bill.InputSbomRef
			var purls []models.PurlDetail
			var vulns []models.ScanResultVulnerabilities

			var output models.ScanResult

			startTime := time.Now()

			tasks := taskin.Tasks{}

			if opts.SbomInput != "" {
				tasks = append(tasks, taskin.Task{
					Title: fmt.Sprintf("Loading SBOM from %s", opts.SbomInput),
					Task: func(t *taskin.Task) error {
						var err error
						sbm, inputRefs, err = bill.LoadSBOM(opts.SbomInput)

						if err != nil {
							return err
						}
						t.Title = fmt.Sprintf("Loaded SBOM from %s", opts.SbomInput)
						return nil
					},
				})
			} else {
				tasks = append(tasks, taskin.Task{
					Title: i18n.C.ScanSbomStart,
					Task: func(t *taskin.Task) error {
						var err error
						sbm, err = bill.GetSBOM(args[0])
						if err != nil {
							return err
						}
						t.Title = i18n.C.ScanSbomEnd
						return nil
					},
				})
			}

			// Add other necessary tasks after the SBOM task
			tasks = append(tasks, taskin.Tasks{
				{
					Title: i18n.C.ScanExtractPurlStart,
					Task: func(t *taskin.Task) error {
						purls = bill.GetPURLDetail(sbm, inputRefs)
						t.Title = fmt.Sprintf(i18n.C.ScanExtractPurlEnd, len(purls))
						return nil
					},
				},
			}...)

			if opts.Offline {
				tasks = append(tasks, taskin.Tasks{
					{
						Title: i18n.C.ScanScanPurlStartOffline,
						Task: func(t *taskin.Task) error {

							indices, err := cache.Indices()
							if err != nil {
								return err
							}

							vulns = []models.ScanResultVulnerabilities{}
							results, err := bill.GetOfflineVulns(indices, purls, func(cur int, total int) {
								t.Title = fmt.Sprintf(i18n.C.ScanScanPurlProgressOffline, cur, total)
								t.Progress(cur, total)
							})
							if err != nil {
								return err
							}
							vulns = results
							output = models.ScanResult{
								Vulnerabilities: vulns,
							}
							t.Title = fmt.Sprintf(i18n.C.ScanScanPurlEndOffline, len(vulns), len(purls))
							return nil
						},
					},
				}...)
			} else {
				tasks = append(tasks, taskin.Tasks{
					{
						Title: i18n.C.ScanScanPurlStart,
						Task: func(t *taskin.Task) error {
							vulns = []models.ScanResultVulnerabilities{}
							results, err := bill.GetVulns(purls, func(cur int, total int) {
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
							results, err := bill.GetMeta(vulns)
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
				}...)

			}

			if opts.SbomFile != "" {
				tasks = append(tasks, taskin.Task{
					Title: fmt.Sprintf("Saving SBOM to %s", opts.SbomFile),
					Task: func(t *taskin.Task) error {
						if err := bill.SaveSBOM(sbm, opts.SbomFile); err != nil {
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
					if opts.Json {
						ui.Json(output)
						return nil
					} else {
						if err := ui.ScanResults(output.Vulnerabilities); err != nil {
							return err
						}
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

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, i18n.C.FlagOutputJson)
	cmd.Flags().BoolVarP(&opts.File, "file", "f", false, i18n.C.FlagSaveResults)
	cmd.Flags().StringVarP(&opts.FileName, "file-name", "n", "output.json", i18n.C.FlagSpecifyFile)
	cmd.Flags().StringVarP(&opts.SbomFile, "sbom-output-file", "o", "", i18n.C.FlagSpecifySbomFile)
	cmd.Flags().StringVarP(&opts.SbomInput, "sbom-input-file", "i", "", i18n.C.FlagSpecifySbomFile)
	cmd.Flags().BoolVar(&opts.Offline, "offline", false, "Use offline mode to find CVEs - requires indices to be cached")

	return cmd

}
