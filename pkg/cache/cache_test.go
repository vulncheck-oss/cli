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
	// Setup
	tempDir := t.TempDir()
	err := config.SetIndicesDir(tempDir)
	assert.NoError(t, err)

	// Test when no sync_info.yaml exists
	info, err := Indices()
	assert.NoError(t, err)
	assert.Empty(t, info.Indices)

	// Create a mock sync_info.yaml
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

	// Test reading existing sync_info.yaml
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

// TestIndicesSync is not included as it requires mocking external API calls.
// Consider writing an integration test for IndicesSync or mocking the session.Connect function.
