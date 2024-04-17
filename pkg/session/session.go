package session

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/build"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/sdk"
	"regexp"
	"strings"
)

type MeResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Data      Me      `json:"data"`
}

type Me struct {
	ID     string
	Email  string
	Name   string
	Avatar string
}

func CheckAuth() bool {
	token := config.Token()
	if token != "" && config.ValidToken(token) {
		return true
	}
	return false
}

func ChangelogURL(version string) string {
	path := "https://github.com/vulncheck-oss/cli"
	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, strings.TrimPrefix(version, "v"))
	return url
}

func VersionFormat(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	var dateStr string
	if buildDate != "" {
		dateStr = fmt.Sprintf(" (%s)", buildDate)
	}

	return fmt.Sprintf("vc version %s%s\n%s\n", version, dateStr, ChangelogURL(version))
}

func Connect(token string) *sdk.Client {
	return sdk.Connect(environment.Env.API, token).SetUserAgent(fmt.Sprintf("VulnCheck CLI %s", build.Version))
}

func CheckToken(token string) (response *sdk.UserResponse, err error) {
	return Connect(token).GetMe()
}
func InvalidateToken(token string) (response *sdk.Response, err error) {
	return Connect(token).Logout()
}

func IsAuthCheckEnabled(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "help", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd:
		return false
	}

	for c := cmd; c.Parent() != nil; c = c.Parent() {
		if c.Annotations != nil && c.Annotations["skipAuthCheck"] == "true" {
			return false
		}
	}

	return true
}

func DisableAuthCheck(cmd *cobra.Command) {
	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	}

	cmd.Annotations["skipAuthCheck"] = "true"
}
