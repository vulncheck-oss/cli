package db

import (
	"strings"
	"testing"

	"github.com/package-url/packageurl-go"
)

func setupPurlTable(t *testing.T) {
	// Get the schema for npm (a package manager index from the "purl PM" schema)
	schema := GetSchema("npm")
	if schema == nil {
		t.Fatalf("failed to get schema for npm")
	}

	// Create the test table based on schema
	var columns []string
	for _, col := range schema.Columns {
		colDef := "\"" + col.Name + "\" " + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		columns = append(columns, colDef)
	}

	// Use npm as the test table name (replacing - with _)
	tableDef := "CREATE TABLE IF NOT EXISTS npm (" + strings.Join(columns, ", ") + ")"
	_, err := testDB.Exec(tableDef)
	if err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}

	// testDB is shared across the whole package - clear prior rows so this
	// helper is idempotent regardless of which test ran first.
	if _, err := testDB.Exec(`DELETE FROM npm`); err != nil {
		t.Fatalf("failed to clear npm table: %v", err)
	}

	// Insert test data with all required values
	_, err = testDB.Exec(`INSERT INTO npm (
		name, version, purl, licenses, cves, vulnerabilities
	) VALUES (?, ?, ?, ?, ?, ?)`,
		"express", "4.17.1",
		`["pkg:npm/express@4.17.1"]`, // purl json array
		`["MIT"]`,                    // licenses json array
		`["CVE-2022-12345"]`,         // cves json array
		`[{"detection":"CVE-2022-12345","fixed_version":"4.18.0"}]`, // vulnerabilities json array
	)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Insert a second test record with a different version that's already fixed
	_, err = testDB.Exec(`INSERT INTO npm (
		name, version, purl, licenses, cves, vulnerabilities
	) VALUES (?, ?, ?, ?, ?, ?)`,
		"express", "4.18.0",
		`["pkg:npm/express@4.18.0"]`, // purl json array
		`["MIT"]`,                    // licenses json array
		`[]`,                         // No CVEs for the fixed version
		`[]`,                         // No vulnerabilities for the fixed version
	)
	if err != nil {
		t.Fatalf("failed to insert second test data: %v", err)
	}
}

// setupPurlAmbiguousVersions inserts two rows whose stored PURLs share a
// version prefix - exercising the substring LIKE risk in purlSearchExact.
func setupPurlAmbiguousVersions(t *testing.T) {
	if _, err := testDB.Exec(`INSERT INTO npm (name, version, purl, licenses, cves, vulnerabilities) VALUES (?,?,?,?,?,?)`,
		"foo", "1.0.1",
		`["pkg:npm/foo@1.0.1"]`,
		`[]`,
		`["CVE-FOO-101"]`,
		`[{"detection":"CVE-FOO-101","fixed_version":"2.0.0"}]`,
	); err != nil {
		t.Fatalf("insert foo@1.0.1: %v", err)
	}
	if _, err := testDB.Exec(`INSERT INTO npm (name, version, purl, licenses, cves, vulnerabilities) VALUES (?,?,?,?,?,?)`,
		"foo", "1.0.10",
		`["pkg:npm/foo@1.0.10"]`,
		`[]`,
		`["CVE-FOO-1010"]`,
		`[{"detection":"CVE-FOO-1010","fixed_version":"2.0.0"}]`,
	); err != nil {
		t.Fatalf("insert foo@1.0.10: %v", err)
	}
}

func TestPURLSearchExactNoVersionPrefixLeak(t *testing.T) {
	setupPurlTable(t)
	setupPurlAmbiguousVersions(t)

	// Searching for 1.0.1 must not bleed into the 1.0.10 row even though "1.0.1"
	// is a substring of "1.0.10".
	results, _, err := PURLSearch("npm", mustParsePurl("pkg:npm/foo@1.0.1"))
	if err != nil {
		t.Fatalf("PURLSearch failed: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected exactly 1 row for foo@1.0.1, got %d: %+v", len(results), results)
	}
	if results[0].Version != "1.0.1" {
		t.Errorf("expected version 1.0.1, got %s", results[0].Version)
	}
	if len(results[0].CVEs) != 1 || results[0].CVEs[0] != "CVE-FOO-101" {
		t.Errorf("expected only CVE-FOO-101, got %v", results[0].CVEs)
	}
}

