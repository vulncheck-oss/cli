package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"time"
)

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
	Duration time.Duration
}

func IPIntelSearch(indexName, country, asn, cidr, countryCode, hostname, id string) ([]IPEntry, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	// Convert table name to use underscores instead of hyphens
	tableName := strings.ReplaceAll(indexName, "-", "_")

	// Build query conditions
	var conditions []string
	var args []interface{}

	if country != "" {
		conditions = append(conditions, "country = ?")
		args = append(args, country)
	}
	if asn != "" {
		conditions = append(conditions, "asn = ?")
		args = append(args, asn)
	}
	if cidr != "" {
		conditions = append(conditions, "ip = ?")
		args = append(args, cidr)
	}
	if countryCode != "" {
		conditions = append(conditions, "country_code = ?")
		args = append(args, countryCode)
	}
	if hostname != "" {
		conditions = append(conditions, "json_array_contains(hostnames, ?)")
		args = append(args, hostname)
	}
	//if id != "" {
	//	conditions = append(conditions, "type_id = ?") // no longer type_id
	//	args = append(args, id)
	//}
	if id != "" {
		conditions = append(conditions, "json_extract(type, '$.id') = ?")
		args = append(args, id)
	}

	// Create WHERE clause if we have conditions
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Execute query
	query := fmt.Sprintf(`SELECT ip, port, ssl, lastSeen, asn, country, country_code, city, cve, matches, hostnames, type, feed_ids FROM "%s" %s`, tableName, whereClause)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []IPEntry
	for rows.Next() {
		var result IPEntry
		var cveJSON, matchesJSON, hostnamesJSON, feedIDsJSON []byte
		var typeJSON []byte

		err := rows.Scan(
			&result.IP,
			&result.Port,
			&result.SSL,
			&result.LastSeen,
			&result.ASN,
			&result.Country,
			&result.CountryCode,
			&result.City,
			&cveJSON,
			&matchesJSON,
			&hostnamesJSON,
			&typeJSON,
			&feedIDsJSON,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if err := json.Unmarshal(typeJSON, &result.Type); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal type: %w", err)
		}

		// Unmarshal JSON arrays
		if err := json.Unmarshal(cveJSON, &result.CVE); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal CVE: %w", err)
		}
		if err := json.Unmarshal(matchesJSON, &result.Matches); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal matches: %w", err)
		}
		if err := json.Unmarshal(hostnamesJSON, &result.Hostnames); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal hostnames: %w", err)
		}
		if err := json.Unmarshal(feedIDsJSON, &result.FeedIDs); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal feed IDs: %w", err)
		}

		results = append(results, result)
	}

	stats := &Stats{
		Duration: time.Since(startTime),
	}

	return results, stats, nil
}
