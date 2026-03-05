package db

import (
	"testing"

	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
)

func TestCPESearch(t *testing.T) {
	// Create the test table based on schema
	if err := createTableFromSchema(t, "cpecve"); err != nil {
		t.Fatalf("failed to create table from schema: %v", err)
	}

	// Insert test data with various vendor/product combinations
	testData := []struct {
		vendor  string
		product string
	}{
		{"microsoft", "windows"},
		{"microsoft", "office"},
		{"apache", "tomcat"},
		{"apache", "httpd"},
		{"google", "chrome"},
		{"mozilla", "firefox"},
		{"test_vendor", "test_product"},
	}

	for _, data := range testData {
		_, err := testDB.Exec(`INSERT INTO cpecve (vendor, product, version, "update", edition, language, sw_edition, target_sw, target_hw, other, cpe23Uri, cves) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			data.vendor, data.product, "", "", "", "", "", "", "", "",
			"cpe:/a:"+data.vendor+":"+data.product, `["CVE-2024-0001"]`)
		if err != nil {
			t.Fatalf("failed to insert test data: %v", err)
		}
	}

	// Define test cases
	tests := []struct {
		name       string
		cpe        cpeutils.CPE
		expectHits int
	}{
		{
			name: "Exact match vendor and product",
			cpe: cpeutils.CPE{
				Vendor:  "microsoft",
				Product: "windows",
			},
			expectHits: 1,
		},
		{
			name: "Exact match - case insensitive",
			cpe: cpeutils.CPE{
				Vendor:  "MICROSOFT",
				Product: "WINDOWS",
			},
			expectHits: 1,
		},
		{
			name: "Wildcard at end of vendor",
			cpe: cpeutils.CPE{
				Vendor:  "micro*",
				Product: "windows",
			},
			expectHits: 1,
		},
		{
			name: "Wildcard at beginning of vendor",
			cpe: cpeutils.CPE{
				Vendor:  "*soft",
				Product: "windows",
			},
			expectHits: 1,
		},
		{
			name: "Wildcard at both ends of vendor",
			cpe: cpeutils.CPE{
				Vendor:  "*cros*",
				Product: "windows",
			},
			expectHits: 1,
		},
		{
			name: "Wildcard at end of product",
			cpe: cpeutils.CPE{
				Vendor:  "apache",
				Product: "tom*",
			},
			expectHits: 1,
		},
		{
			name: "Wildcard at beginning of product",
			cpe: cpeutils.CPE{
				Vendor:  "apache",
				Product: "*cat",
			},
			expectHits: 1,
		},
		{
			name: "Both vendor and product with wildcards",
			cpe: cpeutils.CPE{
				Vendor:  "apa*",
				Product: "*cat",
			},
			expectHits: 1,
		},
		{
			name: "Wildcard matches multiple",
			cpe: cpeutils.CPE{
				Vendor:  "apache",
				Product: "*",
			},
			expectHits: 2, // tomcat and httpd
		},
		{
			name: "Vendor wildcard only (asterisk alone)",
			cpe: cpeutils.CPE{
				Vendor:  "*",
				Product: "chrome",
			},
			expectHits: 1,
		},
		{
			name: "No match - exact matching",
			cpe: cpeutils.CPE{
				Vendor:  "micro",
				Product: "windows",
			},
			expectHits: 0, // Should not match because we use strict matching now
		},
		{
			name: "No match - product doesn't exist",
			cpe: cpeutils.CPE{
				Vendor:  "microsoft",
				Product: "azure",
			},
			expectHits: 0,
		},
		{
			name: "Empty vendor with specific product",
			cpe: cpeutils.CPE{
				Vendor:  "",
				Product: "windows",
			},
			expectHits: 1,
		},
		{
			name: "Specific vendor with empty product",
			cpe: cpeutils.CPE{
				Vendor:  "microsoft",
				Product: "",
			},
			expectHits: 2, // windows and office
		},
		{
			name: "Both empty",
			cpe: cpeutils.CPE{
				Vendor:  "",
				Product: "",
			},
			expectHits: 7, // all records
		},
		{
			name: "Special characters in exact match",
			cpe: cpeutils.CPE{
				Vendor:  "test_vendor",
				Product: "test_product",
			},
			expectHits: 1,
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
				for _, r := range results {
					t.Logf("  Found: vendor=%s, product=%s", r.Vendor, r.Product)
				}
			}

			// For fast system execution, stats.Duration might be 0 but stats should exist
			if stats == nil {
				t.Error("expected valid stats object")
			}
		})
	}
}
