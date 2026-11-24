package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vulncheck-oss/cli/pkg/client"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	_ "modernc.org/sqlite"
)

func MetaByCVE(cve string) (*sdk.IndexVulncheckNvd2Response, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}

	tableName := strings.ReplaceAll("vulncheck-nvd2", "-", "_")

	query := fmt.Sprintf(
		`SELECT "id", "published", "lastModified", "date_added", "vulnStatus", "vulncheckKEVExploitAdd", "metrics", "weaknesses", "description", "references" FROM "%s" WHERE "id" = ? LIMIT 1`,
		tableName,
	)

	row := db.QueryRow(query, cve)

	var result client.ApiNVD20CVEExtended
	var metricsJSON []byte
	var weaknessesJSON []byte
	var referencesJSON []byte
	var vulncheckKEVExploitAdd *string
	var description *string
	var lastModified *string
	var dateAdded *string
	var vulnStatus *string

	err = row.Scan(
		&result.Id,
		&result.Published,
		&lastModified,
		&dateAdded,
		&vulnStatus,
		&vulncheckKEVExploitAdd,
		&metricsJSON,
		&weaknessesJSON,
		&description,
		&referencesJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	result.LastModified = lastModified
	result.DateAdded = dateAdded
	result.VulnStatus = vulnStatus

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

	if len(referencesJSON) > 0 {
		if err := json.Unmarshal(referencesJSON, &result.References); err != nil {
			return nil, fmt.Errorf("failed to unmarshal references: %w", err)
		}
	}

	var descriptionValue string
	if description != nil {
		descriptionValue = *description
	}

	response := &sdk.IndexVulncheckNvd2Response{
		Data:        []client.ApiNVD20CVEExtended{result},
		Description: descriptionValue,
	}

	return response, nil
}
