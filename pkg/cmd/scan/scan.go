package scan

import (
	"fmt"
	"time"

	"github.com/vulncheck-oss/cli/internal/output"
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
		FileName: "output.json",
	}

	cmd := &cobra.Command{
		Use:     "scan <path>",
		Short:   i18n.C.ScanShort,
		Example: i18n.C.ScanExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			ctx := cmd.Context()

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
			// metaAvailable tracks whether the vulncheck-nvd2 index was
			// usable. When --offline-meta is requested but the index is
			// missing (with --warn-on-index), we still surface the CVEs
			// we found - just without empty score columns - and tell the
			// user how to populate them.
			metaAvailable := true

			var result models.ScanResult

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
								cpes = bill.GetCPEDetail(sbm, inputRefs)
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
								vulns = append(cpeVulns, purlVulns...)
								result = models.ScanResult{
									Vulnerabilities: vulns,
								}
								return nil
							},
						},
					}...)
					if opts.OfflineMeta {
						tasks = append(tasks, taskin.Tasks{
							{
								Title: i18n.C.ScanVulnOfflineMetaStart,
								Task: func(t *taskin.Task) error {
									indices, _ := cache.Indices()
									results, ok, err := bill.GetOfflineMeta(indices, vulns, opts.WarnOnIndex)
									if err != nil {
										return err
									}
									vulns = results
									metaAvailable = ok
									if ok {
										t.Title = i18n.C.ScanVulnOfflineMetaEnd
									} else {
										t.Title = i18n.C.ScanVulnOfflineMetaUnavailable
									}
									result = models.ScanResult{
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
								results, err := bill.GetBatchVulns(ctx, purls, func(cur int, total int) {
									t.Title = fmt.Sprintf(i18n.C.ScanScanPurlProgress, cur, total)
									t.Progress(cur, total)
								})
								if err != nil {
									return err
								}
								purlVulns = results
								t.Title = fmt.Sprintf(i18n.C.ScanScanPurlEnd, len(purlVulns), len(purls))
								vulns = purlVulns
								return nil
							},
						},
						{
							Title: i18n.C.ScanVulnMetaStart,
							Task: func(t *taskin.Task) error {
								results, err := bill.GetMeta(ctx, vulns)
								if err != nil {
									return err
								}
								vulns = results
								t.Title = i18n.C.ScanVulnMetaEnd
								result = models.ScanResult{
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
						if err := ui.JsonFile(result, opts.FileName); err != nil {
							return err
						}
						t.Title = fmt.Sprintf("Results saved to %s", opts.FileName)
						return nil
					},
				})
			}

			// Progress UI is suppressed in JSON mode (would corrupt stdout)
			// and whenever the caller explicitly asks for --disable-ui.
			// Headless detection (non-TTY, NO_COLOR, CI) lands in phase 4.
			disableUI := opts.DisableUI || r.IsJSON()

			runners := taskin.New(tasks, taskin.Config{
				DisableUI: disableUI,
				ProgressOptions: []progress.Option{
					progress.WithScaledGradient("#6667AB", "#34D399"),
					progress.WithWidth(20),
					progress.WithoutPercentage(),
				},
			})

			if err := runners.Run(); err != nil {
				return err
			}

			if opts.SbomOnly {
				if opts.SbomFile != "" {
					r.Info("SBOM generation completed successfully")
				}
				return nil
			}

			if r.IsJSON() {
				return r.JSON(result)
			}

			if len(vulns) == 0 {
				r.Info(i18n.C.ScanNoCvesFound, len(purls))
			} else {
				// Hide score columns whenever we don't have nvd2 metadata
				// (either explicitly skipped, or the index wasn't cached).
				hideScores := opts.Offline && (!opts.OfflineMeta || !metaAvailable)
				if err := ui.ScanResults(result.Vulnerabilities, hideScores); err != nil {
					return err
				}
				if opts.OfflineMeta && !metaAvailable {
					r.Info("%s", i18n.C.ScanVulnOfflineMetaUnavailable)
				}
			}

			r.Info(i18n.C.ScanBenchmark, time.Since(startTime))
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.File, "file", "f", false, i18n.C.FlagSaveResults)
	cmd.Flags().StringVarP(&opts.FileName, "file-name", "n", "output.json", i18n.C.FlagSpecifyFile)
	cmd.Flags().StringVarP(&opts.SbomFile, "sbom-output-file", "o", "", i18n.C.FlagSpecifySbomFile)
	cmd.Flags().StringVarP(&opts.SbomInput, "sbom-input-file", "i", "", i18n.C.FlagSpecifySbomInput)
	cmd.Flags().BoolVarP(&opts.SbomOnly, "sbom-only", "s", false, i18n.C.FlagSpecifySbomOnly)
	cmd.Flags().BoolVarP(&opts.Cpes, "include-cpes", "c", false, i18n.C.FlagIncludeCpes)
	cmd.Flags().BoolVar(&opts.Offline, "offline", false, "Use offline mode to find CVEs - requires indices to be cached")
	cmd.Flags().BoolVar(&opts.OfflineMeta, "offline-meta", false, "Use with offline mode to populate CVE metadata - requires the vulncheck-nvd2 index to be cached")
	cmd.Flags().BoolVar(&opts.WarnOnIndex, "warn-on-index", false, "When an index is not present locally, show a warning instead of shutting down")
	cmd.Flags().BoolVar(&opts.DisableUI, "disable-ui", false, "Disable interactive UI elements (progress bars, spinners)")

	return cmd
}
