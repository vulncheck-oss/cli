package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/auth"
	"regexp"
	"strings"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the current version, build date, and changelog URL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Root().Annotations["versionInfo"])
		},
	}

	auth.DisableAuthCheck(cmd)
	return cmd
}

func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	var dateStr string
	if buildDate != "" {
		dateStr = fmt.Sprintf(" (%s)", buildDate)
	}

	return fmt.Sprintf("vc version %s%s\n%s\n", version, dateStr, changelogURL(version))
}

func changelogURL(version string) string {
	path := "https://github.com/vulncheck-oss/cli"
	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, strings.TrimPrefix(version, "v"))
	return url
}
