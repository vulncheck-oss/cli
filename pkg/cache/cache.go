package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fumeapp/taskin"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
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

func syncSingleIndex(index string, configDir string, indexInfo *InfoFile, force bool) taskin.Tasks {
	response, err := session.Connect(config.Token()).GetIndexBackup(index)
	if err != nil {
		return taskin.Tasks{
			{
				Title: fmt.Sprintf("Error syncing index: %s", index),
				Task: func(t *taskin.Task) error {
					return err
				},
			},
		}
	}

	if len(response.GetData()) == 0 {
		return taskin.Tasks{
			{
				Title: fmt.Sprintf("No data received for index: %s", index),
				Task: func(t *taskin.Task) error {
					return fmt.Errorf("no data received for index %s", index)
				},
			},
		}
	}

	file, err := utils.ExtractFileBasename(response.GetData()[0].URL)
	if err != nil {
		return taskin.Tasks{
			{
				Title: fmt.Sprintf("Error extracting file for index: %s", index),
				Task: func(t *taskin.Task) error {
					return err
				},
			},
		}
	}

	filePath := filepath.Join(configDir, file)
	lastUpdated := response.GetData()[0].DateAdded

	if indexInfo.IndexExists(index) && indexInfo.GetIndex(index).LastUpdated == lastUpdated && !force {
		return taskin.Tasks{
			{
				Title: fmt.Sprintf("Index %s is already up to date", index),
				Task: func(t *taskin.Task) error {
					return nil
				},
			},
		}
	}

	childTasks := taskin.Tasks{
		taskDownload(index, filePath),
		taskExtract(index, configDir, filePath),
		taskDB(index, configDir, filePath, lastUpdated, indexInfo),
	}

	return childTasks
}

func IndicesSync(indices []string, force bool) error {
	configDir, err := config.IndicesDir()
	if err != nil {
		return err
	}

	infoPath := filepath.Join(configDir, "sync_info.yaml")
	indexInfo, err := Indices()
	if err != nil {
		return fmt.Errorf("failed to get cached indices: %w", err)
	}

	for i := len(indexInfo.Indices) - 1; i >= 0; i-- {
		if !slices.Contains(indices, indexInfo.Indices[i].Name) {
			indexDir := filepath.Join(configDir, indexInfo.Indices[i].Name)
			if err := os.RemoveAll(indexDir); err != nil {
				return fmt.Errorf("failed to remove index directory: %w", err)
			}
			indexInfo.Indices = append(indexInfo.Indices[:i], indexInfo.Indices[i+1:]...)
		}
	}

	// If there are indices to sync, run the sync tasks
	if len(indices) > 0 {
		tasks := taskin.Tasks{}

		for _, index := range indices {
			idx := index
			parentTask := taskin.Task{
				Title: fmt.Sprintf("Sync index %s", idx),
				Task: func(t *taskin.Task) error {
					t.Title = fmt.Sprintf("Syncing index %s", idx)
					return nil
				},
				Tasks: syncSingleIndex(idx, configDir, &indexInfo, force),
			}

			tasks = append(tasks, parentTask)
		}

		runner := taskin.New(tasks, taskin.Defaults)
		if err := runner.Run(); err != nil {
			return err
		}
	}

	data, err := yaml.Marshal(indexInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal sync info: %w", err)
	}

	if err := os.WriteFile(infoPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write sync info: %w", err)
	}

	return nil
}

func PurgeIndices() error {
	configDir, err := config.IndicesDir()
	if err != nil {
		return fmt.Errorf("failed to get indices directory: %w", err)
	}

	// Remove the entire indices directory
	if err := os.RemoveAll(configDir); err != nil {
		return fmt.Errorf("failed to remove indices directory: %w", err)
	}

	// Recreate an empty indices directory
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to recreate indices directory: %w", err)
	}

	return nil
}
