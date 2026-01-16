package cache

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vulncheck-oss/cli/pkg/config"
	"gopkg.in/yaml.v3"
)

func TestIndices(t *testing.T) {
	originalDir, err := config.IndicesDir()
	assert.NoError(t, err)
	defer func() {
		err := config.SetIndicesDir(originalDir)
		assert.NoError(t, err)
	}()

	tempDir := t.TempDir()
	err = config.SetIndicesDir(tempDir)
	assert.NoError(t, err)

	info, err := Indices()
	assert.NoError(t, err)
	assert.Empty(t, info.Indices)

	mockInfo := InfoFile{
		Indices: []IndexInfo{
			{Name: "test1", LastSync: time.Now(), Size: 1000, LastUpdated: "2023-05-01"},
			{Name: "test2", LastSync: time.Now(), Size: 2000, LastUpdated: "2023-05-02"},
		},
	}
	data, err := yaml.Marshal(mockInfo)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempDir, "sync_info.yaml"), data, 0644)
	assert.NoError(t, err)

	info, err = Indices()
	assert.NoError(t, err)
	assert.Len(t, info.Indices, 2)
	assert.Equal(t, "test1", info.Indices[0].Name)
	assert.Equal(t, "test2", info.Indices[1].Name)
}

func TestInfoFile_IndexExists(t *testing.T) {
	info := InfoFile{
		Indices: []IndexInfo{
			{Name: "test1"},
			{Name: "test2"},
		},
	}

	assert.True(t, info.IndexExists("test1"))
	assert.True(t, info.IndexExists("test2"))
	assert.False(t, info.IndexExists("test3"))
}

// TestIndicesSync is not included as it requires mocking external API calls.
// Consider writing an integration test for IndicesSync or mocking the session.Connect function.

func TestInfoFile_GetIndex(t *testing.T) {
	info := InfoFile{
		Indices: []IndexInfo{
			{Name: "test1", Size: 1000},
			{Name: "test2", Size: 2000},
		},
	}

	index := info.GetIndex("test1")
	assert.NotNil(t, index)
	assert.Equal(t, "test1", index.Name)
	assert.Equal(t, uint64(1000), index.Size)

	index = info.GetIndex("test2")
	assert.NotNil(t, index)
	assert.Equal(t, "test2", index.Name)
	assert.Equal(t, uint64(2000), index.Size)

	index = info.GetIndex("test3")
	assert.Nil(t, index)
}

func TestTaskDownloadDirectoryCreation(t *testing.T) {
	tempDir := t.TempDir()

	// Test nested directory path creation
	nestedPath := filepath.Join(tempDir, "deeply", "nested", "path", "file.zip")

	// Simulate the directory creation logic from taskDownload (lines 50-54)
	err := os.MkdirAll(filepath.Dir(nestedPath), 0755)
	assert.NoError(t, err, "should create nested directories without error")

	// Verify the directory structure was created
	dirPath := filepath.Dir(nestedPath)
	info, err := os.Stat(dirPath)
	assert.NoError(t, err, "directory should exist")
	assert.True(t, info.IsDir(), "should be a directory")
	assert.Equal(t, os.FileMode(0755), info.Mode().Perm(), "should have correct permissions")

	// Test that creating a file works after directory creation
	file, err := os.Create(nestedPath)
	assert.NoError(t, err, "should be able to create file in nested path")
	if file != nil {
		err = file.Close()
		assert.NoError(t, err, "should close file without error")
	}

	// Verify file was created
	_, err = os.Stat(nestedPath)
	assert.NoError(t, err, "file should exist after creation")
}

func TestTaskDownloadDirectoryCreationWithExistingPath(t *testing.T) {
	tempDir := t.TempDir()

	// Create a directory path
	existingPath := filepath.Join(tempDir, "existing", "path")
	err := os.MkdirAll(existingPath, 0755)
	assert.NoError(t, err)

	// Test that MkdirAll is idempotent (lines 50-54 in tasks.go)
	filePath := filepath.Join(existingPath, "file.zip")
	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	assert.NoError(t, err, "MkdirAll should succeed on existing directory")

	// Verify directory still exists with correct permissions
	info, err := os.Stat(existingPath)
	assert.NoError(t, err)
	assert.True(t, info.IsDir())
}
