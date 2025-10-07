package scan

import (
	"fmt"
	"time"

	"github.com/vulncheck-oss/cli/pkg/bill"
	"github.com/vulncheck-oss/cli/pkg/cache"

	"github.com/anchore/syft/syft/sbom"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/fumeapp/taskin"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type Options struct {
	Json        bool
	File        bool
	FileName    string
	SbomFile    string
	SbomInput   string
	SbomOnly    bool
	Cpes        bool
	Offline     bool
	OfflineMeta bool
	DisableUI   bool
	WarnOnIndex bool
}

func Command() *cobra.Command {
	opts := &Options{
		Json:        false,
		File:        false,
		FileName:    "output.json",
		SbomFile:    "",
		SbomInput:   "",
		SbomOnly:    false,
		Cpes:        false,
		DisableUI:   false,
		WarnOnIndex: false,
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
			var cpes []string
			var cpeVulns []models.ScanResultVulnerabilities
			var purlVulns []models.ScanResultVulnerabilities
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

			if !opts.SbomOnly {
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

				if !opts.Offline && opts.Cpes {
					return ui.Error("CPE extraction/scanning for online mode is coming soon")
				}

				if opts.Cpes {
					tasks = append(tasks, taskin.Tasks{
						{
							Title: i18n.C.ScanExtractCpeStart,
							Task: func(t *taskin.Task) error {
								cpes = bill.GetCPEDetail(sbm)
								t.Title = fmt.Sprintf(i18n.C.ScanExtractCpeEnd, len(cpes))
								return nil
							},
						},
					}...)
				}

				if opts.Offline {
					if opts.Cpes {
						tasks = append(tasks, taskin.Tasks{
							{
								Title: i18n.C.ScanScanCpeStartOffline,
								Task: func(t *taskin.Task) error {

									indices, err := cache.Indices()
									if err != nil {
										return err
									}

									results, err := bill.GetOfflineCpeVulns(indices, cpes, func(cur int, total int) {
										t.Title = fmt.Sprintf(i18n.C.ScanScanCpeProgressOffline, cur, total)
										t.Progress(cur, total)
									}, opts.WarnOnIndex)
									if err != nil {
										return err
									}
									cpeVulns = results
									t.Title = fmt.Sprintf(i18n.C.ScanScanCpeEndOffline, len(cpeVulns), len(cpes))
									return nil
								},
							},
						}...)
					}
					tasks = append(tasks, taskin.Tasks{
						{
							Title: i18n.C.ScanScanPurlStartOffline,
							Task: func(t *taskin.Task) error {

								indices, err := cache.Indices()
								if err != nil {
									return err
								}

								purlVulns = []models.ScanResultVulnerabilities{}
								results, err := bill.GetOfflineVulns(indices, purls, func(cur int, total int) {
									t.Title = fmt.Sprintf(i18n.C.ScanScanPurlProgressOffline, cur, total)
									t.Progress(cur, total)
								}, opts.WarnOnIndex)
								if err != nil {
									return err
								}
								purlVulns = results
								t.Title = fmt.Sprintf(i18n.C.ScanScanPurlEndOffline, len(purlVulns), len(purls))
								// we need to mege vulns and cpeVulns here
								vulns = append(cpeVulns, purlVulns...)
								output = models.ScanResult{
									Vulnerabilities: vulns,
								}
								return nil
							},
						},
					}...)
					/*
						1. check if the vulncheck-nvd2 index is cached
						2. populate vulns with metadata
					*/
					if opts.OfflineMeta {

						tasks = append(tasks, taskin.Tasks{
							{
								Title: i18n.C.ScanVulnOfflineMetaStart,
								Task: func(t *taskin.Task) error {
									indices, _ := cache.Indices()
									results, err := bill.GetOfflineMeta(indices, vulns, opts.WarnOnIndex)
									if err != nil {
										return err
									}
									vulns = results
									t.Title = i18n.C.ScanVulnOfflineMetaEnd
									output = models.ScanResult{
										Vulnerabilities: vulns,
									}
									return nil
								},
							},
						}...)

					}
				} else {
					tasks = append(tasks, taskin.Tasks{
						{
							Title: i18n.C.ScanScanPurlStart,
							Task: func(t *taskin.Task) error {
								purlVulns = []models.ScanResultVulnerabilities{}
								results, err := bill.GetVulns(purls, func(cur int, total int) {
									t.Title = fmt.Sprintf(i18n.C.ScanScanPurlProgress, cur, total)
									t.Progress(cur, total)
								})
								if err != nil {
									return err
								}
								purlVulns = results
								t.Title = fmt.Sprintf(i18n.C.ScanScanPurlEnd, len(purlVulns), len(purls))
								// we will combine cpeVUlns when cpe online scanning is available
								vulns = purlVulns
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

			if !opts.SbomOnly && opts.File {
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
				DisableUI: opts.DisableUI,
				ProgressOptions: []progress.Option{
					progress.WithScaledGradient("#6667AB", "#34D399"),
					progress.WithWidth(20),
					progress.WithoutPercentage(),
				},
			})

			if err := runners.Run(); err != nil {
				return err
			}

			// Only display scan results if we're not in SbomOnly mode
			if !opts.SbomOnly {
				if vulns != nil {
					if len(vulns) == 0 {
						ui.Info(fmt.Sprintf(i18n.C.ScanNoCvesFound, len(purls)))
					}
					if len(vulns) > 0 {
						if opts.Json {
							ui.Json(output)
							return nil
						} else {
							if err := ui.ScanResults(output.Vulnerabilities, opts.Offline && !opts.OfflineMeta); err != nil {
								return err
							}
						}
					}
				} else {
					ui.Info(fmt.Sprintf(i18n.C.ScanNoCvesFound, len(purls)))
				}

				elapsedTime := time.Since(startTime)
				ui.Info(fmt.Sprintf(i18n.C.ScanBenchmark, elapsedTime))
			} else if opts.SbomFile != "" {
				ui.Info("SBOM generation completed successfully")
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, i18n.C.FlagOutputJson)
	cmd.Flags().BoolVarP(&opts.File, "file", "f", false, i18n.C.FlagSaveResults)
	cmd.Flags().StringVarP(&opts.FileName, "file-name", "n", "output.json", i18n.C.FlagSpecifyFile)
	cmd.Flags().StringVarP(&opts.SbomFile, "sbom-output-file", "o", "", i18n.C.FlagSpecifySbomFile)
	cmd.Flags().StringVarP(&opts.SbomInput, "sbom-input-file", "i", "", i18n.C.FlagSpecifySbomInput)
	cmd.Flags().BoolVarP(&opts.SbomOnly, "sbom-only", "s", false, i18n.C.FlagSpecifySbomOnly)
	cmd.Flags().BoolVarP(&opts.Cpes, "include-cpes", "c", false, i18n.C.FlagIncludeCpes)
	cmd.Flags().BoolVar(&opts.Offline, "offline", false, "Use offline mode to find CVEs - requires indices to be cached")
	cmd.Flags().BoolVar(&opts.OfflineMeta, "offline-meta", false, "Use with offline mode to populate CVE metadata - requires the vulncheck-nvd2 index to be cached")
	cmd.Flags().BoolVar(&opts.WarnOnIndex, "warn-on-index", false, "When an index is not present locally, show a warning instead of shutting down")
	cmd.Flags().BoolVar(&opts.DisableUI, "disable-ui", false, "Disable interactive UI elements")

	return cmd
}
