package token

import (
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/login"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "token",
		Short: i18n.C.AuthLoginToken,
		RunE:  CmdToken,
	}
}

func CmdToken(cmd *cobra.Command, args []string) error {

	var token string

	input := huh.
		NewInput().
		Title("Enter your authentication token").
		EchoMode(huh.EchoModePassword).
		Placeholder("vulncheck_******************").
		Value(&token)

	if err := input.Run(); err != nil {
		return ui.Error("Token verification failed: %v", err)
	}

	if !config.ValidToken(token) {
		return ui.Error("Invalid token specified")
	}

	return login.SaveToken(token)
}
