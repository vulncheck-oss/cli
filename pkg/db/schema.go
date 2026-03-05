package db

type Column struct {
	Name    string // Column name in database
	Type    string // SQL type (TEXT, INTEGER, etc)
	Index   bool   // Whether to create an index
	NotNull bool   // Whether column can be null
	IsJSON  bool   // Whether value is JSON array
}

type Schema struct {
	Indices  []string
	Name     string
	Fallback bool
	Results  bool // whether the JSON in each file is inside a "results" array
	Columns  []Column
}

var Schemas = []Schema{
	{
		Fallback: true,
		Name:     "fallback",
		Indices:  []string{},
		Columns: []Column{
			{Name: "data", Type: "TEXT", Index: true, NotNull: true, IsJSON: true},
		},
	},
	{
		Indices: []string{"vulncheck-nvd2"},
		Name:    "nvd",
		Results: true,
		Columns: []Column{
			{Name: "id", Type: "TEXT", Index: true, NotNull: true},
			{Name: "published", Type: "TEXT", Index: false, NotNull: false, IsJSON: false},
			{Name: "vulncheckKEVExploitAdd", Type: "TEXT", Index: false, NotNull: false, IsJSON: false},
			{Name: "metrics", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "weaknesses", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "description", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
		},
	},
	{
		Indices: []string{"ipintel-3d", "ipintel-10d", "ipintel-30d"},
		Name:    "ipintel",
		Columns: []Column{
			// Primary search fields - all indexed
			{Name: "ip", Type: "TEXT", Index: true, NotNull: true},
			{Name: "country", Type: "TEXT", Index: true, NotNull: true},
			{Name: "asn", Type: "TEXT", Index: true, NotNull: true},
			{Name: "country_code", Type: "TEXT", Index: true, NotNull: true},
			{Name: "hostnames", Type: "TEXT", Index: true, NotNull: false, IsJSON: true},
			{Name: "type", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},

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
	{
		Name: "purl PM",
		Indices: []string{
			"cargo", "golang", "cocoapods", "hex", "npm", "gem", "pypi", "maven", "nuget", "composer", "hackage", "cran", "pub", "conan", "swift", "go", "dub", "elixir", "julia", "luarocks", "opam", "r", "vcpkg",
		},
		Columns: []Column{
			{Name: "name", Type: "TEXT", Index: false, NotNull: true},
			{Name: "version", Type: "TEXT", Index: false, NotNull: true},
			{Name: "purl", Type: "TEXT", Index: true, NotNull: true, IsJSON: true},
			{Name: "licenses", Type: "TEXT", Index: false, NotNull: false, IsJSON: true}, // Add licenses
			{Name: "cves", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "vulnerabilities", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
		},
	},
	{
		Name: "purl OS",
		Indices: []string{
			"alpine-purls", "rocky-purls", "debian-purls", "ubuntu-purls",
		},
		Columns: []Column{
			{Name: "name", Type: "TEXT", Index: false, NotNull: true},
			{Name: "version", Type: "TEXT", Index: false, NotNull: true},
			{Name: "purl", Type: "TEXT", Index: true, NotNull: true, IsJSON: true},
			{Name: "licenses", Type: "TEXT", Index: false, NotNull: false, IsJSON: true}, // Add licenses
			{Name: "cves", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			{Name: "vulnerabilities", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
			/*
				{Name: "os_version", Type: "TEXT", Index: true, NotNull: true},
				{Name: "os_arch", Type: "TEXT", Index: true, NotNull: true},
				{Name: "id", Type: "TEXT", Index: true, NotNull: true},
				{Name: "title", Type: "TEXT", Index: false, NotNull: true},
				{Name: "cve", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
				{Name: "severity", Type: "TEXT", Index: true, NotNull: false},
				{Name: "type", Type: "TEXT", Index: true, NotNull: false},
				{Name: "description", Type: "TEXT", Index: false, NotNull: false},
				{Name: "references", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
				{Name: "packages", Type: "TEXT", Index: false, NotNull: true, IsJSON: true},
			*/
		},
	},

	{
		Name:    "cpecve",
		Indices: []string{"cpecve"},
		Columns: []Column{
			{Name: "vendor", Type: "TEXT", Index: true, NotNull: false},
			{Name: "product", Type: "TEXT", Index: true, NotNull: false},
			{Name: "version", Type: "TEXT", Index: true, NotNull: false},
			{Name: "update", Type: "TEXT", Index: true, NotNull: false},
			{Name: "edition", Type: "TEXT", Index: true, NotNull: false},
			{Name: "language", Type: "TEXT", Index: true, NotNull: false},
			{Name: "sw_edition", Type: "TEXT", Index: true, NotNull: false},
			{Name: "target_sw", Type: "TEXT", Index: true, NotNull: false},
			{Name: "target_hw", Type: "TEXT", Index: true, NotNull: false},
			{Name: "other", Type: "TEXT", Index: true, NotNull: false},
			{Name: "cpe23Uri", Type: "TEXT", Index: true, NotNull: false},
			{Name: "cves", Type: "TEXT", Index: false, NotNull: false, IsJSON: true},
		},
	},
}

func GetSchema(indexName string) *Schema {
	// First try to find a matching schema by name
	for _, schema := range Schemas {
		for _, index := range schema.Indices {
			if index == indexName {
				return &schema
			}
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
