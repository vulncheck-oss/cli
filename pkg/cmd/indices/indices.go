package indices

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
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

			action := func(index string) error {
				indexCommand := cmd.Root()
				indexCommand.SetArgs([]string{"index", "browse", index})
				if err := indexCommand.Execute(); err != nil {
					return err
				}
				return nil
			}

			response, err := session.Connect(config.Token()).GetIndices()
			if err != nil {
				return err
			}
			if len(args) > 0 && args[0] != "" {
				indices := response.GetData()
				ui.Info(fmt.Sprintf("Browsing %d indices searching for \"%s\"", len(ui.IndicesRows(indices, args[0])), args[0]))
				return ui.Indices(indices, args[0], action)
			}

			ui.Info(fmt.Sprintf("Browsing %d indices", len(response.GetData())))
			return ui.Indices(response.GetData(), "", action)
		},
	}

}
