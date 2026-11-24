package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

// ExportStats summarizes an offline table dump.
type ExportStats struct {
	Index    string
	Rows     int
	Duration time.Duration
}

// KnownOfflineIndex reports whether indexName has a dedicated import schema
// (not the catch-all fallback).
func KnownOfflineIndex(indexName string) bool {
	for _, schema := range Schemas {
		if schema.Fallback {
			continue
		}
		for _, name := range schema.Indices {
			if name == indexName {
				return true
			}
		}
	}
	return false
}

// ExportIndexJSONL writes every row of a synced offline index as one JSON
// object per line. Column names match the schema (and therefore the SQLite
// table). Missing indexes return an error rather than an empty stream so
// callers can tell "not synced" from "synced but empty".
func ExportIndexJSONL(indexName string, out io.Writer) (*ExportStats, error) {
	start := time.Now()

	if !KnownOfflineIndex(indexName) {
		return nil, fmt.Errorf("index %s is not a known offline schema", indexName)
	}

	schema := GetSchema(indexName)
	if schema == nil {
		return nil, fmt.Errorf("no schema found for index %s", indexName)
	}

	db, err := DB()
	if err != nil {
		return nil, err
	}

	tableName := strings.ReplaceAll(indexName, "-", "_")
	colNames := make([]string, len(schema.Columns))
	for i, col := range schema.Columns {
		colNames[i] = `"` + col.Name + `"`
	}
	query := fmt.Sprintf(`SELECT %s FROM "%s"`, strings.Join(colNames, ", "), tableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query %s (is the index synced?): %w", indexName, err)
	}
	defer func() { _ = rows.Close() }()

	stats := &ExportStats{Index: indexName}
	encoder := json.NewEncoder(out)
	encoder.SetEscapeHTML(false)

	for rows.Next() {
		record, err := scanExportRow(rows, schema)
		if err != nil {
			return nil, err
		}
		if err := encoder.Encode(record); err != nil {
			return nil, fmt.Errorf("failed to encode row: %w", err)
		}
		stats.Rows++
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	stats.Duration = time.Since(start)
	return stats, nil
}

func scanExportRow(rows *sql.Rows, schema *Schema) (map[string]interface{}, error) {
	raw := make([]sql.NullString, len(schema.Columns))
	ptrs := make([]interface{}, len(schema.Columns))
	for i := range raw {
		ptrs[i] = &raw[i]
	}
	if err := rows.Scan(ptrs...); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	record := make(map[string]interface{}, len(schema.Columns))
	for i, col := range schema.Columns {
		if !raw[i].Valid {
			record[col.Name] = nil
			continue
		}
		if col.IsJSON {
			var decoded interface{}
			if err := json.Unmarshal([]byte(raw[i].String), &decoded); err == nil {
				record[col.Name] = decoded
				continue
			}
		}
		record[col.Name] = raw[i].String
	}
	return record, nil
}
