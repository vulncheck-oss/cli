package search

import (
	"fmt"
	"github.com/itchyny/gojq"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFilter(t *testing.T) {
	testCases := []struct {
		name     string
		json     string
		query    string
		expected bool
	}{
		{"Match country", `{"country": "Narnia"}`, ".country == \"Narnia\"", true},
		{"Match port", `{"port": 443}`, ".port == 443", true},
		{"Match CVE", `{"cve": ["CVE-2021-36260"]}`, "any(.cve[] | . == \"CVE-2021-36260\")", true},
		{"No match", `{"country": "Neverland"}`, ".country == \"Narnia\"", false},
		{"Match nested field", `{"type": {"id": "initial-access"}}`, ".type.id == \"initial-access\"", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := quickFilter([]byte(tc.json), tc.query)
			assert.Equal(t, tc.expected, result, "Unexpected result for query: %s", tc.query)
		})
	}
}

func TestParseQuery(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected map[string]string
	}{
		{
			name:     "Single condition",
			query:    ".country == \"Narnia\"",
			expected: map[string]string{"country": "Narnia"},
		},
		{
			name:     "Multiple conditions",
			query:    ".country == \"Narnia\" and .port == 443",
			expected: map[string]string{"country": "Narnia", "port": "443"},
		},
		{
			name:     "Nested field",
			query:    ".type.id == \"initial-access\"",
			expected: map[string]string{"type.id": "initial-access"},
		},
		{
			name:     "Array contains",
			query:    "any(.cve[] | . == \"CVE-2021-36260\")",
			expected: map[string]string{"cve": "CVE-2021-36260"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := parseQuery(tc.query)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestListIndexFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test_index_")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a nested directory structure with some files
	files := []string{
		"file1.json",
		"file2.json",
		"file3.txt",
		"subdir/file4.json",
		"subdir/file5.txt",
		"subdir/nested/file6.json",
	}

	for _, file := range files {
		path := filepath.Join(tempDir, file)
		err := os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
		err = os.WriteFile(path, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
	}

	// Call the function we're testing
	result, err := listIndexFiles(tempDir)
	if err != nil {
		t.Fatalf("listIndexFiles returned an error: %v", err)
	}

	// Check the results
	expected := []string{
		filepath.Join(tempDir, "file1.json"),
		filepath.Join(tempDir, "file2.json"),
		filepath.Join(tempDir, "subdir", "file4.json"),
		filepath.Join(tempDir, "subdir", "nested", "file6.json"),
	}

	// Sort both slices to ensure consistent ordering
	sort.Strings(result)
	sort.Strings(expected)

	assert.Equal(t, expected, result, "Unexpected list of JSON files")
}

func TestProcessFile(t *testing.T) {
	// Create a temporary JSON file for testing
	tempFile, err := os.CreateTemp("", "test_process_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the file
	testData := `{"ip":"192.168.1.1","port":80,"country":"TestCountry","cve":["CVE-2021-1234"]}
{"ip":"10.0.0.1","port":443,"country":"AnotherCountry","cve":["CVE-2022-5678"]}
{"ip":"172.16.0.1","port":8080,"country":"TestCountry","cve":["CVE-2023-9876"]}`
	_, err = tempFile.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	tempFile.Close()

	// Set up test parameters
	query := ".country == \"TestCountry\""
	jq, err := gojq.Parse(fmt.Sprintf("select(%s)", query))
	if err != nil {
		t.Fatalf("Failed to parse query: %v", err)
	}
	code, err := gojq.Compile(jq)
	if err != nil {
		t.Fatalf("Failed to compile query: %v", err)
	}

	resultsChan := make(chan Entry, 10)
	errorsChan := make(chan error, 10)
	var wg sync.WaitGroup
	wg.Add(1)

	stats := &Stats{}

	// Run the function we're testing
	go processFile(tempFile.Name(), query, code, resultsChan, errorsChan, &wg, stats)

	// Wait for the function to complete
	wg.Wait()
	close(resultsChan)
	close(errorsChan)

	// Check for errors
	for err := range errorsChan {
		t.Errorf("Unexpected error: %v", err)
	}

	// Collect and check results
	var results []Entry
	for entry := range resultsChan {
		results = append(results, entry)
	}

	// Verify the results
	assert.Len(t, results, 2, "Expected 2 matching entries")
	assert.Equal(t, "192.168.1.1", results[0].IP, "Unexpected IP for first result")
	assert.Equal(t, "172.16.0.1", results[1].IP, "Unexpected IP for second result")

	// Verify the stats
	assert.Equal(t, int64(1), stats.TotalFiles, "Unexpected total files count")
	assert.Equal(t, int64(3), stats.TotalLines, "Unexpected total lines count")
	assert.Equal(t, int64(2), stats.MatchedLines, "Unexpected matched lines count")
}
