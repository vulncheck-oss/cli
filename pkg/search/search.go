package search

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/itchyny/gojq"
	"github.com/package-url/packageurl-go"
	"github.com/tidwall/gjson"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk-go"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type PurlEntry struct {
	Name            string                  `json:"name"`
	Version         string                  `json:"version"`
	Purl            []string                `json:"purl"`
	Licenses        []string                `json:"licenses"`
	CVEs            []string                `json:"cves"`
	Vulnerabilities []sdk.PurlVulnerability `json:"vulnerabilities"`
	Artifacts       struct {
		Source []struct {
			Type      string `json:"type"`
			URL       string `json:"url"`
			Reference string `json:"reference,omitempty"`
		} `json:"source"`
		Binary []struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"binary"`
	} `json:"artifacts"`
	PublishedDate string `json:"published_date"`
}

type IPEntry struct {
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

func IndexCPE(indexName string, cpe cpeutils.CPE, query string) ([]cpeutils.CPEVulnerabilities, *Stats, error) {
	startTime := time.Now()
	var stats Stats

	configDir, err := config.IndicesDir()
	if err != nil {
		return nil, nil, err
	}

	indexDir := filepath.Join(configDir, indexName)
	files, err := listJSONFiles(indexDir)
	if err != nil || len(files) == 0 {
		return nil, nil, fmt.Errorf("failed to find JSON files in index directory %s: %w", indexDir, err)
	}

	filePath := files[0]
	stats.TotalFiles = 1
	stats.Query = query

	// Compile the jq query
	jq, err := gojq.Parse(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse query: %w", err)
	}
	code, err := gojq.Compile(jq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to compile query: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var results []cpeutils.CPEVulnerabilities

	for scanner.Scan() {
		stats.TotalLines++
		line := scanner.Bytes()

		var entry cpeutils.CPEVulnerabilities
		if err := json.Unmarshal(line, &entry); err != nil {
			continue
		}

		// Apply the query using gojq
		var input interface{}
		if err := json.Unmarshal(line, &input); err != nil {
			continue
		}

		iter := code.Run(input)
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			if err, ok := v.(error); ok {
				return nil, nil, fmt.Errorf("query execution error: %w", err)
			}
			if v == true {
				results = append(results, entry)
				stats.MatchedLines++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	stats.Duration = time.Since(startTime)
	return results, &stats, nil
}

func quickFilterCPE(line []byte, cpe cpeutils.CPE) bool {
	if cpe.Vendor != "" && gjson.GetBytes(line, "vendor").String() != cpe.Vendor {
		return false
	}
	if cpe.Product != "" && gjson.GetBytes(line, "product").String() != cpe.Product {
		return false
	}
	return true
}

func QueryIPIntel(country, asn, cidr, countryCode, hostname, id string) string {
	var conditions []string

	if country != "" {
		conditions = append(conditions, fmt.Sprintf(".country == %q", country))
	}
	if asn != "" {
		conditions = append(conditions, fmt.Sprintf(".asn == %q", asn))
	}
	if cidr != "" {
		// Note: CIDR matching would require additional logic
		conditions = append(conditions, fmt.Sprintf(".ip == %q", cidr))
		// conditions = append(conditions, fmt.Sprintf(".ip | startswith(%q)", strings.Split(cidr, "/")[0]))
	}
	if countryCode != "" {
		conditions = append(conditions, fmt.Sprintf(".country_code == %q", countryCode))
	}
	if hostname != "" {
		conditions = append(conditions, fmt.Sprintf(".hostnames | any(. == %q)", hostname))
	}
	if id != "" {
		conditions = append(conditions, fmt.Sprintf(".type.id == %q", id))
	}

	if len(conditions) == 0 {
		return "true"
	}
	return strings.Join(conditions, " and ")
}

func QueryPURL(instance packageurl.PackageURL) string {
	seperator := "/"
	var conditions []string

	if instance.Type == "maven" {
		seperator = ":"
	}

	if instance.Namespace == "alpine" {
		conditions = append(conditions, fmt.Sprintf(".package_name == %q", instance.Name))
	} else {
		if instance.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(".name == %q", fmt.Sprintf("%s%s%s", instance.Namespace, seperator, instance.Name)))
		} else {
			conditions = append(conditions, fmt.Sprintf(".name == %q", instance.Name))
		}
	}

	if instance.Version != "" {
		if instance.Type == "golang" {
			// For golang, match the version at the end of the string
			// conditions = append(conditions, fmt.Sprintf(".version | contains(%q)", instance.Version))
			conditions = append(conditions, fmt.Sprintf("(.version | index(%q)) != null and (.version | rindex(%q)) == (.version | length - %d)", instance.Version, instance.Version, len(instance.Version)))
		} else {
			// For other types, keep the contains check
			conditions = append(conditions, fmt.Sprintf(".version == %q", instance.Version))
		}
	}
	return strings.Join(conditions, " and ")
}

func IPIndex(indexName, query string) ([]IPEntry, *Stats, error) {
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

	resultsChan := make(chan IPEntry)
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

	var results []IPEntry
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

// IndexPurl - for PURL searches
func IndexPurl(indexName, query string) ([]PurlEntry, *Stats, error) {
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

	resultsChan := make(chan PurlEntry)
	errorsChan := make(chan error)
	var wg sync.WaitGroup

	// Start worker goroutines
	for _, file := range files {
		wg.Add(1)
		go processPurlFile(file, query, code, resultsChan, errorsChan, &wg, &stats)
	}

	// Collect results
	go func() {
		wg.Wait()
		close(resultsChan)
		close(errorsChan)
	}()

	var results []PurlEntry
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

// New function to process PURL files
func processPurlFile(filePath, query string, code *gojq.Code, resultsChan chan<- PurlEntry, errorsChan chan<- error, wg *sync.WaitGroup, stats *Stats) {
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

		entry, err := processPurlLine(line, code)
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

// New function to process a single PURL line
func processPurlLine(line []byte, code *gojq.Code) (*PurlEntry, error) {
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
		var entry PurlEntry
		if err := json.Unmarshal(line, &entry); err != nil {
			return nil, fmt.Errorf("error unmarshaling entry: %w", err)
		}
		return &entry, nil
	}

	return nil, nil
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

func processFile(filePath, query string, code *gojq.Code, resultsChan chan<- IPEntry, errorsChan chan<- error, wg *sync.WaitGroup, stats *Stats) {
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

func fullProcess(line []byte, code *gojq.Code) (*IPEntry, error) {
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
		var entry IPEntry
		if err := json.Unmarshal(line, &entry); err != nil {
			return nil, fmt.Errorf("error unmarshaling entry: %w", err)
		}
		return &entry, nil
	}

	return nil, nil
}

func listJSONFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
