package db

import (
	"regexp"
)

type Column struct {
	Name    string // Column name in database
	Type    string // SQL type (TEXT, INTEGER, etc)
	Index   bool   // Whether to create an index
	NotNull bool   // Whether column can be null
	IsJSON  bool   // Whether value is JSON array
}

type Schema struct {
	// IndexMatch will be a regex that matches index names, like ipintel-(3|7|10)d / etc
	IndexMatch string
	Fallback   bool
	Columns    []Column
}

var Schemas = []Schema{
	{
		Fallback:   true,
		IndexMatch: "",
		Columns: []Column{
			{Name: "data", Type: "TEXT", Index: true, NotNull: true, IsJSON: true},
		},
	},
	{
		IndexMatch: `^ipintel-\d+d$`,
		Columns: []Column{
			// Primary search fields - all indexed
			{Name: "ip", Type: "TEXT", Index: true, NotNull: true},
			{Name: "country", Type: "TEXT", Index: true, NotNull: true},
			{Name: "asn", Type: "TEXT", Index: true, NotNull: true},
			{Name: "country_code", Type: "TEXT", Index: true, NotNull: true},
			{Name: "hostnames", Type: "TEXT", Index: true, NotNull: false, IsJSON: true},
			{Name: "type_id", Type: "TEXT", Index: true, NotNull: false},

			// Non-searched fields - no indexes needed
			{Name: "port", Type: "INTEGER", Index: false, NotNull: true},
			{Name: "ssl", Type: "BOOLEAN", Index: false, NotNull: true},
			{Name: "lastSeen", Type: "TEXT", Index: false, NotNull: true},
			{Name: "city", Type: "TEXT", Index: false, NotNull: false},
			{Name: "cve", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "matches", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "type_kind", Type: "TEXT", Index: false, NotNull: false},
			{Name: "type_finding", Type: "TEXT", Index: false, NotNull: false},
			{Name: "feed_ids", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
		},
	},
}

func GetSchema(indexName string) *Schema {
	// First try to find a matching schema
	for _, schema := range Schemas {
		if schema.Fallback {
			continue // Skip fallback during regex matching
		}
		matched, _ := regexp.MatchString(schema.IndexMatch, indexName)
		if matched {
			return &schema
		}
	}

	// If no match found, return the fallback schema
	for _, schema := range Schemas {
		if schema.Fallback {
			return &schema
		}
	}

	return nil
}
