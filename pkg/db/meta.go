package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vulncheck-oss/sdk-go"
	"github.com/vulncheck-oss/sdk-go/pkg/client"
	_ "modernc.org/sqlite"
)

func MetaByCVE(cve string) (*sdk.IndexVulncheckNvd2Response, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}

	tableName := strings.ReplaceAll("vulncheck-nvd2", "-", "_")

	query := fmt.Sprintf(`SELECT "id", "published", "vulncheckKEVExploitAdd", "metrics", "weaknesses" FROM "%s" WHERE "id" = ? LIMIT 1`, tableName)

	row := db.QueryRow(query, cve)

	var result client.ApiNVD20CVEExtended
	var metricsJSON []byte
	var weaknessesJSON []byte
	var vulncheckKEVExploitAdd *string

	err = row.Scan(
		&result.Id,
		&result.Published,
		&vulncheckKEVExploitAdd,
		&metricsJSON,
		&weaknessesJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if vulncheckKEVExploitAdd != nil && *vulncheckKEVExploitAdd != "" {
		result.VulncheckKEVExploitAdd = vulncheckKEVExploitAdd
	}

	if len(metricsJSON) > 0 {
		if err := json.Unmarshal(metricsJSON, &result.Metrics); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metrics: %w", err)
		}
	}

	if len(weaknessesJSON) > 0 {
		if err := json.Unmarshal(weaknessesJSON, &result.Weaknesses); err != nil {
			return nil, fmt.Errorf("failed to unmarshal weaknesses: %w", err)
		}
	}

	response := &sdk.IndexVulncheckNvd2Response{
		Data: []client.ApiNVD20CVEExtended{result},
	}

	return response, nil
}
