package logout

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Invalidate and remove your current authentication token",
		RunE: func(cmd *cobra.Command, args []string) error {

			if !config.HasConfig() {
				return ui.Danger("No configuration file found")
			}

			if !config.HasToken() {
				return ui.Danger("No valid token was found")
			}

			_, err := session.InvalidateToken(config.Token())
			if err == nil {
				if err := config.RemoveToken(); err != nil {
					return ui.Danger("Failed to remove token")
				}
				return ui.Success("Token successfully invalidated")
			}
			if err.Error() == session.ErrorUnauthorized {
				if err := config.RemoveToken(); err != nil {
					return ui.Danger("Token was invalid, removing from config")
				}
			}

			return err
		},
	}

	return cmd
}
