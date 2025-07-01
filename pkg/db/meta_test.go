package db

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMetaByCVE(t *testing.T) {
	schema := GetSchema("vulncheck-nvd2")
	if schema == nil {
		t.Fatal("failed to get schema for vulncheck-nvd2")
	}

	var columns []string
	for _, col := range schema.Columns {
		colDef := "\"" + col.Name + "\" " + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		columns = append(columns, colDef)
	}

	tableDef := "CREATE TABLE IF NOT EXISTS vulncheck_nvd2 (" + strings.Join(columns, ", ") + ")"
	if _, err := testDB.Exec(tableDef); err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	metricsData := map[string]interface{}{
		"cvssMetricV31": []map[string]interface{}{
			{
				"cvssData": map[string]interface{}{
					"baseScore": 7.5,
				},
			},
		},
	}
	metricsJSON, _ := json.Marshal(metricsData)

	weaknessesData := []map[string]interface{}{
		{
			"source": "nvd@nist.gov",
			"type":   "Primary",
			"description": []map[string]interface{}{
				{
					"lang":  "en",
					"value": "CWE-79",
				},
			},
		},
		{
			"source": "nvd@nist.gov",
			"type":   "Primary",
			"description": []map[string]interface{}{
				{
					"lang":  "en",
					"value": "CWE-89",
				},
			},
		},
	}
	weaknessesJSON, _ := json.Marshal(weaknessesData)

	kevData := "2021-12-10"

	_, err := testDB.Exec(`INSERT INTO vulncheck_nvd2 (id, published, vulncheckKEVExploitAdd, metrics, weaknesses) 
		VALUES (?, ?, ?, ?, ?)`,
		"CVE-2021-44228", "2021-12-10T10:15:00Z", kevData, string(metricsJSON), string(weaknessesJSON))
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	tests := []struct {
		name      string
		cve       string
		expectErr bool
	}{
		{
			name:      "Valid CVE",
			cve:       "CVE-2021-44228",
			expectErr: false,
		},
		{
			name:      "Non-existent CVE",
			cve:       "CVE-9999-9999",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MetaByCVE(tt.cve)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("expected non-nil result")
			}

			if len(result.Data) != 1 {
				t.Errorf("expected 1 result, got %d", len(result.Data))
			}

			if result.Data[0].Id == nil || *result.Data[0].Id != tt.cve {
				t.Errorf("expected CVE %s, got %v", tt.cve, result.Data[0].Id)
			}

			if result.Data[0].Published == nil || *result.Data[0].Published != "2021-12-10T10:15:00Z" {
				t.Error("published date mismatch")
			}

			if result.Data[0].VulncheckKEVExploitAdd == nil || *result.Data[0].VulncheckKEVExploitAdd != kevData {
				t.Error("KEV data mismatch")
			}

			if result.Data[0].Metrics == nil {
				t.Error("expected metrics to be populated")
			}

			if result.Data[0].Weaknesses == nil {
				t.Error("expected weaknesses to be populated")
			}
		})
	}
}
