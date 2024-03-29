package purl

import (
	"fmt"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "purl <scheme>",
		Short: "Look up a specified PURL for any CVES or vulnerabilities",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("purl scheme is required")
			}
			response, err := session.Connect(config.Token()).GetPurl(args[0])
			if err != nil {
				return err
			}
			cves := response.Cves()
			ray.Ray(response.PurlMeta())
			if err := ui.PurlMeta(response.PurlMeta()); err != nil {
				return err
			}
			if len(cves) == 0 {
				ui.Info(fmt.Sprintf("No CVEs were found for purl %s", args[0]))
				return nil
			}
			ui.Info(fmt.Sprintf("%d CVEs were found for purl %s", len(cves), args[0]))
			ui.Json(cves)
			return nil
		},
	}
}
