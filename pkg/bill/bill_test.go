package bill

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/anchore/syft/syft"
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

// TestLoadSBOM_CycloneDX17 documents that scanning against a CycloneDX 1.7
// SBOM (the spec version syft >= 1.46.0 defaults to) must succeed. Syft
// versions before 1.46.0 don't recognize specVersion "1.7" and fail to
// decode it, breaking `scan --sbom-input-file` for these inputs.
func TestLoadSBOM_CycloneDX17(t *testing.T) {
	loadedSBOM, _, err := LoadSBOM(filepath.Join("testdata", "cyclonedx-1.7.json"))
	if err != nil {
		t.Fatalf("LoadSBOM failed to decode a CycloneDX 1.7 SBOM: %v", err)
	}

	if loadedSBOM == nil {
		t.Fatal("LoadSBOM returned nil SBOM")
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

// TestBuildSBOMConfig documents the mapping between our --enrich directive
// slice and the pkgcataloging.Config fields we hand to syft. Locks in the
// expected per-scope behaviour so a future refactor can't quietly stop
// flipping (or start over-flipping) fields.
func TestBuildSBOMConfig(t *testing.T) {
	t.Run("empty Enrich returns nil (preserves syft defaults)", func(t *testing.T) {
		if cfg := buildSBOMConfig(SBOMOptions{}); cfg != nil {
			t.Fatalf("expected nil config for empty Enrich, got %+v", cfg)
		}
	})

	t.Run("Enrich=all flips every scope", func(t *testing.T) {
		cfg := buildSBOMConfig(SBOMOptions{Enrich: []string{"all"}})
		if cfg == nil {
			t.Fatal("expected non-nil config")
		}
		if !cfg.Packages.Golang.SearchRemoteLicenses {
			t.Error("golang.SearchRemoteLicenses not enabled under all")
		}
		if !cfg.Packages.Golang.SearchLocalModCacheLicenses {
			t.Error("golang.SearchLocalModCacheLicenses not enabled under all")
		}
		if !cfg.Packages.Golang.SearchLocalVendorLicenses {
			t.Error("golang.SearchLocalVendorLicenses not enabled under all")
		}
		if !cfg.Packages.JavaScript.SearchRemoteLicenses {
			t.Error("javascript.SearchRemoteLicenses not enabled under all")
		}
		if !cfg.Packages.Python.SearchRemoteLicenses {
			t.Error("python.SearchRemoteLicenses not enabled under all")
		}
		if !cfg.Packages.Python.GuessUnpinnedRequirements {
			t.Error("python.GuessUnpinnedRequirements not enabled under all")
		}
		if !cfg.Packages.JavaArchive.UseNetwork {
			t.Error("java.UseNetwork not enabled under all")
		}
		if !cfg.Packages.JavaArchive.UseMavenLocalRepository {
			t.Error("java.UseMavenLocalRepository not enabled under all")
		}
	})

	t.Run("single scope only flips that scope", func(t *testing.T) {
		cfg := buildSBOMConfig(SBOMOptions{Enrich: []string{"golang"}})
		if !cfg.Packages.Golang.SearchRemoteLicenses {
			t.Error("golang enrichment not enabled")
		}
		if cfg.Packages.JavaScript.SearchRemoteLicenses {
			t.Error("javascript enrichment leaked in")
		}
		if cfg.Packages.Python.SearchRemoteLicenses {
			t.Error("python enrichment leaked in")
		}
		if cfg.Packages.JavaArchive.UseNetwork {
			t.Error("java enrichment leaked in")
		}
	})

	t.Run("all with per-scope negation excludes only that scope", func(t *testing.T) {
		cfg := buildSBOMConfig(SBOMOptions{Enrich: []string{"all", "-java"}})
		if !cfg.Packages.Golang.SearchRemoteLicenses {
			t.Error("golang should still be enabled")
		}
		if cfg.Packages.JavaArchive.UseNetwork {
			t.Error("java should be excluded by -java")
		}
	})

	t.Run("none disables everything", func(t *testing.T) {
		cfg := buildSBOMConfig(SBOMOptions{Enrich: []string{"none"}})
		if cfg == nil {
			t.Fatal("expected non-nil config even when none directive is present")
		}
		if cfg.Packages.Golang.SearchRemoteLicenses {
			t.Error("golang should be disabled under none")
		}
		if cfg.Packages.JavaScript.SearchRemoteLicenses {
			t.Error("javascript should be disabled under none")
		}
	})

	// Syft accepts several scope aliases per language (internal/task/package_tasks.go):
	// go/golang, javascript/node/npm, java/maven. Users typing any of them should
	// get the same behaviour as with the publicised name.
	aliases := []struct {
		name  string
		alias string
		check func(cfg *syft.CreateSBOMConfig) bool
	}{
		{"go alias enables golang", "go", func(c *syft.CreateSBOMConfig) bool { return c.Packages.Golang.SearchRemoteLicenses }},
		{"node alias enables javascript", "node", func(c *syft.CreateSBOMConfig) bool { return c.Packages.JavaScript.SearchRemoteLicenses }},
		{"npm alias enables javascript", "npm", func(c *syft.CreateSBOMConfig) bool { return c.Packages.JavaScript.SearchRemoteLicenses }},
		{"maven alias enables java", "maven", func(c *syft.CreateSBOMConfig) bool { return c.Packages.JavaArchive.UseNetwork }},
	}
	for _, a := range aliases {
		t.Run(a.name, func(t *testing.T) {
			cfg := buildSBOMConfig(SBOMOptions{Enrich: []string{a.alias}})
			if cfg == nil || !a.check(cfg) {
				t.Errorf("--enrich %s did not enable the language it aliases", a.alias)
			}
		})
	}

	t.Run("golang enrichment sets UsePackagesLib", func(t *testing.T) {
		cfg := buildSBOMConfig(SBOMOptions{Enrich: []string{"golang"}})
		if !cfg.Packages.Golang.UsePackagesLib {
			t.Error("golang enrichment did not set UsePackagesLib (matches syft's own flip)")
		}
	})
}

func TestEnrichmentScope(t *testing.T) {
	cases := []struct {
		name       string
		directives []string
		scope      string
		want       bool
	}{
		{"empty directives", nil, "golang", false},
		{"bare scope match", []string{"golang"}, "golang", true},
		{"plus scope match", []string{"+golang"}, "golang", true},
		{"negated scope match", []string{"-golang"}, "golang", false},
		{"all enables scope", []string{"all"}, "golang", true},
		{"none disables scope", []string{"none"}, "golang", false},
		{"all beats none absent", []string{"all"}, "python", true},
		{"explicit negation beats all", []string{"all", "-python"}, "python", false},
		{"explicit include beats none", []string{"none", "+python"}, "python", true},
		{"unrelated scope stays off", []string{"golang"}, "python", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := enrichmentScope(tc.directives, tc.scope); got != tc.want {
				t.Errorf("enrichmentScope(%v, %q) = %v, want %v", tc.directives, tc.scope, got, tc.want)
			}
		})
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
