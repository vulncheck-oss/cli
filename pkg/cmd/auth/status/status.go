package status

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Check authentication status",
		Long:  "Check if you're currently authenticated and if so, display the account information",
		RunE: func(cmd *cobra.Command, args []string) error {
			config.Init()

			if !config.HasConfig() {
				ui.Danger("No configuration found. Please run `vc auth login` to authenticate.")
			}

			return nil
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
