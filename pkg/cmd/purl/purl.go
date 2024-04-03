package purl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type Options struct {
	Json bool
}

func Command() *cobra.Command {
	opts := &Options{
		Json: false,
	}

	cmd := &cobra.Command{
		Use:     "purl <scheme>",
		Short:   i18n.C.PurlShort,
		Example: fmt.Sprintf(i18n.C.PurlExample, "pkg:hackage/aeson@0.3.2.8"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.ErrorPurlSchemeRequired)
			}
			response, err := session.Connect(config.Token()).GetPurl(args[0])
			if err != nil {
				return err
			}

			if opts.Json {
				ui.Json(response.GetData())
				return nil
			}
			cves := response.Cves()
			if err := ui.PurlMeta(response.PurlMeta()); err != nil {
				return err
			}
			if len(cves) == 0 {
				ui.Info(fmt.Sprintf(i18n.C.PurlNoCves, args[0]))
				return nil
			}
			ui.Info(fmt.Sprintf(i18n.C.PurlCvesFound, len(cves), args[0]))
			ui.Json(cves)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")

	return cmd
}
