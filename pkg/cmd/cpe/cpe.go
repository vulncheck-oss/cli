package cpe

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "cpe <scheme>",
		Short: "Look up a specified cpe for any related CVEs",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("cpe scheme is required")
			}
			response, err := session.Connect(config.Token()).GetCpe(args[0])
			if err != nil {
				return err
			}
			cves := response.GetData()
			if err := ui.CpeMeta(response.GetCpeMeta()); err != nil {
				return err
			}
			if len(cves) == 0 {
				ui.Info(fmt.Sprintf("No CVEs were found for cpe %s", args[0]))
				return nil
			}
			ui.Info(fmt.Sprintf("%d CVEs were found for cpe %s", len(cves), args[0]))
			ui.Json(cves)
			return nil
		},
	}
}