func TestPURLSearch(t *testing.T) {
	// Setup the test table
	setupPurlTable(t)

	// Define test cases
	tests := []struct {
		name       string
		purl       packageurl.PackageURL
		expectHits int
	}{
		{
			name:       "Match exact purl",
			purl:       mustParsePurl("pkg:npm/express@4.17.1"),
			expectHits: 1, // Should find the vulnerable version
		},
		{
			name:       "Match package only",
			purl:       mustParsePurl("pkg:npm/express"),
			expectHits: 1, // Should find only the vulnerable version due to filter
		},
		{
			name:       "No match for different package",
			purl:       mustParsePurl("pkg:npm/nonexistent"),
			expectHits: 0,
		},
		{
			name:       "No match for fixed version",
			purl:       mustParsePurl("pkg:npm/express@4.18.0"),
			expectHits: 0, // The fixed version should be filtered out
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, stats, err := PURLSearch("npm", tt.purl)
			if err != nil {
				t.Fatalf("PURLSearch failed: %v", err)
			}

			if len(results) != tt.expectHits {
				t.Errorf("expected %d hits, got %d", tt.expectHits, len(results))
			}

			// For fast system execution, stats.Duration might be 0 but stats should exist
			if stats == nil {
				t.Error("expected valid stats object")
			}

			// For matching results, verify the vulnerability data is correct
			if len(results) > 0 {
				result := results[0]

				// Verify package details
				if result.Name != "express" {
					t.Errorf("expected name 'express', got '%s'", result.Name)
				}

				if result.Version != "4.17.1" {
					t.Errorf("expected version '4.17.1', got '%s'", result.Version)
				}

				// Verify PURL
				if len(result.Purl) != 1 || result.Purl[0] != "pkg:npm/express@4.17.1" {
					t.Errorf("invalid purl: %v", result.Purl)
				}

				// Verify CVEs
				if len(result.CVEs) != 1 || result.CVEs[0] != "CVE-2022-12345" {
					t.Errorf("invalid CVEs: %v", result.CVEs)
				}

				// Verify vulnerabilities - using the correct field names from PurlVulnerability struct
				if len(result.Vulnerabilities) != 1 ||
					result.Vulnerabilities[0].Detection != "CVE-2022-12345" ||
					result.Vulnerabilities[0].FixedVersion != "4.18.0" {
					t.Errorf("invalid vulnerabilities: %v", result.Vulnerabilities)
				}
			}
		})
	}
}

// setupAlpinePurlsTable builds an alpine-purls table that mimics the index
func setupAlpinePurlsTable(t *testing.T) {
	schema := GetSchema("alpine-purls")
	if schema == nil {
		t.Fatalf("failed to get schema for alpine-purls")
	}

	var columns []string
	for _, col := range schema.Columns {
		colDef := "\"" + col.Name + "\" " + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		columns = append(columns, colDef)
	}
	tableDef := "CREATE TABLE IF NOT EXISTS alpine_purls (" + strings.Join(columns, ", ") + ")"
	if _, err := testDB.Exec(tableDef); err != nil {
		t.Fatalf("failed to create alpine_purls table: %v", err)
	}

	rows := []struct {
		purl  string
		cves  string
		vulns string
	}{
		// Includes a still-open CVE, an already-fixed CVE, and a fix
		// with an Alpine "_git" build suffix
		{
			`["pkg:apk/alpine/musl@1.1.0-3.0"]`,
			`["CVE-OPEN","CVE-PATCHED","CVE-GIT"]`,
			`[{"detection":"CVE-OPEN","fixed_version":"1.1.22-r3"},` +
				`{"detection":"CVE-PATCHED","fixed_version":"1.1.22-r1"},` +
				`{"detection":"CVE-GIT","fixed_version":"1.2.4_git20230717-r6"}]`,
		},
		// reports a bogus fix for CVE-PATCHED.
		{
			`["pkg:apk/alpine/musl@1.1.22-3.10"]`,
			`["CVE-PATCHED"]`,
			`[{"detection":"CVE-PATCHED","fixed_version":"1.1.22-r2"}]`,
		},
		// A different package that must not leak into musl results
		{
			`["pkg:apk/alpine/zlib@1.2.11-3.10"]`,
			`["CVE-ZLIB"]`,
			`[{"detection":"CVE-ZLIB","fixed_version":"1.2.12-r0"}]`,
		},
	}
	for _, r := range rows {
		if _, err := testDB.Exec(
			`INSERT INTO alpine_purls (name, version, purl, licenses, cves, vulnerabilities) VALUES (?,?,?,?,?,?)`,
			"musl", "1.1.22", r.purl, `[]`, r.cves, r.vulns,
		); err != nil {
			t.Fatalf("failed to insert alpine test data: %v", err)
		}
	}
}

