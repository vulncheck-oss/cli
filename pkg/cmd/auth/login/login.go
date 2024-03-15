package login

import (
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/util"
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

			if config.HasConfig() && config.HasToken() {
				logoutChoice := true
				confirm := huh.NewConfirm().
					Title("You currently have a token saved. Do you want to invalidate it first?").
					Affirmative("Yes").
					Negative("No").
					Value(&logoutChoice)
				confirm.Run()

				if logoutChoice {
					if _, err := session.InvalidateToken(config.Token()); err != nil {
						if err := config.RemoveToken(); err != nil {
							return ui.Danger("Failed to remove token from config")
						}
						return ui.Info("Token was not valid, removing from config")
					} else {
						if err := config.RemoveToken(); err != nil {
							return ui.Danger("Failed to remove token from config")
						}
						ui.Success("Token invalidated successfully")
					}
				} else {
					return nil
				}

			}

			var choice string
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Select an authentication method").
						Options(
							huh.NewOption("Login with a web browser", "web"),
							huh.NewOption("Paste an authentication token", "token"),
						).Value(&choice),
				),
			)

			err := form.Run()
			if err != nil {
				return util.FlagErrorf("Failed to select authentication method: %v", err)
			}

			switch choice {
			case "token":
				return cmdToken(cmd, args)
			default:
				return util.FlagErrorf("Invalid choice")
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
			return util.FlagErrorf("web login is not yet implemented")
		},
	}

	cmd.AddCommand(web, token)

	session.DisableAuthCheck(cmd)
	return cmd
}

func cmdToken(cmd *cobra.Command, args []string) error {

	var token string

	input := huh.
		NewInput().
		Title("Enter your authentication token").
		Password(true).
		Placeholder("vulncheck_******************").
		Value(&token)

	if err := input.Run(); err != nil {
		return ui.Danger(fmt.Sprintf("Token verification failed: %v", err))
	}

	if !config.ValidToken(token) {
		return util.FlagErrorf("Invalid token specified")
	}

	return SaveToken(token)
}

func SaveToken(token string) error {

	spinner := ui.Spinner("Verifying Token...")
	res, err := session.CheckToken(token)
	if err != nil {
		return ui.Danger(fmt.Sprintf("Token verification failed: %v", err))
	}
	spinner.Quit()
	if err := config.SaveToken(token); err != nil {
		return util.FlagErrorf("Failed to save token: %v", err)
	}
	return ui.Success(fmt.Sprintf("Authenticated as %s (%s)", res.Data.Name, res.Data.Email))
}
