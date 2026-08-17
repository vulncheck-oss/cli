package logout

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: i18n.C.AuthLogoutShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			res := config.Resolve()

			// Logout can only clear the config file. When VC_TOKEN supplies the
			// active token, proceeding would revoke the *environment* token
			// server-side (config.Token() returns it) while deleting a
			// different token from disk, then report success even though
			// CheckAuth still passes off the env var. Refuse instead.
			if res.FromEnv() {
				e := errs.Validation(
					"%s is set; `auth logout` cannot unset an environment variable",
					config.EnvToken)
				if res.Shadowed {
					return e.WithHint(
						"run `unset %s`, then `vulncheck auth logout` again to revoke and remove the token saved in your config file.",
						config.EnvToken)
				}
				return e.WithHint(
					"run `unset %s` to stop using it, and revoke it with `vulncheck token remove <id>` if it should no longer work.",
					config.EnvToken)
			}

			if res.ConfigToken == "" {
				return errs.AuthRequired(i18n.C.ErrorNoToken)
			}

			// Revoke the config-file token specifically, not whatever
			// Resolve() happened to prefer.
			_, err := session.InvalidateToken(res.ConfigToken)
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
				ui.Info(i18n.C.AuthLogoutErrorInvalidToken)
				return nil
			}

			return err
		},
	}

	return cmd
}
