package version

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/build"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/session"
)

// versionInfo is the JSON payload for `vulncheck version --json`.
// Agents probe the CLI's version up-front to decide which flags they can
// safely use; keeping this shape stable across releases is the contract.
type versionInfo struct {
	Version      string `json:"version"`
	BuildDate    string `json:"build_date,omitempty"`
	ChangelogURL string `json:"changelog_url"`
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the current version, build date, and changelog URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if r.IsJSON() {
				return r.JSON(versionInfo{
					Version:      build.Version,
					BuildDate:    build.Date,
					ChangelogURL: session.ChangelogURL(build.Version),
				})
			}
			r.Println(session.VersionFormat(build.Version, build.Date))
			return nil
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
