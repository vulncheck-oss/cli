package cpe

import (
	"fmt"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeoffline"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeuri"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"github.com/vulncheck-oss/cli/pkg/search"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {

	var jsonOutput bool

	var statsOnly bool

	cmd := &cobra.Command{
		Use:     "cpe <scheme>",
		Short:   "Offline CPE lookup",
		Long:    "Search offline package data via CPE schemes",
		Example: "vulncheck offline cpe \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cpe, err := cpeuri.ToStruct(args[0])

			if err != nil {
				return err
			}

			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			indexAvailable, err := sync.EnsureIndexSync(indices, "cpecve", false)
			if err != nil {
				return err
			}

			if !indexAvailable {
				return fmt.Errorf("index cpecve is required to proceed")
			}

			query, err := cpeoffline.Query(cpe)

			ray.Ray(query)

			if err != nil {
				return err
			}

			results, stats, err := search.IndexCPE("cpecve", *cpe, query)

			if err != nil {
				return err
			}
			cves, err := cpeutils.Process(cpe, results)

			if err != nil {
				return err
			}

			if jsonOutput || config.IsCI() {
				ui.Json(cves)
				return nil
			}

			ui.Stat("Results found/filtered", fmt.Sprintf("%d/%d", len(results), len(cves)))
			ui.Stat("Files/Lines processed", fmt.Sprintf("%d/%d", stats.TotalFiles, stats.TotalLines))
			ui.Stat("Search duration", fmt.Sprintf("%.2f seconds", stats.Duration.Seconds()))

			if !statsOnly {
				ui.Json(cves)
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON format")
	cmd.Flags().BoolVarP(&statsOnly, "stats", "s", false, "Output stats only")
	return cmd
}
