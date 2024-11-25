package cpe

import (
	"fmt"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/cpeparse"
	"github.com/vulncheck-oss/cli/pkg/search"
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

			fmt.Printf("CPE Vendor is: %s\n", cpe.Vendor)

			ray.Ray(cpe)

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

			ray.Ray(query)

			entries, stats, err := search.Index(cpe.Vendor, query)

			if err != nil {
				return err
			}

			ray.Ray(entries, stats)

			return nil
		},
	}
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON format")
	return cmd
}
