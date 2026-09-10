package root

import (
	"github.com/MakeNowJust/heredoc/v2"
	"os"
)

func authHelp() string {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		return heredoc.Doc(`
			vulncheck: To use VulnCheck CLI in a GitHub Actions workflow, set the VULNCHECK_API_TOKEN environment variable. Example:
			  env:
			    VULNCHECK_API_TOKEN: ${{ secrets.VULNCHECK_API_TOKEN }}
			Substitute your own secret name; vulncheck-oss/action uses VC_TOKEN, which is also still accepted.
			Note that a secret that does not exist expands to an empty string, which reads here as no token at all.
		`)
	}

	if os.Getenv("CI") != "" {
		return heredoc.Doc(`
			vulncheck: To use VulnCheck CLI in automation, set the VULNCHECK_API_TOKEN environment variable.
		`)
	}

	return heredoc.Doc(`
		To get started with VulnCheck CLI, please run: vulncheck auth login
		Alternatively, populate the VULNCHECK_API_TOKEN environment variable with a VulnCheck token acquired from the portal at https://console.vulncheck.com/token.
	`)
}
