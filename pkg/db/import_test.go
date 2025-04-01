package db

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestImportIndex(t *testing.T) {
	originalDBFunc := DBFunc
	defer func() { DBFunc = originalDBFunc }()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	DBFunc = func() (*sql.DB, error) {
		return mockDB, nil
	}

	tableName := "test_index"
	mock.ExpectExec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s"`, tableName)).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(fmt.Sprintf(`CREATE TABLE "%s"`, tableName)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = ImportIndex("someFilePath", "test-index", func(progress int) {})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