func TestPURLSearchOS(t *testing.T) {
	setupAlpinePurlsTable(t)

	// Installed musl 1.1.22-r2.
	purl := mustParsePurl("pkg:apk/alpine/musl@1.1.22-r2?arch=x86_64&distro=alpine-3.10.0")

	results, stats, err := PURLSearch("alpine-purls", purl)
	if err != nil {
		t.Fatalf("PURLSearch failed: %v", err)
	}
	if stats == nil {
		t.Error("expected valid stats object")
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 PurlEntry, got %d", len(results))
	}

	got := map[string]struct{}{}
	for _, v := range results[0].Vulnerabilities {
		got[v.Detection] = struct{}{}
	}

	// CVE-OPEN (fixed 1.1.22-r3 > installed r2): affected, only in the complete row.
	// CVE-GIT (fixed 1.2.4_git...): affected, requires "_git" suffix to parse.
	// CVE-PATCHED (fixed 1.1.22-r1 <= installed r2): not affected.
	wantPresent := []string{"CVE-OPEN", "CVE-GIT"}
	wantAbsent := []string{"CVE-PATCHED", "CVE-ZLIB"}

	for _, cve := range wantPresent {
		if _, ok := got[cve]; !ok {
			t.Errorf("expected %s to be reported, got %v", cve, got)
		}
	}
	for _, cve := range wantAbsent {
		if _, ok := got[cve]; ok {
			t.Errorf("expected %s to be filtered out, got %v", cve, got)
		}
	}

	// Results should be reported against the installed version, like the API.
	if results[0].Version != "1.1.22-r2" {
		t.Errorf("expected reported version 1.1.22-r2, got %s", results[0].Version)
	}
}

func TestIsAffectedVersion(t *testing.T) {
	tests := []struct {
		name           string
		current, fixed string
		want           bool
	}{
		{"apk older revision", "1.30.1-r2", "1.30.1-r5", true},
		{"apk equal revision", "1.30.1-r2", "1.30.1-r2", false},
		{"apk newer revision", "1.2.11-r1", "1.2.11-r0", false},
		{"apk _git build suffix", "1.1.22-r2", "1.2.4_git20230717-r6", true},
		{"apk newer upstream fix", "2.10.4-r1", "2.12.0-r0", true},

		// RPM-style release tags - underscores must be stripped or the parser rejects them.
		// LIMITATION: stripping is lossy, so RPM errata bumps within the same upstream
		// version (1.0.0_4 vs 1.0.0_5, .el8_3 vs .el8_5) collapse to equal and are
		// treated as "not affected". This matches the upstream API and is left as-is
		// here; lifting it requires keeping the release suffix and teaching the
		// comparator about it.
		{"rpm _N release bumps collapse to equal", "1.0.0_4", "1.0.0_5", false},
		{"rpm .elN_N release bumps collapse to equal", "1.0.0.el8_3", "1.0.0.el8_5", false},
		{"rpm .elN newer than no-suffix", "1.0.0", "1.0.0.el8", true},

		// Linux-kernel style versions with embedded underscores - parser would reject without the strip.
		{"kernel +gitAUTOINC fixed", "5.15.201+gitAUTOINC+a_b", "5.15.202", true},
		{"kernel +gitAUTOINC current matches", "5.15.201+gitAUTOINC+a_b", "5.15.201+gitAUTOINC+a_b", false},

		// Unparseable inputs return false rather than panicking - protects the offline scan loop
		// from blowing up on malformed advisories.
		{"garbage current", "not a version", "1.0.0", false},
		{"garbage fixed", "1.0.0", "not a version", false},
		{"empty fixed", "1.0.0", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAffectedVersion(tt.current, tt.fixed); got != tt.want {
				t.Errorf("isAffectedVersion(%q, %q) = %v; want %v", tt.current, tt.fixed, got, tt.want)
			}
		})
	}
}

// Helper function to parse PURLs without error handling in tests
func mustParsePurl(s string) packageurl.PackageURL {
	purl, err := packageurl.FromString(s)
	if err != nil {
		panic("Invalid PURL string: " + s)
	}
	return purl
}
