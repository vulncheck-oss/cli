package purl

import (
	"fmt"

	"github.com/package-url/packageurl-go"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/db"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "purl <scheme>",
		Short:   "Offline PURL lookup",
		Long:    "Search offline package data via PURL schemes",
		Example: "vulncheck offline purl \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			instance, err := packageurl.FromString(args[0])
			if err != nil {
				return err
			}

			if packages.IsOS(instance) {
				return fmt.Errorf("offline PURL lookups for this operating system are not supported yet - please contact support@vulncheck.com")
			}

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

			if !r.IsJSON() {
				if err := ui.PurlInstance(instance); err != nil {
					return err
				}
				r.Info("Searching index %s, last updated on %s", index.Name, utils.ParseDate(index.LastUpdated))
			}

			results, stats, err := db.PURLSearch(index.Name, instance)
			if err != nil {
				return err
			}

			if r.IsJSON() {
				combinedOutput := struct {
					Instance        packageurl.PackageURL   `json:"instance"`
					Vulnerabilities []sdk.PurlVulnerability `json:"vulnerabilities"`
				}{
					Instance: instance,
				}
				for _, result := range results {
					combinedOutput.Vulnerabilities = append(combinedOutput.Vulnerabilities, result.Vulnerabilities...)
				}
				return r.JSON(combinedOutput)
			}

			vulnsFound := 0
			for _, result := range results {
				vulnsFound += len(result.Vulnerabilities)
			}

			r.Stat("Results found", fmt.Sprintf("%d", len(results)))
			r.Stat("Vulnerabilities found", fmt.Sprintf("%d", vulnsFound))
			r.Stat("Search duration", stats.Duration.String())

			for _, result := range results {
				if err := ui.PurlVulns(result.Vulnerabilities); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return cmd
}
