package sync

import (
	"fmt"
	"slices"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
)

var specialIndices = []string{"cpecve"}

// syncResult is the structured summary emitted in JSON mode at the end of
// a sync. Agents can rely on .selected for the list of indices the command
// attempted to sync, and .elapsed_seconds for timing.
type syncResult struct {
	SchemaVersion  int      `json:"schema_version"`
	Action         string   `json:"action"`
	Selected       []string `json:"selected,omitempty"`
	ElapsedSeconds float64  `json:"elapsed_seconds,omitempty"`
}

func Command() *cobra.Command {
	var addIndices, removeIndices []string
	var purge bool
	var force bool

	cmd := &cobra.Command{
		Use:     "sync",
		Short:   "Sync indices",
		Long:    "Sync indices for offline use",
		Example: "vulncheck offline sync",
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			choose, _ := cmd.Flags().GetBool("choose")

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetIndices()
			if err != nil {
				return err
			}
			indices := response.GetData()

			availableIndices := make(map[string]bool)
			for _, index := range indices {
				availableIndices[index.Name] = true
			}
			for _, specialIndex := range specialIndices {
				availableIndices[specialIndex] = true
			}

			if purge {
				if err := cache.PurgeIndices(); err != nil {
					return fmt.Errorf("failed to purge indices: %w", err)
				}
				if r.IsJSON() {
					return r.JSON(syncResult{SchemaVersion: output.SchemaVersion, Action: "purge"})
				}
				r.Info("All cached indices have been purged.")
				return nil
			}

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

			for _, index := range addIndices {
				if !slices.Contains(selectedIndices, index) {
					selectedIndices = append(selectedIndices, index)
				}
			}

			for _, index := range removeIndices {
				selectedIndices = slices.DeleteFunc(selectedIndices, func(s string) bool {
					return s == index
				})
			}

			if (len(selectedIndices) == 0 && len(removeIndices) == 0) || choose {
				if !r.Interactive() {
					return errs.Validation("no indices to sync; pass --add <name> (repeatable), --remove <name>, or --purge")
				}
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

			startTime := time.Now()
			// Suppress the taskin TUI when we can't render into a real TTY
			// (--json, --no-interactive, CI, non-tty stdout). Any info about
			// per-index progress is out of scope in those modes — the final
			// JSON syncResult is the source of truth.
			disableUI := r.IsJSON() || !r.Interactive()
			if err := cache.IndicesSync(cmd.Context(), selectedIndices, force, disableUI); err != nil {
				return err
			}
			elapsed := time.Since(startTime)

			if r.IsJSON() {
				return r.JSON(syncResult{
					SchemaVersion:  output.SchemaVersion,
					Action:         "sync",
					Selected:       selectedIndices,
					ElapsedSeconds: elapsed.Seconds(),
				})
			}
			r.Info("Sync completed in: %s", elapsed)

			return nil
		},
	}

	cmd.Flags().BoolP("choose", "c", false, "Prompt to choose indices to sync, even if cached ones exist")
	cmd.Flags().StringSliceVar(&addIndices, "add", nil, "Add specific indices to sync, separated by commas")
	cmd.Flags().StringSliceVar(&removeIndices, "remove", nil, "Remove specific indices from sync")
	cmd.Flags().BoolVar(&purge, "purge", false, "Purge all cached indices")
	cmd.Flags().BoolVar(&force, "force", false, "Force a sync ignoring if the index is up-to-date")

	return cmd
}

// EnsureIndexSync checks if the given index is synced, and if not, prompts the user to sync it.
// It returns true if the index is available (either already synced or newly synced), and false otherwise.
//
// In non-interactive contexts (CI, --no-interactive, --json) it refuses to
// download silently — the caller must have synced the index ahead of time
// via `vulncheck offline sync --add <name>`.
func EnsureIndexSync(indices cache.InfoFile, indexType string, fail bool) (bool, error) {
	if indices.GetIndex(indexType) != nil {
		return true, nil
	}

	if config.IsCI() || fail || !output.Interactive() {
		return false, fmt.Errorf("index %s is required and not cached yet; run `vulncheck offline sync --add %s` first", indexType, indexType)
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

	indices, err := cache.Indices()
	if err != nil {
		return false, err
	}

	return indices.GetIndex(indexType) != nil, nil
}
