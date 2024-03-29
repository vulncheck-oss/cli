package cpe

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "cpe <scheme>",
		Short: i18n.C.CpeShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.ErrorCpeSchemeRequired)
			}
			response, err := session.Connect(config.Token()).GetCpe(args[0])
			if err != nil {
				return err
			}
			cves := response.GetData()
			if err := ui.CpeMeta(response.CpeMeta()); err != nil {
				return err
			}
			if len(cves) == 0 {
				ui.Info(fmt.Sprintf(i18n.C.CpeNoCves, args[0]))
				return nil
			}
			ui.Info(fmt.Sprintf(i18n.C.CpeCvesFound, len(cves), args[0]))
			ui.Json(cves)
			return nil
		},
	}
}
