package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vulncheck-oss/cli/pkg/config"
	"os"
	"path/filepath"
)

var dbInstance *sql.DB
var maxInsertSize int64 = 1_000_000_000 // Default max length in bytes
const maxSQLiteVariables = 900          // Slightly below SQLite's limit of 999 to be safe
// DB provides a cached database connection.
func DB() (*sql.DB, error) {
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
        dbInstance, err = sql.Open("sqlite3", dbPath+"?_journal_mode=WAL")
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

        dbInstance, err = sql.Open("sqlite3", dbPath+"?_journal_mode=WAL")
        if err != nil {
            return nil, fmt.Errorf("failed to open database: %w", err)
        }
    }

    // Performance optimizations
    _, err = dbInstance.Exec(`
        PRAGMA journal_mode = WAL;
        PRAGMA synchronous = NORMAL;
        PRAGMA cache_size = -2000000; 
        PRAGMA temp_store = MEMORY;
    `)
    if err != nil {
        return nil, fmt.Errorf("failed to set PRAGMA statements: %w", err)
    }

    return dbInstance, nil
}