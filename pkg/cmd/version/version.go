package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/build"
	"github.com/vulncheck-oss/cli/pkg/session"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the current version, build date, and changelog URL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(session.VersionFormat(build.Version, build.Date))
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
