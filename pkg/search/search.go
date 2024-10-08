package search

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/itchyny/gojq"
	"github.com/tidwall/gjson"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Entry struct {
	IP          string   `json:"ip"`
	Port        int      `json:"port"`
	SSL         bool     `json:"ssl"`
	LastSeen    string   `json:"lastSeen"`
	ASN         string   `json:"asn"`
	Country     string   `json:"country"`
	CountryCode string   `json:"country_code"`
	City        string   `json:"city"`
	CVE         []string `json:"cve"`
	Matches     []string `json:"matches"`
	Hostnames   []string `json:"hostnames"`
	Type        struct {
		ID      string `json:"id"`
		Kind    string `json:"kind"`
		Finding string `json:"finding"`
	} `json:"type"`
	FeedIDs []string `json:"feed_ids"`
}

type Stats struct {
	TotalFiles   int64
	TotalLines   int64
	MatchedLines int64
	Duration     time.Duration
	Query        string
}

func Index(indexName, query string) ([]Entry, *Stats, error) {
	startTime := time.Now()
	var stats Stats

	configDir, err := config.IndicesDir()
	if err != nil {
		return nil, nil, err
	}

	indexDir := filepath.Join(configDir, indexName)

	files, err := listIndexFiles(indexDir)
	if err != nil {
		return nil, nil, err
	}

	jq, err := gojq.Parse(fmt.Sprintf("select(%s)", query))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse query: %w", err)
	}
	code, err := gojq.Compile(jq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to compile query: %w", err)
	}

	stats.Query = query

	resultsChan := make(chan Entry)
	errorsChan := make(chan error)
	var wg sync.WaitGroup

	// Start worker goroutines
	for _, file := range files {
		wg.Add(1)
		go processFile(file, query, code, resultsChan, errorsChan, &wg, &stats)
	}

	// Collect results
	go func() {
		wg.Wait()
		close(resultsChan)
		close(errorsChan)
	}()

	var results []Entry
	for result := range resultsChan {
		results = append(results, result)
	}

	// Check for errors
	for err := range errorsChan {
		_ = ui.Error(fmt.Sprintf("Error during processing: %v", err))
	}

	stats.Duration = time.Since(startTime)

	return results, &stats, nil
}

func listIndexFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func processFile(filePath, query string, code *gojq.Code, resultsChan chan<- Entry, errorsChan chan<- error, wg *sync.WaitGroup, stats *Stats) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		errorsChan <- fmt.Errorf("failed to open file %s: %w", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	atomic.AddInt64(&stats.TotalFiles, 1)

	for scanner.Scan() {
		line := scanner.Bytes()
		atomic.AddInt64(&stats.TotalLines, 1)

		if !quickFilter(line, query) {
			continue
		}

		entry, err := fullProcess(line, code)
		if err != nil {
			errorsChan <- fmt.Errorf("error processing line in file %s: %w", filePath, err)
			continue
		}
		if entry != nil {
			resultsChan <- *entry
			atomic.AddInt64(&stats.MatchedLines, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		errorsChan <- fmt.Errorf("error reading file %s: %w", filePath, err)
	}
}

func quickFilter(line []byte, query string) bool {
	queryFields := parseQuery(query)

	for field, value := range queryFields {
		result := gjson.GetBytes(line, field)
		if result.Exists() {
			if result.IsArray() {
				for _, item := range result.Array() {
					if item.String() == value {
						return true
					}
				}
			} else if result.String() == value {
				return true
			}
		}
	}

	return false
}

func parseQuery(query string) map[string]string {
	fields := make(map[string]string)
	if strings.HasPrefix(query, "any(") && strings.HasSuffix(query, ")") {
		// Handle array contains query
		query = strings.TrimPrefix(query, "any(")
		query = strings.TrimSuffix(query, ")")
		parts := strings.Split(query, " | ")
		if len(parts) == 2 {
			key := strings.Trim(parts[0], ".[]")
			value := strings.Trim(parts[1], ". =")
			value = strings.Trim(value, "\"")
			fields[key] = value
		}
	} else {
		// Handle regular equality queries
		parts := strings.Split(query, " and ")
		for _, part := range parts {
			kv := strings.Split(part, " == ")
			if len(kv) == 2 {
				key := strings.Trim(kv[0], ". ")
				value := strings.Trim(kv[1], "\"")
				fields[key] = value
			}
		}
	}
	return fields
}

func fullProcess(line []byte, code *gojq.Code) (*Entry, error) {
	var input map[string]interface{}
	if err := json.Unmarshal(line, &input); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	iter := code.Run(input)
	v, ok := iter.Next()
	if !ok {
		return nil, nil
	}
	if err, ok := v.(error); ok {
		return nil, fmt.Errorf("error processing line: %w", err)
	}

	if _, ok := v.(map[string]interface{}); ok {
		var entry Entry
		if err := json.Unmarshal(line, &entry); err != nil {
			return nil, fmt.Errorf("error unmarshaling entry: %w", err)
		}
		return &entry, nil
	}

	return nil, nil
}
