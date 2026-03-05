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

	tableName := strings.ReplaceAll(indexName, "-", "_")

	var conditions []string
	var args []interface{}

	// Add vendor condition with strict matching or wildcard support
	if cpe.Vendor != "" && cpe.Vendor != "*" {
		vendor := strings.ToLower(cpe.Vendor)
		operator := "="

		if strings.HasPrefix(vendor, "*") {
			vendor = "%" + vendor[1:]
			operator = "LIKE"
		}
		if strings.HasSuffix(vendor, "*") && len(vendor) > 1 {
			vendor = vendor[:len(vendor)-1] + "%"
			operator = "LIKE"
		}

		conditions = append(conditions, fmt.Sprintf("vendor %s ?", operator))
		args = append(args, vendor)
	}

	// Add product condition with strict matching or wildcard support
	if cpe.Product != "" && cpe.Product != "*" {
		product := strings.ToLower(cpe.Product)
		operator := "="

		if strings.HasPrefix(product, "*") {
			product = "%" + product[1:]
			operator = "LIKE"
		}
		if strings.HasSuffix(product, "*") && len(product) > 1 {
			product = product[:len(product)-1] + "%"
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
