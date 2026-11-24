package db

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestExportIndexJSONL(t *testing.T) {
	schema := GetSchema("vulncheck-nvd2")
	if schema == nil {
		t.Fatal("missing nvd schema")
	}
	if _, err := testDB.Exec(`DROP TABLE IF EXISTS vulncheck_nvd2`); err != nil {
		t.Fatal(err)
	}
	cols := make([]string, len(schema.Columns))
	for i, col := range schema.Columns {
		colDef := `"` + col.Name + `" ` + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		cols[i] = colDef
	}
	if _, err := testDB.Exec(`CREATE TABLE vulncheck_nvd2 (` + strings.Join(cols, ", ") + `)`); err != nil {
		t.Fatal(err)
	}
	defer func() { _, _ = testDB.Exec(`DROP TABLE IF EXISTS vulncheck_nvd2`) }()

	_, err := testDB.Exec(`INSERT INTO vulncheck_nvd2 (id, published, description)
		VALUES (?, ?, ?)`,
		"CVE-2024-0001", "2024-01-01T00:00:00Z", "Example")
	if err != nil {
		t.Fatalf("insert: %v", err)
	}

	var buf bytes.Buffer
	stats, err := ExportIndexJSONL("vulncheck-nvd2", &buf)
	if err != nil {
		t.Fatalf("export: %v", err)
	}
	if stats.Rows != 1 {
		t.Fatalf("expected 1 row, got %d", stats.Rows)
	}

	var row map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &row); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if row["id"] != "CVE-2024-0001" {
		t.Errorf("id=%v", row["id"])
	}
	if row["description"] != "Example" {
		t.Errorf("description=%v", row["description"])
	}
	if row["published"] != "2024-01-01T00:00:00Z" {
		t.Errorf("published=%v", row["published"])
	}
}

func TestExportUnknownIndex(t *testing.T) {
	_, err := ExportIndexJSONL("not-a-real-index", &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected error for unknown index")
	}
}

func TestKnownOfflineIndex(t *testing.T) {
	if !KnownOfflineIndex("npm") || !KnownOfflineIndex("cpecve") || !KnownOfflineIndex("vulncheck-nvd2") {
		t.Fatal("expected known indexes")
	}
	if KnownOfflineIndex("nope") {
		t.Fatal("expected unknown")
	}
}
