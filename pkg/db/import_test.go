package db

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestImportIndex_NDJSONLargeRecord(t *testing.T) {
	tempRoot := t.TempDir()
	indexDir := filepath.Join(tempRoot, "cpecve")
	if err := os.MkdirAll(indexDir, 0o755); err != nil {
		t.Fatalf("failed to create index dir: %v", err)
	}

	// Build a record larger than the historical scanner token limit (256KB).
	largeCVEs := make([]string, 0, 32000)
	for i := 0; i < 32000; i++ {
		largeCVEs = append(largeCVEs, fmt.Sprintf("\"CVE-2026-%05d\"", i))
	}

	largeLine := fmt.Sprintf("{\"part\":\"o\",\"vendor\":\"linux\",\"product\":\"linux_kernel\",\"version\":\"5\\\\.15\\\\.201\",\"update\":\"*\",\"edition\":\"*\",\"language\":\"*\",\"sw_edition\":\"*\",\"target_sw\":\"*\",\"target_hw\":\"*\",\"other\":\"*\",\"cpe23Uri\":\"cpe:/o:linux:linux_kernel:5.15.201\",\"cves\":[%s]}", strings.Join(largeCVEs, ","))
	smallLine := "{\"part\":\"o\",\"vendor\":\"linux\",\"product\":\"linux_kernel\",\"version\":\"5\\\\.15\\\\.202\",\"update\":\"*\",\"edition\":\"*\",\"language\":\"*\",\"sw_edition\":\"*\",\"target_sw\":\"*\",\"target_hw\":\"*\",\"other\":\"*\",\"cpe23Uri\":\"cpe:/o:linux:linux_kernel:5.15.202\",\"cves\":[\"CVE-2026-99999\"]}"

	filePath := filepath.Join(indexDir, "chunk-1.json")
	content := largeLine + "\n" + smallLine + "\n"
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write ndjson fixture: %v", err)
	}

	if err := ImportIndex(filePath, indexDir, func(int) {}); err != nil {
		t.Fatalf("import failed for large NDJSON record: %v", err)
	}

	var count int
	if err := testDB.QueryRow(`SELECT COUNT(*) FROM cpecve`).Scan(&count); err != nil {
		t.Fatalf("failed to count imported rows: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2 imported rows, got %d", count)
	}
}
