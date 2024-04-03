package auth

import (
	"github.com/spf13/cobra"

	cmdLogin "github.com/vulncheck-oss/cli/pkg/cmd/auth/login"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth/logout"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth/status"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auth <command>",
		Short:   i18n.C.AuthShort,
		GroupID: "core",
	}

	session.DisableAuthCheck(cmd)

	cmd.AddCommand(cmdLogin.Command())
	cmd.AddCommand(status.Command())
	cmd.AddCommand(logout.Command())

	return cmd
}
