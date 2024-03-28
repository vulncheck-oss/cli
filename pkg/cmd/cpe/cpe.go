package cpe

import (
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "cpe <scheme>",
		Short: "Look up a specified cpe",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("cpe scheme is required")
			}
			response, err := session.Connect(config.Token()).GetCpe(args[0])
			if err != nil {
				return err
			}
			ray.Ray(response)
			ui.Json(response.GetData())
			return nil
		},
	}
}
