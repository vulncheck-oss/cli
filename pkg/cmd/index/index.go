package index

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
)

type indexOptions struct {
	Json   bool
	Browse bool
}

func Command() *cobra.Command {

	opts := &indexOptions{}

	cmd := &cobra.Command{
		Use:   "index <index>",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("index name is required")
			}
			response, err := sdk.Connect(environment.Env.API, config.Token()).GetIndex(args[0])
			if err != nil {
				return err
			}

			if opts.Json {
				ui.Json(response.GetData())
				return nil
			}

			if opts.Browse {
				ui.Viewport(args[0], response.GetData())
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&opts.Json, "json", false, "Output as JSON")
	cmd.Flags().BoolVar(&opts.Browse, "browse", false, "Browse the index in a pager")

	return cmd
}
