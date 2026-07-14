package indices

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
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

func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <search>",
		Short: i18n.C.ListIndicesShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetIndices()
			if err != nil {
				return err
			}

			if len(args) > 0 && args[0] != "" {
				indices := response.GetData()
				rows := ui.IndicesRows(indices, args[0])
				if r.IsJSON() {
					return r.JSON(rows)
				}
				r.Info(i18n.C.ListIndicesSearch, len(rows), args[0])
				return ui.IndicesList(indices, args[0])
			}

			if r.IsJSON() {
				return r.JSON(response.GetData())
			}

			r.Info(i18n.C.ListIndicesFull, len(response.GetData()))
			return ui.IndicesList(response.GetData(), "")
		},
	}

	return cmd
}

func Browse() *cobra.Command {
	return &cobra.Command{
		Use:   "browse <search>",
		Short: i18n.C.BrowseIndicesShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetIndices()
			if err != nil {
				return err
			}

			// Fall back to JSON list output when we can't render the TUI:
			// symmetric with `vulncheck index browse --json`. Agents get
			// a usable result instead of a validation error.
			if !r.Interactive() {
				search := ""
				if len(args) > 0 {
					search = args[0]
				}
				indices := response.GetData()
				if search != "" {
					return r.JSON(ui.IndicesRows(indices, search))
				}
				return r.JSON(indices)
			}
			search := ""
			if len(args) > 0 && args[0] != "" {
				search = args[0]
			}
			indices := response.GetData()

			for {
				ui.ClearScreen()

				if search != "" {
					r.Info(i18n.C.BrowseIndicesSearch, len(ui.IndicesRows(indices, search)), search)
				} else {
					r.Info(i18n.C.BrowseIndicesFull, len(ui.IndicesRows(indices, search)))
				}

				selectedIndex, err := ui.IndicesBrowse(indices, search)

				if err != nil {
					return err
				}

				if selectedIndex == "" {
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
