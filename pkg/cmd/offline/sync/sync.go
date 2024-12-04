package sync

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"slices"
)

var specialIndices = []string{"cpecve"}

func Command() *cobra.Command {

	var addIndices, removeIndices []string
	var purge bool

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

			// Create a map of available indices for quick lookup
			availableIndices := make(map[string]bool)
			for _, index := range indices {
				availableIndices[index.Name] = true
			}

			// Add special indices to availableIndices
			for _, specialIndex := range specialIndices {
				availableIndices[specialIndex] = true
			}

			// Handle purge flag
			if purge {
				if err := cache.PurgeIndices(); err != nil {
					return fmt.Errorf("failed to purge indices: %w", err)
				}
				ui.Info("All cached indices have been purged.")
				return nil
			}

			// Validate addIndices and removeIndices
			for _, index := range append(addIndices, removeIndices...) {
				if !availableIndices[index] {
					return fmt.Errorf("index '%s' does not exist", index)
				}
			}

			indexInfo, err := cache.Indices()
			if err != nil {
				return err
			}

			selectedIndices := make([]string, 0, len(indexInfo.Indices))
			for _, info := range indexInfo.Indices {
				selectedIndices = append(selectedIndices, info.Name)
			}

			// Add indices
			for _, index := range addIndices {
				if !slices.Contains(selectedIndices, index) {
					selectedIndices = append(selectedIndices, index)
				}
			}

			// Remove indices
			for _, index := range removeIndices {
				selectedIndices = slices.DeleteFunc(selectedIndices, func(s string) bool {
					return s == index
				})
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
	cmd.Flags().StringSliceVar(&addIndices, "add", nil, "Add specific indices to sync")
	cmd.Flags().StringSliceVar(&removeIndices, "remove", nil, "Remove specific indices from sync")
	cmd.Flags().BoolVar(&purge, "purge", false, "Purge all cached indices")

	return cmd
}

// EnsureIndexSync checks if the given index is synced, and if not, prompts the user to sync it.
// It returns true if the index is available (either already synced or newly synced), and false otherwise.
func EnsureIndexSync(indices cache.InfoFile, indexType string, fail bool) (bool, error) {

	if indices.GetIndex(indexType) != nil {
		return true, nil
	}

	if config.IsCI() || fail {
		return false, fmt.Errorf("index %s is required and not cached yet", indexType)
	}

	shouldSync := true
	prompt := huh.NewConfirm().
		Title(fmt.Sprintf("Index %s is required and not cached yet. Do you want to download it?", indexType)).
		Value(&shouldSync).WithTheme(huh.ThemeCatppuccin())

	if err := prompt.Run(); err != nil {
		return false, err
	}

	if !shouldSync {
		return false, nil
	}

	syncCmd := Command()
	syncCmd.SetArgs([]string{"--add", indexType})
	if err := syncCmd.Execute(); err != nil {
		return false, fmt.Errorf("failed to sync index: %w", err)
	}

	// Refresh indices after syncing
	indices, err := cache.Indices()
	if err != nil {
		return false, err
	}

	return indices.GetIndex(indexType) != nil, nil
}
