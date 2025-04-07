package db

import (
	"database/sql"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Set the TEST_ENV environment variable for tests
	os.Setenv("TEST_ENV", "true")
	os.Exit(m.Run())
}

func createTableFromSchema(db *sql.DB, schema *Schema) error {
	var columns []string
	for _, col := range schema.Columns {
		colDef := "\"" + col.Name + "\" " + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		columns = append(columns, colDef)
	}

	tableDef := "CREATE TABLE IF NOT EXISTS " + schema.Name + " (" + strings.Join(columns, ", ") + ")"
	_, err := db.Exec(tableDef)
	return err
}

func TestCPESearch(t *testing.T) {
	db, err := DB() // Use the test DB instance explicitly
	if err != nil {
		t.Fatalf("failed to initialize test database: %v", err)
	}
	defer db.Close()

	schema := GetSchema("cpecve")
	if schema == nil {
		t.Fatalf("failed to get schema for cpecve")
	}

	if err := createTableFromSchema(db, schema); err != nil {
		t.Fatalf("failed to create table from schema: %v", err)
	}

	// Insert test data with values for all columns and correct CVEs format (array of strings)
	_, err = db.Exec(`INSERT INTO cpecve (vendor, product, version, "update", edition, language, sw_edition, target_sw, target_hw, other, cpe23Uri, cves) 
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

			if stats == nil || stats.Duration <= 0 {
				t.Error("expected valid stats with non-zero duration")
			}
		})
	}
}
