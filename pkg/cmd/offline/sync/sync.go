package sync

import (
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync",
		Short:   "Sync indices",
		Long:    "Sync indices for offline use",
		Example: "vulncheck offline sync",
		RunE: func(cmd *cobra.Command, args []string) error {
			choose, _ := cmd.Flags().GetBool("choose")

			response, err := session.Connect(config.Token()).GetIndices()

			if err != nil {
				return err
			}
			indices := response.GetData()

			indexInfo, err := cache.Indices()
			selectedIndices := make([]string, 0, len(indexInfo.Indices))

			for _, info := range indexInfo.Indices {
				selectedIndices = append(selectedIndices, info.Name)
			}
			if err != nil {
				return err
			}

			if len(selectedIndices) == 0 || choose {

				options := make([]huh.Option[string], len(indices))

				for i, index := range indices {
					options[i] = huh.Option[string]{
						Value: index.Name,
						Key:   index.Name,
					}
				}

				form := huh.NewForm(
					huh.NewGroup(
						huh.NewMultiSelect[string]().
							Title("Select indices to sync").
							Options(options...).
							Height(10).
							Filterable(true).
							Value(&selectedIndices),
					),
				)

				err = form.Run()
				if err != nil {
					return err
				}
			}

			if err := cache.IndicesSync(selectedIndices); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolP("choose", "c", false, "Prompt to choose indices to sync, even if cached ones exist")

	return cmd
}
