package login

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type CmdCopy struct {
	Short string
	Long  string
}

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log in with a VulnCheck account",

		Long: heredoc.Docf(`
			Authenticate with a VulnCheck account.

			The default authentication mode is a web-based browser flow.

			Alternatively, use %[1]stoken%[1]s to specify an issued token directly.

			Alternatively, vc will use the authentication token found in the %[1]sVC_TOKEN%[1]s environment variable.
			This method is most suitable for "headless" use of vc such as in automation.
		`, "`"),
		Example: heredoc.Doc(`
			# Start interactive authentication
			$ vc auth login

			# Authenticate with vulncheck.com by passing in a token
			$ vc auth login token vulncheck_******************
		`),
		RunE: func(cmd *cobra.Command, args []string) error {

			if config.IsCI() {
				return ui.Error("This command is interactive and cannot be run in a CI environment, use the VC_TOKEN environment variable instead")
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
		Short: "Connect a VulnCheck account using an authentication token",
		RunE:  cmdToken,
	}

	web := &cobra.Command{
		Use:   "web",
		Short: "Log in with a VulnCheck account using a web browser",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.Error("web login is not yet implemented")
		},
	}

	cmd.AddCommand(web, token)

	session.DisableAuthCheck(cmd)
	return cmd
}
