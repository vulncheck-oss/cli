package login

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth/login/token"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth/login/web"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	pkgLogin "github.com/vulncheck-oss/cli/pkg/login"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type CmdCopy struct {
	Short string
	Long  string
}

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "login",
		Short:   i18n.C.AuthLoginShort,
		Long:    i18n.C.AuthLoginLong,
		Example: i18n.C.AuthLoginExample,
		RunE: func(cmd *cobra.Command, args []string) error {

			if config.IsCI() {
				return ui.Error("%s", i18n.C.AuthLoginErrorCI)
			}

			if config.HasConfig() && config.HasToken() {
				if err := pkgLogin.ExistingToken(); err != nil {
					return err
				}
			}

			choice, err := pkgLogin.ChooseAuthMethod()

			if err != nil {
				return err
			}

			switch choice {
			case "token":
				return token.CmdToken(cmd, args)
			case "web":
				return web.CmdWeb(cmd, args)
			default:
				return ui.Error("%s", "Invalid choice")
			}
		},
	}

	cmd.AddCommand(web.Command(), token.Command())

	session.DisableAuthCheck(cmd)
	return cmd
}
