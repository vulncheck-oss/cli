package cpe

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeuri"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"github.com/vulncheck-oss/cli/pkg/db"
)

func Command() *cobra.Command {
	var statsOnly bool

	cmd := &cobra.Command{
		Use:     "cpe <scheme>",
		Short:   "Offline CPE lookup",
		Long:    "Search offline package data via CPE schemes",
		Example: "vulncheck offline cpe \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

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

			results, stats, err := db.CPESearch("cpecve", *cpe)
			if err != nil {
				return err
			}
			cves, err := cpeutils.Process(cpe, results)
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(cves)
			}

			r.Stat("Results found/filtered", fmt.Sprintf("%d/%d", len(results), len(cves)))
			r.Stat("Search duration", fmt.Sprintf("%.2f seconds", stats.Duration.Seconds()))

			if !statsOnly {
				return r.JSON(cves)
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&statsOnly, "stats", "s", false, "Output stats only")
	return cmd
}
