package login

import (
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/util"
	"strings"
)

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
	}

	web := &cobra.Command{
		Use:   "web",
		Short: "Log in with a VulnCheck account using a web browser",
		RunE: func(cmd *cobra.Command, args []string) error {
			return util.FlagErrorf("web login is not yet implemented")
		},
	}

	token := &cobra.Command{
		Use:   "token",
		Short: "Connect a VulnCheck account using a token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				token, err := ui.TokenPrompt()

				if err != nil {
					return util.FlagErrorf("Failed to read token: %v", err)
				}
				args = []string{token}
			}
			ray.Ray(args[0])
			if !ValidToken(args[0]) {
				return util.FlagErrorf("Invalid token specified")
			}

			spinner := ui.Spinner("Verifying Token...")
			res, err := session.CheckToken(args[0])
			if err != nil {
				return util.FlagErrorf("Token verification failed: %v", err)
			}
			spinner.Quit()
			ui.Success(fmt.Sprintf("Authenticated as %s (%s)", res.Data.Name, res.Data.Email))
			_, err = config.SaveToken(args[0])
			if err != nil {
				return util.FlagErrorf("Failed to save config: %v", err)
			}
			return nil
		},
	}

	cmd.AddCommand(web, token)

	session.DisableAuthCheck(cmd)
	return cmd
}

func ValidToken(token string) bool {
	if !strings.HasPrefix(token, "vulncheck_") {
		return false
	}

	if len(token) != 74 {
		return false
	}
	return true
}
