package indices

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "indices <command>",
		Short: "Manage indices",
	}

	cmd.AddCommand(Browse())

	return cmd
}

func Browse() *cobra.Command {

	return &cobra.Command{
		Use:   "browse <search>",
		Short: "Browse indices",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := sdk.Connect(environment.Env.API, config.Token()).GetIndices()
			if err != nil {
				return err
			}
			if len(args) > 0 && args[0] != "" {
				indices := response.GetData()
				ui.Info(fmt.Sprintf("Browsing %d indices searching for \"%s\"", len(ui.IndicesRows(indices, args[0])), args[0]))
				return ui.Indices(indices, args[0])
			}

			ui.Info(fmt.Sprintf("Browsing %d indices", len(response.GetData())))
			return ui.Indices(response.GetData(), "")
		},
	}

}
