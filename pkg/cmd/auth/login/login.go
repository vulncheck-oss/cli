package login

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
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
				return ui.Error(i18n.C.AuthLoginErrorCI)
			}

			if config.HasConfig() && config.HasToken() {
				if err := existingToken(); err != nil {
					return err
				}
			}

			choice, err := chooseAuthMethod()

			if err != nil {
				return err
			}

			switch choice {
			case "token":
				return cmdToken(cmd, args)
			case "web":
				return ui.Error("Command currently under construction")
			default:
				return ui.Error("Invalid choice")
			}
		},
	}

	token := &cobra.Command{
		Use:   "token",
		Short: i18n.C.AuthLoginToken,
		RunE:  cmdToken,
	}

	web := &cobra.Command{
		Use:   "web",
		Short: i18n.C.AuthLoginWeb,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.Error("web login is not yet implemented")
		},
	}

	cmd.AddCommand(web, token)

	session.DisableAuthCheck(cmd)
	return cmd
}
