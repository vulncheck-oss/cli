package session

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/sdk"
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

func CheckAuth() bool {
	return true
}

func CheckToken(token string) (response *sdk.UserResponse, err error) {
	return sdk.Connect(environment.Env.API, token).GetMe()
}
