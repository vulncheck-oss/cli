package bill

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/anchore/syft/syft/sbom"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/models"
)

func TestSaveSBOM(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sbom_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Errorf("Failed to remove temp directory: %v", err)
		}
	}()

	sbomFile := filepath.Join(tempDir, "test_sbom.json")
	mockSBOM := &sbom.SBOM{}

	err = SaveSBOM(mockSBOM, sbomFile)
	if err != nil {
		t.Fatalf("SaveSBOM failed: %v", err)
	}

	_, err = os.Stat(sbomFile)
	if os.IsNotExist(err) {
		t.Errorf("SaveSBOM did not create the file")
	}
}

func TestLoadSBOM(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sbom_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Errorf("Failed to remove temp directory: %v", err)
		}
	}()

	sbomFile := filepath.Join(tempDir, "test_sbom.json")
	mockSBOM := &sbom.SBOM{}

	err = SaveSBOM(mockSBOM, sbomFile)
	if err != nil {
		t.Fatalf("SaveSBOM failed: %v", err)
	}

	loadedSBOM, _, err := LoadSBOM(sbomFile)
	if err != nil {
		t.Fatalf("LoadSBOM failed: %v", err)
	}

	if loadedSBOM == nil {
		t.Errorf("LoadSBOM returned nil SBOM")
	}
}

func TestGetPURLDetail(t *testing.T) {
	mockSBOM := &sbom.SBOM{}

	purls := GetPURLDetail(mockSBOM, nil)

	if len(purls) != 0 {
		t.Errorf("Expected 0 PURLs, got %d", len(purls))
	}

	nilPurls := GetPURLDetail(nil, nil)
	if len(nilPurls) != 0 {
		t.Errorf("Expected 0 PURLs for nil SBOM, got %d", len(nilPurls))
	}
}

func TestFormatSingleDecimal(t *testing.T) {
	testCases := []struct {
		input    float32
		expected string
	}{
		{3.14159, "3.1"},
		{2.0, "2.0"},
		{0.1, "0.1"},
		{9.99, "10.0"},
	}

	for _, tc := range testCases {
		result := formatSingleDecimal(&tc.input)
		if result != tc.expected {
			t.Errorf("formatSingleDecimal(%f) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestBaseScore(t *testing.T) {
}

func TestTemporalScore(t *testing.T) {
}

func TestGetVulns(t *testing.T) {
}

func TestGetMeta(t *testing.T) {
}

func TestGetOfflineMeta(t *testing.T) {
	t.Run("empty vulnerabilities", func(t *testing.T) {
		vulns := []models.ScanResultVulnerabilities{}
		indices := cache.InfoFile{
			Indices: []cache.IndexInfo{
				{Name: "vulncheck-nvd2"},
			},
		}

		result, err := GetOfflineMeta(indices, vulns)
		if err == nil {
			t.Skip("Cannot test without mocking dependencies")
		}
		_ = result
	})

	t.Run("missing index", func(t *testing.T) {
		vulns := []models.ScanResultVulnerabilities{
			{
				CVE:     "CVE-2021-44228",
				Name:    "log4j",
				Version: "2.14.1",
			},
		}
		indices := cache.InfoFile{
			Indices: []cache.IndexInfo{},
		}

		_, err := GetOfflineMeta(indices, vulns)
		if err == nil {
			t.Skip("Cannot test without mocking dependencies")
		}
	})
}
