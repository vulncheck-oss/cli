package root

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/cmd/upgrade"
	"os"

	"github.com/vulncheck-oss/cli/pkg/cmd/offline"

	"github.com/vulncheck-oss/cli/pkg/cmd/token"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/about"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth"
	"github.com/vulncheck-oss/cli/pkg/cmd/backup"
	"github.com/vulncheck-oss/cli/pkg/cmd/cpe"
	"github.com/vulncheck-oss/cli/pkg/cmd/index"
	"github.com/vulncheck-oss/cli/pkg/cmd/indices"
	"github.com/vulncheck-oss/cli/pkg/cmd/pdns"
	"github.com/vulncheck-oss/cli/pkg/cmd/purl"
	"github.com/vulncheck-oss/cli/pkg/cmd/rule"
	"github.com/vulncheck-oss/cli/pkg/cmd/scan"
	"github.com/vulncheck-oss/cli/pkg/cmd/tag"
	"github.com/vulncheck-oss/cli/pkg/cmd/version"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk-go"
)

type AuthError struct {
	err error
}

func (ae *AuthError) Error() string {
	return ae.err.Error()
}

func NewCmdRoot() *cobra.Command {
	i18n.Init()
	cmd := &cobra.Command{
		Use:   "vulncheck <command> <subcommand> [flags]",
		Short: "VulnCheck CLI.",
		Long:  i18n.C.RootLong,
		Example: heredoc.Doc(`
		$ vulncheck indices list
		$ vulncheck index abb
		$ vulncheck backup abb
	`),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			environment.Init()
			config.Init()

			if session.IsAuthCheckEnabled(cmd) && cmd.Parent().Name() != "completion" && !session.CheckAuth() {
				fmt.Println(authHelp())
				return ui.Error(i18n.C.ErrorNoToken)
			}

			return nil

		},
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core Commands",
	})

	cmd.AddCommand(version.Command())
	cmd.AddCommand(about.Command())
	cmd.AddCommand(auth.Command())
	cmd.AddCommand(token.Command())
	cmd.AddCommand(upgrade.Command())
	cmd.AddCommand(indices.Command())
	cmd.AddCommand(index.Command())
	cmd.AddCommand(backup.Command())
	cmd.AddCommand(cpe.Command())
	cmd.AddCommand(purl.Command())
	cmd.AddCommand(scan.Command())
	cmd.AddCommand(rule.Command())
	cmd.AddCommand(tag.Command())
	cmd.AddCommand(pdns.Command())
	cmd.AddCommand(offline.Command())

	return cmd
}

func Execute() {
	if err := NewCmdRoot().Execute(); err != nil {
		if errors.Is(err, sdk.ErrorUnauthorized) {
			fmt.Println(ui.Danger(i18n.C.ErrorUnauthorized))
		} else {
			fmt.Println(ui.Danger(err.Error()))
		}

		os.Exit(1)
	}
}
