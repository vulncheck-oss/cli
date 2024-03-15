package auth

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth/login"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth/status"
	"github.com/vulncheck-oss/cli/pkg/session"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auth <command",
		Short:   "Authenticate vc with the VulnCheck portal",
		GroupID: "core",
	}

	session.DisableAuthCheck(cmd)

	cmd.AddCommand(login.Command())
	cmd.AddCommand(status.Command())

	return cmd
}
