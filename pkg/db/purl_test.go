package db

import (
	"github.com/package-url/packageurl-go"
	"strings"
	"testing"
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

// Helper function to parse PURLs without error handling in tests
func mustParsePurl(s string) packageurl.PackageURL {
	purl, err := packageurl.FromString(s)
	if err != nil {
		panic("Invalid PURL string: " + s)
	}
	return purl
}
