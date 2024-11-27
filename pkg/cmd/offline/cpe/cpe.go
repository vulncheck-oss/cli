package cpe

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeparse"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeprocess"
	"github.com/vulncheck-oss/cli/pkg/search"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {

	var jsonOutput bool

	cmd := &cobra.Command{
		Use:     "cpe <scheme>",
		Short:   "Offline CPE lookup",
		Long:    "Search offline package data via CPE schemes",
		Example: "vulncheck offline cpe \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cpe, err := cpeparse.Parse(args[0])

			if err != nil {
				return err
			}

			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			indexAvailable, err := sync.EnsureIndexSync(indices, cpe.Vendor, false)
			if err != nil {
				return err
			}

			if !indexAvailable {
				return fmt.Errorf("index %s is required to proceed", cpe.Vendor)
			}

			query := search.QueryCPE(*cpe)

			results, stats, err := search.IndexAdvisories(cpe.Vendor, query)

			if err != nil {
				return err
			}
			cves, err := cpeprocess.Process(*cpe, results)

			if err != nil {
				return err
			}

			ui.Stat("Results found/filtered", fmt.Sprintf("%d/%d", len(results), len(cves)))
			ui.Stat("Files/Lines processed", fmt.Sprintf("%d/%d", stats.TotalFiles, stats.TotalLines))
			ui.Stat("Search duration", stats.Duration.String())

			ui.Json(cves)

			return nil
		},
	}
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON format")
	return cmd
}
