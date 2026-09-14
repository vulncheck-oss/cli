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

// writeIndex writes NDJSON records to a temp index dir and returns the dir.
// The basename decides the schema: "osv" has none, so it gets the fallback
// single-"data"-column table.
func writeIndex(t *testing.T, index, contents string) string {
	t.Helper()
	indexDir := filepath.Join(t.TempDir(), index)
	if err := os.MkdirAll(indexDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(indexDir, "data.json"), []byte(contents), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	t.Cleanup(func() {
		_, _ = testDB.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %q`, index))
	})
	return indexDir
}

func fallbackRecord(t *testing.T, id string, padBytes int) string {
	t.Helper()
	b, err := json.Marshal(map[string]string{"id": id, "pad": strings.Repeat("a", padBytes)})
	if err != nil {
		t.Fatalf("marshal %s: %v", id, err)
	}
	return string(b)
}

// TestImportIndexOversizedRecord covers records past the old 4MB scanner
// ceiling, which is what broke osv for #3628 (its largest record is ~19MB).
// ReadBytes has no ceiling, so the only limit left is SQLITE_MAX_LENGTH.
func TestImportIndexOversizedRecord(t *testing.T) {
	big := fallbackRecord(t, "big", 5_000_000) // comfortably over the old 4MB
	records := []string{
		fallbackRecord(t, "before", 10),
		big,
		// A record after the oversized one: pre-fix the scanner stopped on the
		// previous line and this was silently dropped.
		fallbackRecord(t, "after", 10),
	}
	indexDir := writeIndex(t, "osv", strings.Join(records, "\n")+"\n")

	if err := ImportIndex("", indexDir, func(int) {}); err != nil {
		t.Fatalf("ImportIndex: %v", err)
	}

	var count int
	if err := testDB.QueryRow(`SELECT COUNT(*) FROM osv`).Scan(&count); err != nil {
		t.Fatalf("count: %v", err)
	}
	if count != len(records) {
		t.Errorf("expected %d rows, got %d", len(records), count)
	}

	// The oversized record must be stored whole, not truncated.
	var stored int
	if err := testDB.QueryRow(
		`SELECT LENGTH(data) FROM osv WHERE json_extract(data,'$.id') = 'big'`).Scan(&stored); err != nil {
		t.Fatalf("read big record: %v", err)
	}
	if stored != len(big) {
		t.Errorf("oversized record truncated: stored %d bytes, wrote %d", stored, len(big))
	}
}

// TestImportIndexFallbackFlushesBatches guards the flush in the fallback fast
// path. It used to skip the batch-size check and accumulate the whole file,
// which cost 12GB of heap on osv.
func TestImportIndexFallbackFlushesBatches(t *testing.T) {
	const want = 1200 // crosses the 500-record flush threshold twice
	records := make([]string, want)
	for i := range records {
		records[i] = fallbackRecord(t, fmt.Sprintf("rec-%04d", i), 10)
	}
	indexDir := writeIndex(t, "osv", strings.Join(records, "\n")+"\n")

	// Each flush reports progress exactly once, so the callback count tells us
	// whether the import streamed or buffered the whole file. Row counts alone
	// cannot: the old code accumulated everything and flushed once at the end,
	// which is correct but is the 12GB-of-heap behaviour we are fixing.
	flushes := 0
	if err := ImportIndex("", indexDir, func(int) { flushes++ }); err != nil {
		t.Fatalf("ImportIndex: %v", err)
	}

	var count int
	if err := testDB.QueryRow(`SELECT COUNT(*) FROM osv`).Scan(&count); err != nil {
		t.Fatalf("count: %v", err)
	}
	if count != want {
		t.Errorf("expected %d rows, got %d - records lost across batch flushes", want, count)
	}
	if flushes < 2 {
		t.Errorf("got %d flush(es) for %d records: the fallback path is buffering the whole file instead of flushing every 500", flushes, want)
	}
}

// TestImportIndexLineEndings covers the framing cases ReadBytes has to handle:
// CRLF endings, blank lines, and a final record with no trailing newline.
// Blank lines are skipped rather than fatal, which is a deliberate relaxation:
// bufio.Scanner emitted an empty token and the import died with "invalid JSON".
func TestImportIndexLineEndings(t *testing.T) {
	contents := fallbackRecord(t, "crlf", 10) + "\r\n" +
		"\n" +
		fallbackRecord(t, "after-blank", 10) + "\n" +
		"\r\n" +
		fallbackRecord(t, "no-trailing-newline", 10) // deliberately unterminated
	indexDir := writeIndex(t, "osv", contents)

	if err := ImportIndex("", indexDir, func(int) {}); err != nil {
		t.Fatalf("ImportIndex: %v", err)
	}

	for _, id := range []string{"crlf", "after-blank", "no-trailing-newline"} {
		var count int
		if err := testDB.QueryRow(
			`SELECT COUNT(*) FROM osv WHERE json_extract(data,'$.id') = ?`, id).Scan(&count); err != nil {
			t.Fatalf("count %q: %v", id, err)
		}
		if count != 1 {
			t.Errorf("record %q: expected 1 row, got %d", id, count)
		}
	}

	// Blank lines must not have produced rows of their own.
	var total int
	if err := testDB.QueryRow(`SELECT COUNT(*) FROM osv`).Scan(&total); err != nil {
		t.Fatalf("total: %v", err)
	}
	if total != 3 {
		t.Errorf("expected 3 rows, got %d - blank lines should be skipped", total)
	}
}
