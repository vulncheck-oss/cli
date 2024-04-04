package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"regexp"
)

type Options struct {
	Json bool
}

func Command() *cobra.Command {

	opts := &Options{
		Json: false,
	}

	cmd := &cobra.Command{
		Use:   "list <id>",
		Short: i18n.C.SbomScanListShort,
		RunE: func(cmd *cobra.Command, args []string) error {

			// validate args[0] as a proper uuid

			if len(args) == 0 {
				return fmt.Errorf("an SBOM UUID must be provided")
			}

			uuid := args[0]
			matched, err := regexp.MatchString(`^\w{10}-\w{4}-\w{4}-\w{4}-\w{10}$`, uuid)
			if err != nil {
				return fmt.Errorf("error validating UUID format: %v", err)
			}
			if !matched {
				return fmt.Errorf("the provided SBOM UUID '%s' is not in the valid format (10-4-4-4-10)", uuid)
			}

			response, err := session.Connect(config.Token()).SbomScans(args[0])
			if err != nil {
				return err
			}

			if opts.Json == true {
				ui.Json(response.GetData())
				return nil
			}

			if err := ui.SbomScans(response.GetData()); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	return cmd
}
