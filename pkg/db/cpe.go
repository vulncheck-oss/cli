package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/vulncheck-oss/cli/pkg/cpe/cpeutils"
	_ "modernc.org/sqlite"
)

// processWildcard converts asterisks at the beginning or end of a string to SQL wildcards
func processWildcard(value string) (string, bool) {
	value = strings.ToLower(value)
	hasWildcard := false
	
	if strings.HasPrefix(value, "*") {
		value = "%" + value[1:]
		hasWildcard = true
	}
	if strings.HasSuffix(value, "*") {
		value = value[:len(value)-1] + "%"
		hasWildcard = true
	}
	
	return value, hasWildcard
}

func CPESearch(indexName string, cpe cpeutils.CPE) ([]cpeutils.CPEVulnerabilities, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	tableName := strings.ReplaceAll(indexName, "-", "_")
	
	var conditions []string
	var args []interface{}

	if cpe.Vendor != "" && cpe.Vendor != "*" {
		vendor, hasWildcard := processWildcard(cpe.Vendor)
		operator := "="
		if hasWildcard {
			operator = "LIKE"
		}
		conditions = append(conditions, fmt.Sprintf("vendor %s ?", operator))
		args = append(args, vendor)
	}

	if cpe.Product != "" && cpe.Product != "*" {
		product, hasWildcard := processWildcard(cpe.Product)
		operator := "="
		if hasWildcard {
			operator = "LIKE"
		}
		conditions = append(conditions, fmt.Sprintf("product %s ?", operator))
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
