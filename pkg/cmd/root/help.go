package root

import (
	"github.com/MakeNowJust/heredoc/v2"
	"os"
)

func authHelp() string {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		return heredoc.Doc(`
			vc: To use GitHub CLI in a GitHub Actions workflow, set the GH_TOKEN environment variable. Example:
			  env:
			    VC_TOKEN: ${{ secrets.vulncheck_token }}
		`)
	}

	if os.Getenv("CI") != "" {
		return heredoc.Doc(`
			vc: To use VulnCheck CLI in automation, set the VC_TOKEN environment variable.
		`)
	}

	return heredoc.Doc(`
		To get started with VulnCheck CLI, please run: vc auth login
		Alternatively, populate the VC_TOKEN environment variable with a VulnCheck token acquired from the portal at https://vulncheck.com/token.
	`)
}
