package logout

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk-go"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: i18n.C.AuthLogoutShort,
		RunE: func(cmd *cobra.Command, args []string) error {

			if !config.HasToken() {
				return ui.Danger(i18n.C.ErrorNoToken)
			}

			_, err := session.InvalidateToken(config.Token())
			if err == nil {
				if err := config.RemoveToken(); err != nil {
					return ui.Danger(i18n.C.AuthLogoutErrorFailed)
				}
				ui.Success(i18n.C.AuthLogoutTokenRemoved)
				return nil
			}
			if errors.Is(err, sdk.ErrorUnauthorized) {
				if err := config.RemoveToken(); err != nil {
					return ui.Danger(i18n.C.AuthLogoutErrorInvalidToken)
				}
			}

			return err
		},
	}

	return cmd
}
