package purl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "purl <scheme>",
		Short:   i18n.C.PurlShort,
		Example: fmt.Sprintf(i18n.C.PurlExample, "pkg:hackage/aeson@0.3.2.8"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.PurlErrorSchemeRequired)
			}
			response, err := session.Connect(config.Token()).GetPurl(args[0])
			if err != nil {
				return fmt.Errorf("error fetching purl %s: %w", args[0], err)
			}

			if r.IsJSON() {
				return r.JSON(response.GetData())
			}

			vulns := response.Vulnerabilities()
			if err := ui.PurlMeta(response.PurlMeta()); err != nil {
				return err
			}
			if len(vulns) == 0 {
				r.Info(i18n.C.PurlNoVulns, args[0])
				return nil
			}
			if len(vulns) == 1 {
				r.Info(i18n.C.PurlVulnFound, args[0])
			} else {
				r.Info(i18n.C.PurlVulnsFound, len(vulns), args[0])
			}
			return ui.PurlVulns(vulns)
		},
	}

	return cmd
}
