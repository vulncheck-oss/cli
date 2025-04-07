package db

import (
	"database/sql"
	"os"
	"strings"
	"testing"
)

var testDB *sql.DB

// TestMain is the main entry point for all tests in the db package
func TestMain(m *testing.M) {
	// Set TEST_ENV for in-memory database
	os.Setenv("TEST_ENV", "true")

	// Initialize the database once for all tests
	var err error
	testDB, err = DB()
	if err != nil {
		panic("failed to initialize test database: " + err.Error())
	}

	// Run all tests
	code := m.Run()

	// Close the database connection
	testDB.Close()

	os.Exit(code)
}

// createTableFromSchema is a helper function to create a test table from a schema
func createTableFromSchema(t *testing.T, tableName string) error {
	schema := GetSchema(tableName)
	if schema == nil {
		t.Fatalf("failed to get schema for %s", tableName)
		return nil
	}

	var columns []string
	for _, col := range schema.Columns {
		colDef := "\"" + col.Name + "\" " + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		columns = append(columns, colDef)
	}

	tableDef := "CREATE TABLE IF NOT EXISTS " + schema.Name + " (" + strings.Join(columns, ", ") + ")"
	_, err := testDB.Exec(tableDef)
	return err
}