package web

import (
	"fmt"
	"github.com/charmbracelet/huh/spinner"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/inquiry"
	"github.com/vulncheck-oss/cli/pkg/login"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

/**
step 1. generate an inquiry.
step 2. prompt the user to visit the inquiry URL.
step 3. loop and sleep waiting for an inquiry response.
*/

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: i18n.C.AuthLoginWeb,
		RunE:  CmdWeb,
	}
}

func CmdWeb(cmd *cobra.Command, args []string) error {

	if !inquiry.IsPortAvailable(inquiry.Port) {
		return fmt.Errorf("this method is not available, try vc auth login token")
	}

	ui.Info("Attempting to launch console.vulncheck.com in your browser...")

	var errorResponse error
	var token string

	_ = spinner.New().
		Style(ui.Pantone).
		Title(" Awaiting Verification...").Action(func() {

		if err := browser.OpenURL(fmt.Sprintf("%s/inquiry", environment.Env.WEB)); err != nil {
			errorResponse = err
			return
		}

		_, err := inquiry.ListenForHash()
		if err != nil {
			errorResponse = err
			return
		}

		tkn, err := inquiry.ListenForToken()
		if err != nil {
			errorResponse = err
			return
		}
		if tkn == "denied" {
			errorResponse = fmt.Errorf("request denied")
			return
		}
		token = tkn

	}).Run()

	if errorResponse != nil {
		return errorResponse
	}

	return login.SaveToken(token)
}
