package bill

import (
	"github.com/anchore/syft/syft/sbom"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveSBOM(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sbom_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

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
	defer os.RemoveAll(tempDir)

	sbomFile := filepath.Join(tempDir, "test_sbom.json")
	mockSBOM := &sbom.SBOM{}

	err = SaveSBOM(mockSBOM, sbomFile)
	if err != nil {
		t.Fatalf("SaveSBOM failed: %v", err)
	}

	loadedSBOM, err := LoadSBOM(sbomFile)
	if err != nil {
		t.Fatalf("LoadSBOM failed: %v", err)
	}

	if loadedSBOM == nil {
		t.Errorf("LoadSBOM returned nil SBOM")
	}
}

func TestGetPURLDetail(t *testing.T) {
	mockSBOM := &sbom.SBOM{}
	// TODO: Populate mockSBOM with test data

	purls := GetPURLDetail(mockSBOM)

	if len(purls) != 0 {
		t.Errorf("Expected 0 PURLs, got %d", len(purls))
	}

	// Test with nil SBOM
	nilPurls := GetPURLDetail(nil)
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
	// TODO: Implement test cases for baseScore function
	// This would require creating mock client.ApiNVD20CVEExtended objects
}

func TestTemporalScore(t *testing.T) {
	// TODO: Implement test cases for temporalScore function
	// This would require creating mock client.ApiNVD20CVEExtended objects
}

func TestGetVulns(t *testing.T) {
	// TODO: Implement test cases for GetVulns function
	// This would require mocking the session.Connect and its methods
}

func TestGetMeta(t *testing.T) {
	// TODO: Implement test cases for GetMeta function
	// This would require mocking the session.Connect and its methods
}
