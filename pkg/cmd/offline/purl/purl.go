package purl

import (
	"fmt"
	"github.com/package-url/packageurl-go"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/db"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"github.com/vulncheck-oss/sdk-go"
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

			/*
				if packages.IsOS(instance) {
					return fmt.Errorf("operating system package support coming soon")
				}
			*/

			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			indexName := packages.IndexFromInstance(instance)

			indexAvailable, err := sync.EnsureIndexSync(indices, indexName, false)
			if err != nil {
				return err
			}

			if !indexAvailable {
				return fmt.Errorf("index %s is required to proceed", instance.Type)
			}

			indices, err = cache.Indices()

			if err != nil {
				return err
			}

			index := indices.GetIndex(indexName)

			if !jsonOutput && !config.IsCI() {
				if err := ui.PurlInstance(instance); err != nil {
					return err
				}
				ui.Info(fmt.Sprintf("Searching index %s, last updated on %s", index.Name, utils.ParseDate(index.LastUpdated)))
			}

			results, stats, err := db.PURLSearch(index.Name, instance)
			if err != nil {
				return err
			}

			if jsonOutput || config.IsCI() {

				// Create a combined structure for JSON output
				combinedOutput := struct {
					Instance        packageurl.PackageURL   `json:"instance"`
					Vulnerabilities []sdk.PurlVulnerability `json:"vulnerabilities"`
				}{
					Instance: instance,
				}

				// Collect all vulnerabilities
				for _, result := range results {
					combinedOutput.Vulnerabilities = append(combinedOutput.Vulnerabilities, result.Vulnerabilities...)
				}
				ui.Json(combinedOutput)
				return nil
			}

			ui.Stat("Results found", fmt.Sprintf("%d", len(results)))
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
