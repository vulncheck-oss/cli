package purl

import (
	"fmt"
	"github.com/package-url/packageurl-go"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/ipintel"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/search"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func Command() *cobra.Command {

	var jsonOutput bool

	cmd := &cobra.Command{
		Use:     "purl <scheme>",
		Short:   "Offline PURL lookup",
		Long:    "Search offline package data via PURL schemes",
		Example: "vulncheck offline purl \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			instance, err := packageurl.FromString(args[0])

			if err != nil {
				return err
			}

			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			index := indices.GetIndex(instance.Type)
			if index == nil {
				return fmt.Errorf("index %s is required for this command, and is not cached", instance.Type)
			}

			ui.PurlInstance(instance)

			query := ipintel.BuildPurlQuery(instance)

			if !jsonOutput && !config.IsCI() {
				ui.Info(fmt.Sprintf("Searching index %s, last updated on %s", index.Name, utils.ParseDate(index.LastUpdated)))
			}

			results, stats, err := search.IndexPurl(index.Name, query)
			if err != nil {
				return err
			}

			if jsonOutput || config.IsCI() {
				ui.Json(results)
				return nil
			}

			ui.Stat("Results found", fmt.Sprintf("%d", len(results)))
			ui.Stat("Files/Lines processed", fmt.Sprintf("%d/%d", stats.TotalFiles, stats.TotalLines))
			ui.Stat("Search duration", stats.Duration.String())

			for _, result := range results {

				if err := ui.PurlVulns(result.Vulnerabilities); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output results in JSON format")

	return cmd
}
