package db

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"testing"
)

func TestCPESearch(t *testing.T) {
	// Save original DBFunc function to restore later
	originalDBFunc := DBFunc
	defer func() {
		DBFunc = originalDBFunc
	}()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	// Override DBFunc function
	DBFunc = func() (*sql.DB, error) {
		return mockDB, nil
	}

	indexName := "test-index"
	// One condition
	input := cpeutils.CPE{Vendor: "testvendor"}
	tableName := "test_index"

	// Expect a query with LIKE '%testvendor%'
	mock.ExpectQuery(fmt.Sprintf(`SELECT .* FROM "%s" WHERE vendor LIKE \?`, tableName)).
		WithArgs("%testvendor%").
		WillReturnRows(
			sqlmock.NewRows([]string{
				"vendor", "product", "version", "update", "edition", "language",
				"sw_edition", "target_sw", "target_hw", "other", "cpe23Uri", "cves",
			}).AddRow(
				"testvendor", "testproduct", "1.0", "", "", "",
				"", "", "", "", "cpe:2.3:a:testvendor:testproduct:1.0:*:*:*:*:*:*:*", `["CVE-1234","CVE-5678"]`,
			),
		)

	results, _, err := CPESearch(indexName, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Vendor != "testvendor" {
		t.Errorf("expected vendor testvendor, got %s", results[0].Vendor)
	}
}
