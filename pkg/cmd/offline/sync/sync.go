package sync

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"os"
	"strings"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync",
		Short:   "Sync indices",
		Long:    "Sync indices for offline use",
		Example: "vulncheck offline sync",
		RunE: func(cmd *cobra.Command, args []string) error {

			response, err := session.Connect(config.Token()).GetIndices()

			if err != nil {
				return err
			}
			indices := response.GetData()

			selectedIndices := make([]string, 0)
			options := make([]huh.Option[string], len(indices))

			for i, index := range indices {
				options[i] = huh.Option[string]{Value: index.Name, Key: index.Name}
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

			ui.Info(fmt.Sprintf("Syncing indices: %s", strings.Join(selectedIndices, ", ")))

			IndicesSync(selectedIndices)

			return nil
		},
	}

	return cmd
}

func IndicesSync(indices []string) error {
	for _, index := range indices {
		response, err := session.Connect(config.Token()).GetIndexBackup(index)
		if err != nil {
			return err
		}

		file, err := utils.ExtractFile(response.GetData()[0].URL)
		if err != nil {
			return err
		}

		configDir, err := config.IndicesDir()
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s/%s", configDir, file)
		indexDir := fmt.Sprintf("%s/%s", configDir, index)

		date := utils.ParseDate(response.GetData()[0].DateAdded)

		ui.Info(fmt.Sprintf("[%s] last updated %s", index, date))
		ui.Info(fmt.Sprintf("[%s] Downloading %s", index, file))

		if err := ui.Download(response.GetData()[0].URL, filePath); err != nil {
			return err
		}

		// Check if the index directory exists
		if _, err := os.Stat(indexDir); !os.IsNotExist(err) {
			// Remove the existing directory and its contents
			if err := os.RemoveAll(indexDir); err != nil {
				return fmt.Errorf("failed to remove existing index directory: %w", err)
			}
		}

		// Create the index directory
		if err := os.MkdirAll(indexDir, 0755); err != nil {
			return fmt.Errorf("failed to create index directory: %w", err)
		}

		// Unzip the downloaded file into the index directory
		if err := utils.Unzip(filePath, indexDir); err != nil {
			return fmt.Errorf("failed to unzip index file: %w", err)
		}

		// Calculate and display the size of the extracted index
		size, err := utils.GetDirectorySize(indexDir)
		if err != nil {
			ui.Error(fmt.Sprintf("Failed to calculate size of index directory: %s", err))
		} else {
			ui.Info(fmt.Sprintf("Successfully synced %s (Size: %s)", index, size))
		}

		// Optionally, remove the downloaded zip file
		if err := os.Remove(filePath); err != nil {
			ui.Error(fmt.Sprintf("Failed to remove downloaded zip file: %s", err))
		}
	}

	return nil
}
