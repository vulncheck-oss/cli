package indices

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "indices <command>",
		Short: i18n.C.IndicesShort,
	}

	cmd.AddCommand(List())
	cmd.AddCommand(Browse())

	return cmd
}

type ListOptions struct {
	Json bool
}

func List() *cobra.Command {

	opts := &ListOptions{
		Json: false,
	}

	cmd := &cobra.Command{
		Use:   "list <search>",
		Short: i18n.C.ListIndicesShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := session.Connect(config.Token()).GetIndices()
			if err != nil {
				return err
			}
			if len(args) > 0 && args[0] != "" {
				indices := response.GetData()
				ui.Info(fmt.Sprintf(i18n.C.ListIndicesSearch, len(ui.IndicesRows(indices, args[0])), args[0]))
				return ui.IndicesList(indices, args[0])
			}
			ui.Info(fmt.Sprintf(i18n.C.ListIndicesFull, len(response.GetData())))
			if opts.Json {
				ui.Json(response.GetData())
				return nil
			}
			if err := ui.IndicesList(response.GetData(), ""); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	return cmd
}

func Browse() *cobra.Command {

	return &cobra.Command{
		Use:   "browse <search>",
		Short: i18n.C.BrowseIndicesShort,
		RunE: func(cmd *cobra.Command, args []string) error {

			response, err := session.Connect(config.Token()).GetIndices()

			if err != nil {
				return err
			}
			search := ""
			if len(args) > 0 && args[0] != "" {
				search = args[0]
			}
			indices := response.GetData()
			if search != "" {
				ui.Info(fmt.Sprintf(i18n.C.BrowseIndicesSearch, len(ui.IndicesRows(indices, search)), search))
			} else {
				ui.Info(fmt.Sprintf(i18n.C.BrowseIndicesFull, len(ui.IndicesRows(indices, search))))
			}
			for {

				selectedIndex, err := ui.IndicesBrowse(indices, search)

				if err != nil {
					return err
				}

				if selectedIndex == "" {
					// User quit the browse view
					return nil
				}

				indexCommand := cmd.Root()
				indexCommand.SetArgs([]string{"index", "browse", selectedIndex})
				if err := indexCommand.Execute(); err != nil {
					return err
				}
			}

		},
	}

}
