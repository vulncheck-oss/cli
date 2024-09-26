package sync

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"time"
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

			selectedIndices, err := IndicesCurrent()
			if err != nil {
				return err
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

			ui.Info(fmt.Sprintf("Syncing indices: %s", strings.Join(selectedIndices, ", ")))

			if err := IndicesSync(selectedIndices); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

type IndexInfo struct {
	Name        string    `yaml:"name"`
	LastSync    time.Time `yaml:"last_sync"`
	Size        uint64    `yaml:"size"`         // Size in bytes
	LastUpdated string    `yaml:"last_updated"` // From the server
}

type InfoFile struct {
	Indices []IndexInfo `yaml:"indices"`
}

func IndicesCurrent() ([]string, error) {
	configDir, err := config.IndicesDir()
	if err != nil {
		return nil, err
	}

	infoPath := filepath.Join(configDir, "sync_info.yaml")
	var indexInfo InfoFile

	// Load existing sync info
	data, err := os.ReadFile(infoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil // No indices synced yet
		}
		return nil, fmt.Errorf("failed to read sync info: %w", err)
	}

	if err := yaml.Unmarshal(data, &indexInfo); err != nil {
		return nil, fmt.Errorf("failed to parse sync info: %w", err)
	}

	indices := make([]string, 0, len(indexInfo.Indices))
	if len(indexInfo.Indices) != 0 {
		ui.Info(fmt.Sprintf("Found %d currently cached indices", len(indexInfo.Indices)))

		// Find the longest index name for proper alignment
		maxNameLength := 0
		for _, info := range indexInfo.Indices {
			if len(info.Name) > maxNameLength {
				maxNameLength = len(info.Name)
			}
		}

		// Print each index info
		for _, info := range indexInfo.Indices {
			ui.Info(fmt.Sprintf("%-*s  Last updated: %-19s  Size: %s",
				maxNameLength,
				info.Name,
				utils.ParseDate(info.LastUpdated),
				utils.GetSizeHuman(info.Size)))
			indices = append(indices, info.Name)
		}

	}

	return indices, nil
}

func IndicesSync(indices []string) error {

	configDir, err := config.IndicesDir()
	if err != nil {
		return err
	}

	infoPath := filepath.Join(configDir, "sync_info.yaml")
	var indexInfo InfoFile

	// Load existing sync info if it exists
	if data, err := os.ReadFile(infoPath); err == nil {
		if err := yaml.Unmarshal(data, &indexInfo); err != nil {
			return fmt.Errorf("failed to parse existing sync info: %w", err)
		}
	}

	if indexInfo.Indices == nil {
		indexInfo.Indices = make([]IndexInfo, 0)
	}

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

		lastUpdated := response.GetData()[0].DateAdded
		date := utils.ParseDate(lastUpdated)

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
		}

		// Update or add sync info for this specific index
		updatedInfo := IndexInfo{
			Name:        index,
			LastSync:    time.Now(),
			Size:        size,
			LastUpdated: lastUpdated,
		}

		found := false
		for i, info := range indexInfo.Indices {
			if info.Name == index {
				indexInfo.Indices[i] = updatedInfo
				found = true
				break
			}
		}
		if !found {
			indexInfo.Indices = append(indexInfo.Indices, updatedInfo)
		}

		ui.Info(fmt.Sprintf("Successfully synced %s (Size: %s)", index, utils.GetSizeHuman(size)))

		// Optionally, remove the downloaded zip file
		if err := os.Remove(filePath); err != nil {
			ui.Error(fmt.Sprintf("Failed to remove downloaded zip file: %s", err))
		}
	}

	// Save updated sync info
	data, err := yaml.Marshal(indexInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal sync info: %w", err)
	}

	if err := os.WriteFile(infoPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write sync info: %w", err)
	}

	return nil
}
