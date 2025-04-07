package db

import (
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"testing"
)

func TestCPESearch(t *testing.T) {
	// Create the test table based on schema
	if err := createTableFromSchema(t, "cpecve"); err != nil {
		t.Fatalf("failed to create table from schema: %v", err)
	}

	// Insert test data with values for all columns and correct CVEs format (array of strings)
	_, err := testDB.Exec(`INSERT INTO cpecve (vendor, product, version, "update", edition, language, sw_edition, target_sw, target_hw, other, cpe23Uri, cves) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		"test_vendor", "test_product", "", "", "", "", "", "", "", "",
		"cpe:/a:test_vendor:test_product", `["CVE-1234-5678"]`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Define test cases
	tests := []struct {
		name       string
		cpe        cpeutils.CPE
		expectHits int
	}{
		{
			name: "Match vendor and product",
			cpe: cpeutils.CPE{
				Vendor:  "test_vendor",
				Product: "test_product",
			},
			expectHits: 1,
		},
		{
			name: "No match",
			cpe: cpeutils.CPE{
				Vendor:  "nonexistent_vendor",
				Product: "nonexistent_product",
			},
			expectHits: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, stats, err := CPESearch("cpecve", tt.cpe)
			if err != nil {
				t.Fatalf("CPESearch failed: %v", err)
			}

			if len(results) != tt.expectHits {
				t.Errorf("expected %d hits, got %d", tt.expectHits, len(results))
			}

			// For fast system execution, stats.Duration might be 0 but stats should exist
			if stats == nil {
				t.Error("expected valid stats object")
			}
		})
	}
}
