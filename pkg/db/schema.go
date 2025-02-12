package db

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/octoper/go-ray"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Column struct {
	Name    string // Column name in database
	Type    string // SQL type (TEXT, INTEGER, etc)
	Index   bool   // Whether to create an index
	NotNull bool   // Whether column can be null
	IsJSON  bool   // Whether value is JSON array
}

type Schema struct {
	// IndexMatch will be a regex that matches index names, like ipintel-(3|7|10)d / etc
	IndexMatch string
	Fallback   bool
	Columns    []Column
}

var Schemas = []Schema{
	{
		Fallback:   true,
		IndexMatch: "",
		Columns: []Column{
			{Name: "data", Type: "TEXT", Index: true, NotNull: true, IsJSON: true},
		},
	},
	{
		IndexMatch: `^ipintel-\d+d$`,
		Columns: []Column{
			// Primary search fields - all indexed
			{Name: "ip", Type: "TEXT", Index: true, NotNull: true},
			{Name: "country", Type: "TEXT", Index: true, NotNull: true},
			{Name: "asn", Type: "TEXT", Index: true, NotNull: true},
			{Name: "country_code", Type: "TEXT", Index: true, NotNull: true},
			{Name: "hostnames", Type: "TEXT", Index: true, NotNull: false, IsJSON: true},
			{Name: "type_id", Type: "TEXT", Index: true, NotNull: false},

			// Non-searched fields - no indexes needed
			{Name: "port", Type: "INTEGER", Index: false, NotNull: true},
			{Name: "ssl", Type: "BOOLEAN", Index: false, NotNull: true},
			{Name: "lastSeen", Type: "TEXT", Index: false, NotNull: true},
			{Name: "city", Type: "TEXT", Index: false, NotNull: false},
			{Name: "cve", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "matches", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "type_kind", Type: "TEXT", Index: false, NotNull: false},
			{Name: "type_finding", Type: "TEXT", Index: false, NotNull: false},
			{Name: "feed_ids", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
		},
	},
}

func GetSchema(indexName string) *Schema {
	// First try to find a matching schema
	for _, schema := range Schemas {
		if schema.Fallback {
			continue // Skip fallback during regex matching
		}
		matched, _ := regexp.MatchString(schema.IndexMatch, indexName)
		if matched {
			return &schema
		}
	}

	// If no match found, return the fallback schema
	for _, schema := range Schemas {
		if schema.Fallback {
			return &schema
		}
	}

	return nil
}

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

    scanner := bufio.NewScanner(file)
    scanner.Buffer(make([]byte, 1024*1024), 1024*1024) // 1MB buffer

    var batch [][]interface{}
    var batchSize int64

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for scanner.Scan() {
        line := scanner.Bytes()

        // Handle fallback schema (single JSON column) differently
        if len(schema.Columns) == 1 && schema.Columns[0].Name == "data" {
            // Verify it's valid JSON before inserting
            var decoded interface{}
            if err := json.Unmarshal(line, &decoded); err != nil {
                return fmt.Errorf("invalid JSON: %w", err)
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

            switch {
            case col.IsJSON:
                // Always marshal JSON fields
                jsonStr, err := json.Marshal(val)
                if err != nil {
                    return fmt.Errorf("failed to marshal JSON field %s: %w", col.Name, err)
                }
                values[i] = string(jsonStr)
            default:
                values[i] = val
            }
        }

        batch = append(batch, values)
        batchSize += int64(len(line))

        if batchSize >= maxSize {
            if err := executeBatch(tx, baseInsertSQL, batch); err != nil {
                return err
            }
            progressFn(batchSize)
            
            // Reset batch
            batch = nil
            batchSize = 0

            // Commit and start new transaction
            if err := tx.Commit(); err != nil {
                return err
            }
            tx, err = db.Begin()
            if err != nil {
                return err
            }
        }
    }

    // Handle remaining batch
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

	// Calculate how many rows we can include per statement
	// based on SQLite's variable limit
	varsPerRow := len(batch[0])
	maxRowsPerBatch := maxSQLiteVariables / varsPerRow

	// Process in sub-batches if needed
	for i := 0; i < len(batch); i += maxRowsPerBatch {
		end := i + maxRowsPerBatch
		if end > len(batch) {
			end = len(batch)
		}
		subBatch := batch[i:end]

		// Build VALUES clause and args for this sub-batch
		values := make([]string, len(subBatch))
		args := make([]interface{}, 0, len(subBatch)*varsPerRow)
		placeholder := "(" + strings.Repeat("?,", varsPerRow-1) + "?)"

		for j := range subBatch {
			values[j] = placeholder
			args = append(args, subBatch[j]...)
		}

		// Execute this sub-batch
		query := baseSQL + " " + strings.Join(values, ",")
		ray.Ray(query, args)
		if _, err := tx.Exec(query, args...); err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}
	}

	return nil
}
