package index

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
)

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "index <command>",
		Short: "Browse or list an index",
	}

	cmdList := &cobra.Command{
		Use:   "list <index>",
		Short: "List documents of a specified index",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("index name is required")
			}
			response, err := sdk.Connect(environment.Env.API, config.Token()).GetIndex(args[0])
			if err != nil {
				return err
			}
			ui.Json(response.GetData())
			return nil
		},
	}

	cmdBrowse := &cobra.Command{
		Use:   "browse <index>",
		Short: "Browse documents of an index interactively",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("index name is required")
			}
			response, err := sdk.Connect(environment.Env.API, config.Token()).GetIndex(args[0])
			if err != nil {
				return err
			}
			ui.Viewport(args[0], response.GetData())
			return nil
		},
	}

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdBrowse)

	return cmd
}
