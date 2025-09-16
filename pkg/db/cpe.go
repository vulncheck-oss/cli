package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	_ "modernc.org/sqlite"
)

func CPESearch(indexName string, cpe cpeutils.CPE) ([]cpeutils.CPEVulnerabilities, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	// Convert table name to use underscores instead of hyphens
	tableName := strings.ReplaceAll(indexName, "-", "_")

	// Build query with strict matching or wildcard only when asterisk is present
	var conditions []string
	var args []interface{}

	if cpe.Vendor != "" && cpe.Vendor != "*" {
		vendor := strings.ToLower(cpe.Vendor)
		if strings.HasPrefix(vendor, "*") || strings.HasSuffix(vendor, "*") {
			vendor = strings.ReplaceAll(vendor, "*", "%")
			conditions = append(conditions, "vendor LIKE ?")
		} else {
			conditions = append(conditions, "vendor = ?")
		}
		args = append(args, vendor)
	}

	if cpe.Product != "" && cpe.Product != "*" {
		product := strings.ToLower(cpe.Product)
		if strings.HasPrefix(product, "*") || strings.HasSuffix(product, "*") {
			product = strings.ReplaceAll(product, "*", "%")
			conditions = append(conditions, "product LIKE ?")
		} else {
			conditions = append(conditions, "product = ?")
		}
		args = append(args, product)
	}

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
	defer func() {
		if err := rows.Close(); err != nil {
			_ = err
		}
	}()

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
