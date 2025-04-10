package db

import (
	"encoding/json"
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	_ "modernc.org/sqlite"
	"strings"
	"time"
)

func CPESearch(indexName string, cpe cpeutils.CPE) ([]cpeutils.CPEVulnerabilities, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	// Convert table name to use underscores instead of hyphens
	tableName := strings.ReplaceAll(indexName, "-", "_")

	// Build query based on vendor and product like search.matchesCPE
	var conditions []string
	var args []interface{}

	if cpe.Vendor != "" {
		conditions = append(conditions, "vendor LIKE ?")
		args = append(args, "%"+strings.ToLower(cpe.Vendor)+"%")
	}

	if cpe.Product != "" {
		conditions = append(conditions, "product LIKE ?")
		args = append(args, "%"+strings.ToLower(cpe.Product)+"%")
	}

	// Create WHERE clause if we have conditions
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Execute query
	query := fmt.Sprintf(`SELECT "vendor", "product", "version", "update", "edition", "language", "sw_edition", "target_sw", "target_hw", "other", "cpe23Uri", "cves" FROM "%s" %s`, tableName, whereClause)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []cpeutils.CPEVulnerabilities
	for rows.Next() {
		var result cpeutils.CPEVulnerabilities
		var cvesJSON []byte

		err := rows.Scan(
			&result.Vendor,
			&result.Product,
			&result.Version,
			&result.Update,
			&result.Edition,
			&result.Language,
			&result.SoftwareEdition,
			&result.TargetSoftware,
			&result.TargetHardware,
			&result.Other,
			&result.CPE23URI,
			&cvesJSON,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if err := json.Unmarshal(cvesJSON, &result.Cves); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal CVEs: %w", err)
		}

		results = append(results, result)
	}

	stats := &Stats{
		Duration: time.Since(startTime),
	}

	return results, stats, nil
}
