package scan

import (
	"context"
	"fmt"
	"github.com/anchore/syft/syft"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
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

			ui.Info(fmt.Sprintf(i18n.C.ScanPurlsFound, len(purls)))

			for _, purl := range purls {
				response, err := session.Connect(config.Token()).GetPurl(purl)
				if err != nil {
					return err
				}
				if len(response.Data.Cves) > 0 {
					ray.Ray(response.Data)
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	cmd.Flags().BoolVarP(&opts.Annotate, "annotate", "a", false, "Output as Github Annotations")

	return cmd

}
