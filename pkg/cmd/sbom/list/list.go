package list

import (
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
		Use:   "list <search>",
		Short: i18n.C.SbomListShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := session.Connect(config.Token()).SbomList()
			if err != nil {
				return err
			}

			if opts.Json == true {
				ui.Json(response.GetData())
				return nil
			}

			if err := ui.SbomList(response.GetData()); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	return cmd
}
