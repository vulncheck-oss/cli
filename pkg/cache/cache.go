package cache

import (
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type IndexInfo struct {
	Name        string    `yaml:"name"`
	LastSync    time.Time `yaml:"last_sync"`
	Size        uint64    `yaml:"size"`         // Size in bytes
	LastUpdated string    `yaml:"last_updated"` // From the server
}

type InfoFile struct {
	Indices []IndexInfo `yaml:"indices"`
}

func Indices() (InfoFile, error) {
	configDir, err := config.IndicesDir()
	if err != nil {
		return InfoFile{}, err
	}
	infoPath := filepath.Join(configDir, "sync_info.yaml")
	var indexInfo InfoFile
	// Load existing sync info
	data, err := os.ReadFile(infoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return InfoFile{}, nil // No indices synced yet
		}
		return InfoFile{}, fmt.Errorf("failed to read sync info: %w", err)
	}

	if err := yaml.Unmarshal(data, &indexInfo); err != nil {
		return InfoFile{}, fmt.Errorf("failed to parse sync info: %w", err)
	}

	return indexInfo, nil
}

func (i *InfoFile) IndexExists(name string) bool {
	for _, index := range i.Indices {
		if index.Name == name {
			return true
		}
	}
	return false
}

func (i *InfoFile) GetIndex(name string) *IndexInfo {
	for _, index := range i.Indices {
		if index.Name == name {
			return &index
		}
	}
	return nil
}

func IndicesSync(indices []string) error {
	configDir, err := config.IndicesDir()
	if err != nil {
		return err
	}

	infoPath := filepath.Join(configDir, "sync_info.yaml")
	indexInfo, err := Indices()
	if err != nil {
		return fmt.Errorf("failed to get cached indices: %w", err)
	}

	// Remove indices not in the provided list
	for i := len(indexInfo.Indices) - 1; i >= 0; i-- {
		if !slices.Contains(indices, indexInfo.Indices[i].Name) {
			// Remove the index folder
			indexDir := filepath.Join(configDir, indexInfo.Indices[i].Name)
			if err := os.RemoveAll(indexDir); err != nil {
				return fmt.Errorf("failed to remove index directory: %w", err)
			}
			// Remove the index from the list
			indexInfo.Indices = append(indexInfo.Indices[:i], indexInfo.Indices[i+1:]...)
		}
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
			_ = ui.Error(fmt.Sprintf("Failed to calculate size of index directory: %s", err))
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
			_ = ui.Error(fmt.Sprintf("Failed to remove downloaded zip file: %s", err))
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
