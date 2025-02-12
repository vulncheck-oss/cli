package db

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ImportIndex(filePath string, indexDir string, progressCallback func(int)) error {
	db, err := DB()
	if err != nil {
		return err
	}

	// Get schema for this index type
	indexName := filepath.Base(indexDir)
	schema := GetSchema(indexName)
	if schema == nil {
		return fmt.Errorf("no schema found for index %s", indexName)
	}

	// Drop existing table if it exists
	tableName := strings.Replace(indexName, "-", "_", -1)
	if _, err := db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s"`, tableName)); err != nil {
		return fmt.Errorf("failed to drop existing table: %w", err)
	}

	// Create table with schema
	cols := make([]string, len(schema.Columns))
	for i, col := range schema.Columns {
		def := fmt.Sprintf("%s %s", col.Name, col.Type)
		if col.NotNull {
			def += " NOT NULL"
		}
		cols[i] = def
	}
	createTableSQL := fmt.Sprintf(`CREATE TABLE "%s" (%s)`,
		tableName, strings.Join(cols, ", "))

	if _, err := db.Exec(createTableSQL); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Create indexes
	for _, col := range schema.Columns {
		if col.Index {
			indexSQL := fmt.Sprintf("CREATE INDEX idx_%s_%s ON %s(%s)",
				tableName, col.Name, tableName, col.Name)
			if _, err := db.Exec(indexSQL); err != nil {
				return fmt.Errorf("failed to create index: %w", err)
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
		colNames[i] = col.Name
		placeholders[i] = "?"
	}
	baseInsertSQL := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES`,
		tableName, strings.Join(colNames, ","))

	totalSize := int64(0)
	for _, f := range files {
		if info, err := os.Stat(f); err == nil {
			totalSize += info.Size()
		}
	}

	processedSize := int64(0)
	for fileNum, file := range files {
		if err := importFile(db, file, schema, baseInsertSQL, maxInsertSize, func(size int64) {
			processedSize += size
			progress := int(float64(processedSize) / float64(totalSize) * 100)
			progressCallback(progress)
		}); err != nil {
			return fmt.Errorf("failed to import file %s: %w", file, err)
		}
		progressCallback(int(float64(fileNum+1) / float64(len(files)) * 100))
	}

	return nil
}

func importFile(db *sql.DB, filePath string, schema *Schema, baseInsertSQL string, maxSize int64, progressFn func(int64)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Use larger scanner buffer for better performance
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 4*1024*1024), 4*1024*1024) // 4MB buffer

	// Pre-allocate batch slice to reduce allocations
	batch := make([][]interface{}, 0, 10000)
	var batchSize int64

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Cache JSON array column indices
	jsonColumns := make(map[int]bool)
	for i, col := range schema.Columns {
		if col.IsJSON {
			jsonColumns[i] = true
		}
	}

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

		// Handle structured schema
		var entry map[string]interface{}
		if err := json.Unmarshal(line, &entry); err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}

		values := make([]interface{}, len(schema.Columns))
		for i, col := range schema.Columns {
			val, exists := entry[col.Name]
			if !exists {
				if col.NotNull {
					return fmt.Errorf("missing required field %s", col.Name)
				}
				values[i] = nil
				continue
			}

			if jsonColumns[i] {
				// Only marshal JSON for array fields
				if arr, ok := val.([]interface{}); ok {
					jsonStr, _ := json.Marshal(arr)
					values[i] = string(jsonStr)
				} else {
					values[i] = nil
				}
			} else {
				values[i] = val
			}
		}

		batch = append(batch, values)
		batchSize += int64(len(line))

		// Use larger batch size for better performance
		if batchSize >= 50*1024*1024 { // 50MB batches
			if err := executeBatch(tx, baseInsertSQL, batch); err != nil {
				return err
			}
			progressFn(batchSize)

			batch = batch[:0] // Reuse slice
			batchSize = 0

			if err := tx.Commit(); err != nil {
				return err
			}
			tx, err = db.Begin()
			if err != nil {
				return err
			}
		}
	}

	if len(batch) > 0 {
		if err := executeBatch(tx, baseInsertSQL, batch); err != nil {
			return err
		}
		progressFn(batchSize)
	}

	return tx.Commit()
}

func executeBatch(tx *sql.Tx, baseSQL string, batch [][]interface{}) error {
	if len(batch) == 0 {
		return nil
	}

	varsPerRow := len(batch[0])
	maxRowsPerBatch := maxSQLiteVariables / varsPerRow

	// Pre-allocate slices
	values := make([]string, 0, maxRowsPerBatch)
	args := make([]interface{}, 0, maxRowsPerBatch*varsPerRow)
	placeholder := "(" + strings.Repeat("?,", varsPerRow-1) + "?)"

	for i := 0; i < len(batch); i += maxRowsPerBatch {
		end := i + maxRowsPerBatch
		if end > len(batch) {
			end = len(batch)
		}

		values = values[:0]
		args = args[:0]

		// Build batch
		for j := i; j < end; j++ {
			values = append(values, placeholder)
			args = append(args, batch[j]...)
		}

		query := baseSQL + " " + strings.Join(values, ",")
		if _, err := tx.Exec(query, args...); err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}
	}

	return nil
}
