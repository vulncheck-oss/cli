package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	_ "modernc.org/sqlite"
)

type PurlEntry struct {
	Name            string                  `json:"name"`
	Version         string                  `json:"version"`
	Purl            []string                `json:"purl"`
	CVEs            []string                `json:"cves"`
	Vulnerabilities []sdk.PurlVulnerability `json:"vulnerabilities"`
}

func PURLSearch(indexName string, instance packageurl.PackageURL) ([]PurlEntry, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	tableName := strings.ReplaceAll(indexName, "-", "_")
	query := fmt.Sprintf("SELECT name, version, purl, cves, vulnerabilities FROM `%s` WHERE purl LIKE ?", tableName)

	rows, err := db.Query(query, "%"+instance.String()+"%")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []PurlEntry
	for rows.Next() {
		var result PurlEntry
		var purlJSON, cvesJSON, vulnerabilitiesJSON sql.NullString

		err := rows.Scan(
			&result.Name,
			&result.Version,
			&purlJSON,
			&cvesJSON,
			&vulnerabilitiesJSON,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if purlJSON.Valid {
			if err := json.Unmarshal([]byte(purlJSON.String), &result.Purl); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal purl: %w", err)
			}
		}
		if cvesJSON.Valid {
			if err := json.Unmarshal([]byte(cvesJSON.String), &result.CVEs); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal cves: %w", err)
			}
		}
		if vulnerabilitiesJSON.Valid {
			if err := json.Unmarshal([]byte(vulnerabilitiesJSON.String), &result.Vulnerabilities); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal vulnerabilities: %w", err)
			}
		}

		// only add the result if Version !== Fixed
		if len(result.Vulnerabilities) > 0 && result.Version != result.Vulnerabilities[0].FixedVersion {
			results = append(results, result)
		}
	}

	stats := &Stats{
		Duration: time.Since(startTime),
	}

	return results, stats, nil
}
