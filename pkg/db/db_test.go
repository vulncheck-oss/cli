package db

import (
	"database/sql"
	"os"
	"testing"
)

var dbPath string

func CreateTestDB() (*sql.DB, error) {
	// Create a temporary database file
	f, err := os.CreateTemp("", "testdb")
	if err != nil {
		return nil, err
	}
	dbPath = f.Name()
	f.Close()

	// Open the database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		os.Remove(dbPath) // Clean up the temporary file
		return nil, err
	}

	return db, nil
}

func TestDB(t *testing.T) {
	// Call CreateTestDB to set up the test database
	testDB, err := CreateTestDB()
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}
	defer func() {
		testDB.Close()
		os.Remove(dbPath) // Clean up the temporary database file
	}()

	// Call the DB function to get the database connection
	db, err := DB()
	if err != nil {
		t.Fatalf("DB() failed: %v", err)
	}
	defer db.Close()

	// Check if the returned database connection is valid
	err = db.Ping()
	if err != nil {
		t.Errorf("Ping failed: %v", err)
	}
}
