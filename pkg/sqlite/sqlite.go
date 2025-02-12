package sqlite

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vulncheck-oss/cli/pkg/config"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var dbInstance *sql.DB
var maxInsertSize int64 = 1_000_000_000 // Default max length in bytes

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
		dbInstance, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}
		return dbInstance, nil
	} else if !os.IsNotExist(err) {
		// An error other than "file does not exist" occurred
		return nil, fmt.Errorf("failed to check database file: %w", err)
	}

	// File does not exist, create a new database file
	file, err := os.Create(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create database file: %w", err)
	}
	file.Close()

	dbInstance, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return dbInstance, nil
}

func ParseAndInsertJSON(db *sql.DB, tableName string, filePath string, progressCallback func(int)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	totalLines := 0
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read JSON file: %w", err)
		}
		totalLines++
	}

	file.Seek(0, 0)
	reader = bufio.NewReader(file)
	currentLine := 0
	var batch []string
	var batchSize int64

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read JSON file: %w", err)
		}

		var entry map[string]interface{}
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			return fmt.Errorf("failed to unmarshal JSON data: %w", err)
		}

		entryData, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON entry: %w", err)
		}

		batch = append(batch, string(entryData))
		batchSize += int64(len(entryData))

		if batchSize >= maxInsertSize {
			if err := insertBatch(db, tableName, batch); err != nil {
				return err
			}
			batch = nil
			batchSize = 0
		}

		currentLine++
		if currentLine%100 == 0 {
			progressCallback(currentLine * 100 / totalLines)
		}
	}

	if len(batch) > 0 {
		if err := insertBatch(db, tableName, batch); err != nil {
			return err
		}
	}

	progressCallback(100)
	return nil
}

func insertBatch(db *sql.DB, tableName string, batch []string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	insertSQL := fmt.Sprintf(`INSERT INTO "%s" (data) VALUES (?);`, tableName)
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, entry := range batch {
		if _, err := stmt.Exec(entry); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute batch insert: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func JSONTable(filePath string, indexDir string, progressCallback func(int)) error {
	db, err := DB()
	if err != nil {
		return err
	}

	parts := strings.Split(filepath.Base(filePath), "-")
	if len(parts) < 2 {
		return fmt.Errorf("invalid tableName format, expected {index}-{timestamp}.zip")
	}

	tableName := parts[0]

	dropTableSQL := fmt.Sprintf(`DROP TABLE IF EXISTS "%s";`, tableName)
	_, err = db.Exec(dropTableSQL)
	if err != nil {
		return fmt.Errorf("failed to drop existing table: %w", err)
	}

	createTableSQL := fmt.Sprintf(`CREATE TABLE "%s" (data JSON);`, tableName)
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	files, err := os.ReadDir(indexDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	totalFiles := len(files)
	processedFiles := 0

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(indexDir, file.Name())
			if err := ParseAndInsertJSON(db, tableName, filePath, func(progress int) {
				progressCallback((processedFiles*100 + progress) / totalFiles)
			}); err != nil {
				return err
			}
			processedFiles++
		}
	}

	progressCallback(100)
	return nil
}
