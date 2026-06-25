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

func TestGetCPEDetail(t *testing.T) {
	// nil SBOM with no refs yields nothing.
	if got := GetCPEDetail(nil, nil); len(got) != 0 {
		t.Errorf("expected 0 CPEs for nil SBOM/refs, got %d", len(got))
	}

	// CPEs declared on CycloneDX "file" components are not surfaced by Syft as
	// packages, so they reach us only through the raw inputRefs.
	refs := []InputSbomRef{
		{SbomRef: "ref-1", PURL: "pkg:generic/qnx_software_development_platform@7.1",
			CPE: "cpe:2.3:a:blackberry:qnx_software_development_platform:7.1:*:*:*:*:*:*:*"},
		{SbomRef: "ref-2", PURL: "pkg:generic/qnx_software_development_platform@7.1",
			CPE: "cpe:2.3:a:blackberry:qnx_software_development_platform:7.1:*:*:*:*:*:*:*"}, // duplicate
		{SbomRef: "ref-3", PURL: "pkg:generic/other@1.0", CPE: ""},                          // no CPE
	}

	got := GetCPEDetail(nil, refs)

	if len(got) != 1 {
		t.Fatalf("expected 1 deduplicated CPE, got %d: %v", len(got), got)
	}
	want := "cpe:2.3:a:blackberry:qnx_software_development_platform:7.1:*:*:*:*:*:*:*"
	if got[0] != want {
		t.Errorf("expected %q, got %q", want, got[0])
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

		result, _, err := GetOfflineMeta(indices, vulns, false)
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

		_, _, err := GetOfflineMeta(indices, vulns, false)
		if err == nil {
			t.Skip("Cannot test without mocking dependencies")
		}
	})

	// Regression: with --warn-on-index, a missing nvd2 index must not drop
	// the already-found vulns. They come back unchanged, ok=false signals to
	// the caller that score columns should be hidden and a hint shown.
	t.Run("missing index with warnOnly preserves vulns", func(t *testing.T) {
		vulns := []models.ScanResultVulnerabilities{
			{CVE: "CVE-2021-44228", Name: "log4j", Version: "2.14.1"},
			{CVE: "CVE-2022-12345", Name: "express", Version: "4.17.1"},
		}
		indices := cache.InfoFile{Indices: []cache.IndexInfo{}}

		out, ok, err := GetOfflineMeta(indices, vulns, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected ok=false when nvd2 index is missing")
		}
		if len(out) != len(vulns) {
			t.Fatalf("expected %d vulns to pass through, got %d", len(vulns), len(out))
		}
		for i := range vulns {
			if out[i].CVE != vulns[i].CVE {
				t.Errorf("vuln %d: got CVE %s, want %s", i, out[i].CVE, vulns[i].CVE)
			}
		}
	})
}
