package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFilter(t *testing.T) {
	testCases := []struct {
		name     string
		json     string
		query    string
		expected bool
	}{
		{"Match country", `{"country": "Narnia"}`, ".country == \"Narnia\"", true},
		{"Match port", `{"port": 443}`, ".port == 443", true},
		{"Match CVE", `{"cve": ["CVE-2021-36260"]}`, "any(.cve[] | . == \"CVE-2021-36260\")", true},
		{"No match", `{"country": "Neverland"}`, ".country == \"Narnia\"", false},
		{"Match nested field", `{"type": {"id": "initial-access"}}`, ".type.id == \"initial-access\"", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := quickFilter([]byte(tc.json), tc.query)
			assert.Equal(t, tc.expected, result, "Unexpected result for query: %s", tc.query)
		})
	}
}

func TestParseQuery(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected map[string]string
	}{
		{
			name:     "Single condition",
			query:    ".country == \"Narnia\"",
			expected: map[string]string{"country": "Narnia"},
		},
		{
			name:     "Multiple conditions",
			query:    ".country == \"Narnia\" and .port == 443",
			expected: map[string]string{"country": "Narnia", "port": "443"},
		},
		{
			name:     "Nested field",
			query:    ".type.id == \"initial-access\"",
			expected: map[string]string{"type.id": "initial-access"},
		},
		{
			name:     "Array contains",
			query:    "any(.cve[] | . == \"CVE-2021-36260\")",
			expected: map[string]string{"cve": "CVE-2021-36260"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := parseQuery(tc.query)
			assert.Equal(t, tc.expected, result)
		})
	}
}
