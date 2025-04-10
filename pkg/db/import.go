package db

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"os"
	"path/filepath"
	"strings"
)

// Maximum batch size in bytes for database inserts
const maxInsertSize int64 = 50 * 1024 * 1024 // 50 MB

func ImportIndex(filePath string, indexDir string, progressCallback func(int)) error {
	db, err := DB()
	if err != nil {
		return err
	}

	// Get schema for this index type
	indexName := packages.IndexFromName(filepath.Base(indexDir))
	schema := GetSchema(indexName)
	if schema == nil {
		return fmt.Errorf("no schema found for index %s", indexName)
	}

	// Convert table name to use underscores instead of hyphens
	tableName := strings.Replace(indexName, "-", "_", -1)

	// Drop existing table if it exists
	dropTableSQL := fmt.Sprintf(`DROP TABLE IF EXISTS "%s"`, tableName)
	if _, err := db.Exec(dropTableSQL); err != nil {
		return fmt.Errorf("failed to drop existing table: %w", err)
	}

	// Create table with schema
	cols := make([]string, len(schema.Columns))
	for i, col := range schema.Columns {
		def := fmt.Sprintf(`"%s" %s`, col.Name, col.Type)
		if col.NotNull {
			def += " NOT NULL"
		}
		cols[i] = def
	}
	createTableSQL := fmt.Sprintf(`CREATE TABLE "%s" (%s`, tableName, strings.Join(cols, ", "))

	if indexName == "golang" {
		createTableSQL += `, UNIQUE("name", "version")`
	}

	createTableSQL += ")"

	if _, err := db.Exec(createTableSQL); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Drop indexes before import
	for _, col := range schema.Columns {
		if col.Index {
			indexName := fmt.Sprintf("idx_%s_%s", tableName, col.Name)
			dropIndexSQL := fmt.Sprintf(`DROP INDEX IF EXISTS "idx_%s_%s"`, tableName, col.Name)
			if _, err := db.Exec(dropIndexSQL); err != nil {
				return fmt.Errorf("failed to drop index %s: %w", indexName, err)
			}
		}
	}

	// Find all JSON files
	files, err := filepath.Glob(filepath.Join(indexDir, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to list JSON files: %w", err)
	}

	// Prepare base insert statement
	colNames := make([]string, len(schema.Columns))
	placeholders := make([]string, len(schema.Columns))
	for i, col := range schema.Columns {
		colNames[i] = fmt.Sprintf(`"%s"`, col.Name)
		placeholders[i] = "?"
	}

	var baseInsertSQL string

	if indexName == "golang" {
		baseInsertSQL = fmt.Sprintf(`INSERT OR REPLACE INTO "%s" (%s) VALUES`,
			tableName, strings.Join(colNames, ","))
	} else {
		baseInsertSQL = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES`,
			tableName, strings.Join(colNames, ","))
	}

	totalSize := int64(0)
	for _, f := range files {
		if info, err := os.Stat(f); err == nil {
			totalSize += info.Size()
		}
	}

	// Define a consistent progress tracking function
	processedSize := int64(0)
	updateProgress := func(size int64) {
		processedSize += size
		if totalSize > 0 {
			progress := int(float64(processedSize) / float64(totalSize) * 100)
			// Ensure progress doesn't exceed 100%
			if progress > 100 {
				progress = 100
			}
			progressCallback(progress)
		}
	}

	// Process each file
	for _, file := range files {
		if err := importFile(db, file, schema, baseInsertSQL, maxInsertSize, updateProgress); err != nil {
			return fmt.Errorf("failed to import file %s: %w", file, err)
		}
	}

	// Recreate indexes after import
	for _, col := range schema.Columns {
		if col.Index {
			indexSQL := fmt.Sprintf(`CREATE INDEX "idx_%s_%s" ON "%s"("%s")`, tableName, col.Name, tableName, col.Name)
			if _, err := db.Exec(indexSQL); err != nil {
				return fmt.Errorf("failed to create index: %w", err)
			}
		}
	}

	return nil
}

func importFile(db *sql.DB, filePath string, schema *Schema, baseInsertSQL string, maxSize int64, progressFn func(int64)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read first character to check if it's an array
	firstByte := make([]byte, 1)
	if _, err := file.Read(firstByte); err != nil {
		return fmt.Errorf("failed to read first byte: %w", err)
	}
	isArray := firstByte[0] == '['

	// Reset file pointer
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %w", err)
	}

	var batch [][]interface{}
	var batchSize int64

	// Cache JSON array column indices
	jsonColumns := make(map[int]bool)
	for i, col := range schema.Columns {
		if col.IsJSON {
			jsonColumns[i] = true
		}
	}

	if isArray {
		// Handle array of objects
		decoder := json.NewDecoder(file)
		// Read opening bracket
		_, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read array start: %w", err)
		}

		// Read array elements
		for decoder.More() {
			var entry map[string]interface{}
			if err := decoder.Decode(&entry); err != nil {
				return fmt.Errorf("failed to decode JSON object: %w", err)
			}

			if values, size, err := processEntry(entry, schema, jsonColumns); err == nil {
				batch = append(batch, values)
				batchSize += size

				if batchSize >= maxSize { // Use maxSize parameter instead of hardcoded value
					if err := executeBatch(db, baseInsertSQL, batch); err != nil {
						return err
					}
					progressFn(batchSize)
					batch = batch[:0]
					batchSize = 0
				}
			}
		}
	} else {
		// Original line-by-line processing
		scanner := bufio.NewScanner(file)
		scanner.Buffer(make([]byte, 4*1024*1024), 4*1024*1024)

		for scanner.Scan() {
			line := scanner.Bytes()

			// Fast path for fallback schema
			if len(schema.Columns) == 1 && schema.Columns[0].Name == "data" {
				if !json.Valid(line) {
					return fmt.Errorf("invalid JSON")
				}
				batch = append(batch, []interface{}{string(line)})
				batchSize += int64(len(line))
				continue
			}

			var entry map[string]interface{}
			if err := json.Unmarshal(line, &entry); err != nil {
				return fmt.Errorf("failed to unmarshal JSON: %w", err)
			}

			if values, size, err := processEntry(entry, schema, jsonColumns); err == nil {
				batch = append(batch, values)
				batchSize += size

				if batchSize >= maxSize { // Use maxSize parameter instead of hardcoded value
					if err := executeBatch(db, baseInsertSQL, batch); err != nil {
						return err
					}
					progressFn(batchSize)
					batch = batch[:0]
					batchSize = 0
				}
			}
		}
	}

	// Process remaining batch
	if len(batch) > 0 {
		if err := executeBatch(db, baseInsertSQL, batch); err != nil {
			return err
		}
		progressFn(batchSize)
	}

	return nil
}

func processEntry(entry map[string]interface{}, schema *Schema, jsonColumns map[int]bool) ([]interface{}, int64, error) {
	// Only check for CVEs if it's a PM schema
	if schema.Name == "purl PM" {
		cves, hasCVEs := entry["cves"].([]interface{})
		if !hasCVEs || len(cves) == 0 {
			return nil, 0, fmt.Errorf("skip entry without CVEs")
		}
	}

	// Rest of the function remains the same...
	values := make([]interface{}, len(schema.Columns))
	for i, col := range schema.Columns {
		val, exists := entry[col.Name]
		if !exists {
			if col.NotNull {
				return nil, 0, fmt.Errorf("missing required field %s", col.Name)
			}
			values[i] = nil
			continue
		}

		// Process the version field to extract base version
		if col.Name == "version" {
			if versionStr, ok := val.(string); ok {
				// Extract base version (part before first hyphen)
				parts := strings.Split(versionStr, "-")
				val = parts[0]
			}
		}
		if jsonColumns[i] {
			if arr, ok := val.([]interface{}); ok {
				jsonStr, _ := json.Marshal(arr)
				values[i] = string(jsonStr)
			} else {
				jsonStr, _ := json.Marshal(val)
				values[i] = string(jsonStr)
			}
		} else {
			values[i] = val
		}
	}

	size := int64(len(fmt.Sprintf("%v", entry)))
	return values, size, nil
}

func executeBatch(db *sql.DB, baseSQL string, batch [][]interface{}) error {
	if len(batch) == 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	varsPerRow := len(batch[0])
	maxRowsPerBatch := maxSQLiteVariables / varsPerRow

	for i := 0; i < len(batch); i += maxRowsPerBatch {
		end := i + maxRowsPerBatch
		if end > len(batch) {
			end = len(batch)
		}
		subBatch := batch[i:end]

		// Build VALUES clause and args
		values := make([]string, len(subBatch))
		args := make([]interface{}, 0, len(subBatch)*len(batch[0]))
		placeholder := "(" + strings.Repeat("?,", len(batch[0])-1) + "?)"

		for j, row := range subBatch {
			values[j] = placeholder
			args = append(args, row...)
		}

		// Execute batch insert
		query := baseSQL + " " + strings.Join(values, ",")
		_, err = tx.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}
	}

	return tx.Commit()
}
