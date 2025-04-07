package db

import (
	"strings"
	"testing"
)

func setupIPIntelTable(t *testing.T) {
	// Get the schema for ipintel-3d
	schema := GetSchema("ipintel-3d")
	if schema == nil {
		t.Fatalf("failed to get schema for ipintel-3d")
	}

	// Create the table explicitly with the correct name
	var columns []string
	for _, col := range schema.Columns {
		colDef := "\"" + col.Name + "\" " + col.Type
		if col.NotNull {
			colDef += " NOT NULL"
		}
		columns = append(columns, colDef)
	}

	// Use the correct table name
	tableDef := "CREATE TABLE IF NOT EXISTS ipintel_3d (" + strings.Join(columns, ", ") + ")"
	_, err := testDB.Exec(tableDef)
	if err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}
	
	// Insert test data with values for all required columns into the ipintel_3d table
	_, err = testDB.Exec(`INSERT INTO ipintel_3d (
		ip, country, asn, country_code, hostnames, type,
		port, ssl, lastSeen, city, cve, matches, type_kind, type_finding, feed_ids
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		"192.168.1.1", "United States", "AS12345", "US", 
		`["example.com", "test.com"]`, // hostnames json array
		`{"id": "web-server", "kind": "server", "finding": "nginx"}`, // type as json object
		80, true, "2023-01-01", "New York", 
		`["CVE-2023-1234"]`, // cve json array 
		`["nginx/1.18.0"]`, // matches json array
		"server", "nginx", 
		`["feed1", "feed2"]`) // feed_ids json array
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}
}

func TestIPIntelSearch(t *testing.T) {
	// Setup the test table
	setupIPIntelTable(t)

	// Define test cases
	tests := []struct {
		name         string
		country      string
		asn          string
		cidr         string
		countryCode  string
		hostname     string
		id           string
		expectHits   int
	}{
		{
			name:        "Match country",
			country:     "United States",
			expectHits:  1,
		},
		{
			name:        "Match ASN",
			asn:         "AS12345",
			expectHits:  1,
		},
		{
			name:        "Match IP",
			cidr:        "192.168.1.1",
			expectHits:  1,
		},
		{
			name:        "Match country code",
			countryCode: "US",
			expectHits:  1,
		},
		{
			name:        "Match hostname",
			hostname:    "example.com",
			expectHits:  1,
		},
		{
			name:        "Match type ID",
			id:          "web-server",
			expectHits:  1,
		},
		{
			name:        "No match",
			country:     "Canada",
			expectHits:  0,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, stats, err := IPIntelSearch("ipintel-3d", tt.country, tt.asn, tt.cidr, tt.countryCode, tt.hostname, tt.id)
			if err != nil {
				t.Fatalf("IPIntelSearch failed: %v", err)
			}

			if len(results) != tt.expectHits {
				t.Errorf("expected %d hits, got %d", tt.expectHits, len(results))
			}

			if stats == nil || stats.Duration <= 0 {
				t.Error("expected valid stats with non-zero duration")
			}
		})
	}
}