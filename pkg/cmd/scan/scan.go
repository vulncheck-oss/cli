package scan

import (
	"context"
	"fmt"
	"github.com/anchore/syft/syft"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
)

type Options struct {
	Json     bool
	Annotate bool
}

func Command() *cobra.Command {
	opts := &Options{
		Json:     false,
		Annotate: false,
	}

	cmd := &cobra.Command{
		Use:     "scan <path>",
		Short:   i18n.C.ScanShort,
		Example: i18n.C.ScanExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.ScanErrorDirectoryRequired)
			}

			src, err := syft.GetSource(context.Background(), args[0], nil)

			if err != nil {
				return err
			}

			sbom, err := syft.CreateSBOM(context.Background(), src, nil)

			var purls []string

			for p := range sbom.Artifacts.Packages.Enumerate() {
				purls = append(purls, p.PURL)
			}

			ui.Info(fmt.Sprintf(i18n.C.ScanPackagesFound, len(purls)))

			ui.NewProgress(len(purls))

			var vulns []sdk.PurlVulnerability

			for index, purl := range purls {
				response, err := session.Connect(config.Token()).GetPurl(purl)
				if err != nil {
					return err
				}
				if len(response.Data.Vulnerabilities) > 0 {
					vulns = append(vulns, response.Data.Vulnerabilities...)
				}
				ui.UpdateProgress(index + 1)
			}
			ui.Info(fmt.Sprintf(i18n.C.ScanCvesFound, len(vulns), len(purls)))

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	cmd.Flags().BoolVarP(&opts.Annotate, "annotate", "a", false, "Output as Github Annotations")

	return cmd

}
