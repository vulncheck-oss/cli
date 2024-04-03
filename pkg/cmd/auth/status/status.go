package status

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/login"
	"github.com/vulncheck-oss/cli/pkg/session"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: i18n.C.AuthStatusShort,
		Long:  i18n.C.AuthStatusLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			config.Init()

			if !config.HasToken() {
				return fmt.Errorf(i18n.C.ErrorNoToken)
			}

			token := config.Token()
			return login.SaveToken(token)
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
