package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestImportIndexBufferOverflow guards against silent truncation when a record
// in a line-by-line index exceeds the scanner buffer. Pre-fix, bufio.Scanner's
// 256KB ceiling stopped at the first oversized line and the rest of the file
// was lost without an error - the symptom that caused #2580 (cpecve missing
// ~50k of its ~1.6M records, including every linux_kernel 5.15.* row).
func TestImportIndexBufferOverflow(t *testing.T) {
	dir := t.TempDir()
	indexDir := filepath.Join(dir, "cpecve")
	if err := os.MkdirAll(indexDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	// Build a CVE list big enough that the encoded line clears 256KB.
	// At ~14 bytes per CVE this gives ~400KB, well over the old ceiling.
	largeCves := make([]string, 30_000)
	for i := range largeCves {
		largeCves[i] = fmt.Sprintf("CVE-2099-%05d", i)
	}
	cvesJSON, err := json.Marshal(largeCves)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	records := []string{
		// Small record - imports trivially.
		`{"part":"a","vendor":"first","product":"p","version":"1.0","cves":["CVE-A"]}`,
		// Oversized record - mimics the cpecve rows that broke the old import.
		fmt.Sprintf(`{"part":"o","vendor":"jumbo","product":"big","version":"5.15.201","cves":%s}`, string(cvesJSON)),
		// Small record AFTER the big one - this is the proof. Pre-fix the
		// scanner stopped on the previous line and never reached this one.
		`{"part":"a","vendor":"after","product":"tail","version":"3.0","cves":["CVE-Z"]}`,
	}

	fixture := filepath.Join(indexDir, "data.json")
	if err := os.WriteFile(fixture, []byte(strings.Join(records, "\n")+"\n"), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	// Clean the table when we're done so unrelated cpecve tests aren't
	// affected by whatever order they run in.
	t.Cleanup(func() {
		_, _ = testDB.Exec(`DROP TABLE IF EXISTS cpecve`)
	})

	if err := ImportIndex(fixture, indexDir, func(int) {}); err != nil {
		t.Fatalf("ImportIndex: %v", err)
	}

	// All three vendors must be present. Pre-fix only "first" would survive.
	for _, vendor := range []string{"first", "jumbo", "after"} {
		var c int
		if err := testDB.QueryRow(`SELECT COUNT(*) FROM cpecve WHERE vendor = ?`, vendor).Scan(&c); err != nil {
			t.Fatalf("count %q: %v", vendor, err)
		}
		if c != 1 {
			t.Errorf("vendor %q: expected 1 row, got %d (pre-fix the tail-of-file records were silently dropped)", vendor, c)
		}
	}

	// And the oversized row's CVE array must be intact end-to-end - not
	// truncated, not encoded as the wrong shape.
	var storedCves string
	if err := testDB.QueryRow(`SELECT cves FROM cpecve WHERE vendor = 'jumbo'`).Scan(&storedCves); err != nil {
		t.Fatalf("read jumbo cves: %v", err)
	}
	if !strings.Contains(storedCves, "CVE-2099-29999") {
		t.Errorf("oversized row's last CVE missing - was the value truncated? len=%d", len(storedCves))
	}
}
