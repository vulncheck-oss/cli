package root

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/build"
	"github.com/vulncheck-oss/cli/pkg/cmd/ascii"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth"
	"github.com/vulncheck-oss/cli/pkg/cmd/indices"
	cmdVersion "github.com/vulncheck-oss/cli/pkg/cmd/version"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
	"os"
)

type AuthError struct {
	err error
}

type exitCode int

const (
	exitOK        exitCode = 0
	exitError     exitCode = 1
	exitCancel    exitCode = 2
	exitAuthError exitCode = 3
)

func (ae *AuthError) Error() string {
	return ae.err.Error()
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vc <command> <subcommand> [flags]",
		Short: "VulnCheck CLI.",
		Long:  "Work seamlessly with the VulnCheck API.",
		Example: heredoc.Doc(`
		$ vc indices list
		$ vc index abb
		$ vc backup abb
	`),
		Annotations: map[string]string{
			"versionInfo": cmdVersion.Format(build.Version, build.Date),
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			environment.Init()
			config.Init()

			if session.IsAuthCheckEnabled(cmd) && !session.CheckAuth() {
				fmt.Println(authHelp())
				return &AuthError{}
			}

			return nil

		},
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core Commands",
	})

	cmd.AddCommand(cmdVersion.Command())
	cmd.AddCommand(ascii.Command())
	cmd.AddCommand(auth.Command())
	cmd.AddCommand(indices.Command())

	return cmd
}

func Execute() {
	if err := NewCmdRoot().Execute(); err != nil {
		if errors.Is(err, sdk.ErrorUnauthorized) {
			fmt.Println(ui.Danger("Error: %s, Try authenticating with: vc auth login", err.Error()))
			os.Exit(1)
		}
		// fmt.Println(err)
		// os.Exit(1)
	}
}
