package root

import (
	_ "embed"
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/auth"
	"github.com/vulncheck-oss/cli/pkg/build"
	"github.com/vulncheck-oss/cli/pkg/cmd/ascii"
	cmdVersion "github.com/vulncheck-oss/cli/pkg/cmd/version"
	"os"
)

type AuthError struct {
	err error
}

func (ae *AuthError) Error() string {
	return ae.err.Error()
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vc <command> <subcommand> [flags]",
		Short: "VulnCheck CLI.",
		Long:  "Work seamlessly with the VulnCheck API.",
		Example: heredoc.Doc(`
		$ vc index list
		$ vc index abb
		$ vc backup abb
	`),
		Annotations: map[string]string{
			"versionInfo": cmdVersion.Format(build.Version, build.Date),
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			if auth.IsAuthCheckEnabled(cmd) && !auth.CheckAuth() {
				fmt.Println(authHelp())
				return &AuthError{}
			}

			return nil

		},
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.AddCommand(cmdVersion.Command())
	cmd.AddCommand(ascii.Command())

	return cmd
}

func Execute(version string, date string) {

	if err := NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
