package ipintel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/itchyny/gojq"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"os"
	"path/filepath"
	"slices"
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

type SearchStats struct {
	TotalFiles   int64
	TotalLines   int64
	MatchedLines int64
	Duration     time.Duration
}

func Command() *cobra.Command {
	var country, asn, cidr, countryCode, hostname, id string

	validTimeframes := []string{"3d", "10d", "30d"}

	cmd := &cobra.Command{
		Use:     "ipintel <3d|10d|30d> [flags]",
		Short:   "IP Intel offline search",
		Long:    "Search offline IP Intel data",
		Example: "vulncheck offline ipintel 3d --country=Sweden",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires a timeframe argument: the timeframe (3d, 10d, or 30d)\n\nUsage:\n  %s", cmd.UseLine())
			}
			if !slices.Contains(validTimeframes, args[0]) {
				return fmt.Errorf("invalid timeframe: %s. Must be one of %v\n\nUsage:\n  %s", args[0], validTimeframes, cmd.UseLine())
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			// we need to error if zero flags were specified
			if country == "" && asn == "" && cidr == "" && countryCode == "" && hostname == "" && id == "" {
				return fmt.Errorf("requires at least one filter flag\n\nUsage:\n  %s", cmd.UseLine())
			}

			indices, err := cache.CachedIndices()
			if err != nil {
				return err
			}

			index := indices.GetIndex(fmt.Sprintf("ipintel-%s", args[0]))

			if index == nil {
				return fmt.Errorf("index ipintel-%s is required for this command, and is not cached", args[0])
			}

			ui.Info(fmt.Sprintf("Searchning index %s, last updated on %s", index.Name, utils.ParseDate(index.LastUpdated)))

			query := buildQuery(country, asn, cidr, countryCode, hostname, id)

			results, stats, err := searchIndex(index.Name, query)
			if err != nil {
				return err
			}

			ui.Stat("Results found", fmt.Sprintf("%d", len(results)))
			ui.Stat("Files processed", fmt.Sprintf("%d", stats.TotalFiles))
			ui.Stat("Total entries", fmt.Sprintf("%d", stats.TotalLines))
			ui.Stat("Search duration", stats.Duration.String())

			for i, result := range results {
				if i >= 10 {
					break
				}
				ui.Info(fmt.Sprintf("%2d. IP: %-15s Country: %-15s ASN: %-10s ID: %s",
					i+1, result.IP, result.Country, result.ASN, result.Type.ID))
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&country, "country", "", "Filter by country")
	cmd.Flags().StringVar(&asn, "asn", "", "Filter by ASN")
	cmd.Flags().StringVar(&cidr, "cidr", "", "Filter by CIDR")
	cmd.Flags().StringVar(&countryCode, "country-code", "", "Filter by country code")
	cmd.Flags().StringVar(&hostname, "hostname", "", "Filter by hostname")
	cmd.Flags().StringVar(&id, "id", "", "Filter by ID")

	return cmd
}

func buildQuery(country, asn, cidr, countryCode, hostname, id string) string {
	conditions := []string{}

	if country != "" {
		conditions = append(conditions, fmt.Sprintf(".country == %q", country))
	}
	if asn != "" {
		conditions = append(conditions, fmt.Sprintf(".asn == %q", asn))
	}
	if cidr != "" {
		// Note: CIDR matching would require additional logic
		conditions = append(conditions, fmt.Sprintf(".ip | startswith(%q)", strings.Split(cidr, "/")[0]))
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

/*
func searchIndex(indexName, query string) ([]Entry, *SearchStats, error) {
	startTime := time.Now()
	var stats SearchStats

	configDir, err := config.IndicesDir()
	if err != nil {
		return nil, nil, err
	}

	indexDir := filepath.Join(configDir, indexName)
	files, err := os.ReadDir(indexDir)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read index directory: %w", err)
	}

	var results []Entry

	jq, err := gojq.Parse(fmt.Sprintf("select(%s)", query))
	code, err := gojq.Compile(jq)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse query: %w", err)
	}

	ui.Info(fmt.Sprintf("Using query: select(%s)", query)) // Add this line

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(indexDir, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			ui.Error(fmt.Sprintf("Failed to read file %s: %v", file.Name(), err))
			continue
		}

		lines := strings.Split(string(content), "\n")
		atomic.AddInt64(&stats.TotalFiles, 1)
		atomic.AddInt64(&stats.TotalLines, int64(len(lines)))

		for _, line := range lines {
			if line == "" {
				continue
			}

			var input map[string]any
			if err := json.Unmarshal([]byte(line), &input); err != nil {
				ui.Error(fmt.Sprintf("Error parsing JSON: %v", err))
				continue
			}

			iter := code.Run(input)
			v, ok := iter.Next()
			if !ok {
				continue
			}
			if err, ok := v.(error); ok {
				ui.Error(fmt.Sprintf("Error processing line: %v", err))
				continue
			}

			if _, ok := v.(map[string]interface{}); ok {
				var entry Entry
				if err := json.Unmarshal([]byte(line), &entry); err == nil {
					results = append(results, entry)
					atomic.AddInt64(&stats.MatchedLines, 1)
				}

			}
		}

	}

	stats.Duration = time.Since(startTime)

	return results, &stats, nil
}
*/

func searchIndex(indexName, query string) ([]Entry, *SearchStats, error) {
	startTime := time.Now()
	var stats SearchStats

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
	code, err := gojq.Compile(jq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse query: %w", err)
	}

	ui.Info(fmt.Sprintf("Using query: select(%s)", query))

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
		ui.Error(fmt.Sprintf("Error during processing: %v", err))
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

func processFile(filePath, query string, code *gojq.Code, resultsChan chan<- Entry, errorsChan chan<- error, wg *sync.WaitGroup, stats *SearchStats) {
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
	// Parse the query to extract key fields and values
	queryFields := parseQuery(query)

	// Check if any of the query fields match the JSON data
	for field, value := range queryFields {
		result := gjson.GetBytes(line, field)
		if result.Exists() && strings.Contains(strings.ToLower(result.String()), strings.ToLower(value)) {
			return true
		}
	}

	return false
}

func parseQuery(query string) map[string]string {
	fields := make(map[string]string)
	// This is a simple parser and should be adjusted based on your actual query format
	parts := strings.Split(query, " and ")
	for _, part := range parts {
		kv := strings.Split(part, " == ")
		if len(kv) == 2 {
			key := strings.Trim(kv[0], ". ")
			value := strings.Trim(kv[1], "\"")
			fields[key] = value
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
