package db

import (
	"database/sql"
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/config"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

var dbInstance *sql.DB

const maxSQLiteVariables = 900 // Slightly below SQLite's limit of 999 to be safe
// DB provides a cached database connection.
func DB() (*sql.DB, error) {
	if os.Getenv("TEST_ENV") == "true" {
		if dbInstance == nil {
			var err error
			dbInstance, err = sql.Open("sqlite", ":memory:")
			if err != nil {
				return nil, fmt.Errorf("failed to open in-memory database: %w", err)
			}
		}
		return dbInstance, nil
	}

	if dbInstance != nil {
		return dbInstance, nil
	}

	configDir, err := config.IndicesDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get indices directory: %w", err)
	}

	dbPath := filepath.Join(configDir, "data.db")
	if _, err := os.Stat(dbPath); err == nil {
		// File exists, open the existing database
		dbInstance, err = sql.Open("sqlite", dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check database file: %w", err)
	} else {
		// File does not exist, create a new database file
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %w", err)
		}
		file.Close()

		dbInstance, err = sql.Open("sqlite", dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}
	}

	return dbInstance, nil
}
