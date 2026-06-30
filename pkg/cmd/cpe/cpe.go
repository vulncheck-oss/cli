package cpe

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
		Use:     "cpe <scheme>",
		Short:   i18n.C.CpeShort,
		Example: fmt.Sprintf(i18n.C.CpeExample, "cpe:2.3:a:sap:businessobjects_business_intelligence_platform:4.2:-:*"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.CpeErrorSchemeRequired)
			}
			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetCpe(args[0])
			if err != nil {
				return err
			}

			cves := response.GetData()
			if r.IsJSON() {
				return r.JSON(cves)
			}

			if err := ui.CpeMeta(response.CpeMeta()); err != nil {
				return err
			}
			if len(cves) == 0 {
				r.Info(i18n.C.CpeNoCves, args[0])
				return nil
			}
			r.Info(i18n.C.CpeCvesFound, len(cves), args[0])
			// In text mode the existing UI dumps the CVE list as pretty JSON.
			// Preserve that behaviour by routing through the renderer's
			// payload stream so it remains stdout-pure.
			return r.JSON(cves)
		},
	}

	return cmd
}
