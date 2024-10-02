package ipintel

import (
	"encoding/json"
	"fmt"
	"github.com/itchyny/gojq"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"os"
	"path/filepath"
	"slices"
	"strings"
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
