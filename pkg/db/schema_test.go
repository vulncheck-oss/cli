package db

import (
	"testing"
)

func TestGetSchema(t *testing.T) {
	// Test valid index names from different schemas
	testCases := []struct {
		indexName    string
		expectedName string
	}{
		{"ipintel-3d", "ipintel"},
		{"ipintel-10d", "ipintel"},
		{"ipintel-30d", "ipintel"},
		{"cargo", "purl PM"},
		{"npm", "purl PM"},
		{"golang", "purl PM"},
		{"alpine-purls", "purl OS"},
		{"rocky-purls", "purl OS"},
		{"cpecve", "cpecve"},
	}

	for _, tc := range testCases {
		t.Run("Index: "+tc.indexName, func(t *testing.T) {
			schema := GetSchema(tc.indexName)
			if schema == nil {
				t.Errorf("GetSchema(%q) returned nil, expected schema with name %q", tc.indexName, tc.expectedName)
				return
			}
			if schema.Name != tc.expectedName {
				t.Errorf("GetSchema(%q) returned schema with name %q, expected %q", tc.indexName, schema.Name, tc.expectedName)
			}
		})
	}

	// Test invalid index name should return fallback schema
	t.Run("Invalid index name", func(t *testing.T) {
		schema := GetSchema("non-existent-index")
		if schema == nil {
			t.Error("GetSchema with invalid index name returned nil, expected fallback schema")
			return
		}
		if !schema.Fallback {
			t.Errorf("GetSchema with invalid index name returned schema %q, expected fallback schema", schema.Name)
		}
		if schema.Name != "fallback" {
			t.Errorf("GetSchema with invalid index name returned schema %q, expected 'fallback'", schema.Name)
		}
	})
}

func TestSchemaStructure(t *testing.T) {
	// Test a few schemas to ensure their structures are as expected
	testCases := []struct {
		schemaName      string
		expectedColumns int
		testColumn      string
		columnIsJSON    bool
	}{
		{"ipintel", 15, "hostnames", true},
		{"purl PM", 6, "purl", true},
		{"cpecve", 12, "cves", true},
		{"fallback", 1, "data", true},
	}

	for _, tc := range testCases {
		t.Run("Schema: "+tc.schemaName, func(t *testing.T) {
			var schema *Schema
			// Find the schema by name
			for i := range Schemas {
				if Schemas[i].Name == tc.schemaName {
					schema = &Schemas[i]
					break
				}
			}

			if schema == nil {
				t.Errorf("Schema %q not found", tc.schemaName)
				return
			}

			if len(schema.Columns) != tc.expectedColumns {
				t.Errorf("Schema %q has %d columns, expected %d", tc.schemaName, len(schema.Columns), tc.expectedColumns)
			}

			// Test for specific column
			found := false
			for _, col := range schema.Columns {
				if col.Name == tc.testColumn {
					found = true
					if col.IsJSON != tc.columnIsJSON {
						t.Errorf("Column %q in schema %q has IsJSON=%v, expected %v",
							tc.testColumn, tc.schemaName, col.IsJSON, tc.columnIsJSON)
					}
					break
				}
			}

			if !found {
				t.Errorf("Column %q not found in schema %q", tc.testColumn, tc.schemaName)
			}
		})
	}
}
