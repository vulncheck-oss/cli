package sdk

import (
	"encoding/json"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func metrics(t *testing.T, raw string) []advisoryMetric {
	t.Helper()
	var m []advisoryMetric
	assert.NoError(t, json.Unmarshal([]byte(raw), &m))
	return m
}

// CVE 5.x puts each CVSS version in its own metric object, so version
// preference has to span the slice. Taking the maximum instead renders a
// superseded score.
func TestBestCVSSPrefersTheNewestVersionAcrossEntries(t *testing.T) {
	tests := []struct {
		name  string
		raw   string
		want  float64
		found bool
	}{
		{
			// The live regression: cisa-csaf CVE-2019-13548.
			name:  "v3.1 supersedes a higher v3.0",
			raw:   `[{"cvssV3_0":{"baseScore":10}},{"cvssV3_1":{"baseScore":9.8}}]`,
			want:  9.8,
			found: true,
		},
		{
			name:  "v4.0 supersedes v3.1",
			raw:   `[{"cvssV3_1":{"baseScore":9.8}},{"cvssV4_0":{"baseScore":6.1}}]`,
			want:  6.1,
			found: true,
		},
		{
			name:  "order within the slice does not matter",
			raw:   `[{"cvssV3_1":{"baseScore":9.8}},{"cvssV3_0":{"baseScore":10}}]`,
			want:  9.8,
			found: true,
		},
		{
			name:  "highest wins within one version",
			raw:   `[{"cvssV3_1":{"baseScore":4.3}},{"cvssV3_1":{"baseScore":7.5}}]`,
			want:  7.5,
			found: true,
		},
		{
			name:  "falls back to v2.0 when it is all there is",
			raw:   `[{"cvssV2_0":{"baseScore":10}}]`,
			want:  10,
			found: true,
		},
		{
			name:  "epss carries no base score",
			raw:   `[{"format":"EPSS","other":{"content":{"score":0.7525},"type":"epss"}}]`,
			found: false,
		},
		{
			name:  "no metrics at all",
			raw:   `[]`,
			found: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score, found := bestCVSS(metrics(t, tt.raw))
			assert.Equal(t, tt.found, found)
			if tt.found {
				assert.Equal(t, tt.want, score)
			}
		})
	}
}

func TestParseAdvisoryRecordsProjectsAFullRecord(t *testing.T) {
	record := `{
	  "dataType":"CVE_RECORD","dataVersion":"5.2",
	  "cveMetadata":{"cveId":"CVE-2019-12255","dateUpdated":"2023-11-14T10:02:11.000Z","datePublished":"2019-08-09T00:00:00Z"},
	  "containers":{"cna":{
	    "providerMetadata":{"shortName":"siemens"},
	    "title":"Urgent/11 TCP/IP Stack Vulnerabilities",
	    "affected":[{"vendor":"Siemens"},{"vendor":"Siemens"}],
	    "references":[{"url":"a"},{"url":"b"},{"url":"c"}],
	    "metrics":[{"cvssV3_1":{"baseScore":9.8}}]
	  }}
	}`

	views := ParseAdvisoryRecords([]json.RawMessage{json.RawMessage(record)})
	assert.Len(t, views, 1)

	v := views[0]
	assert.Equal(t, "siemens", v.Feed)
	assert.Equal(t, "CVE-2019-12255", v.CveID)
	assert.Equal(t, "Urgent/11 TCP/IP Stack Vulnerabilities", v.Title)
	assert.Equal(t, "9.8", v.CVSS)
	assert.Equal(t, 2, v.Affected)
	assert.Equal(t, 3, v.References)
	assert.Equal(t, "2023-11-14", v.Updated)
}

// epss records have an empty title and nothing but metrics; ncsc-cves have only
// a title and references. Every field has to survive being absent.
func TestParseAdvisoryRecordsToleratesSparseRecords(t *testing.T) {
	views := ParseAdvisoryRecords([]json.RawMessage{
		json.RawMessage(`{"cveMetadata":{"cveId":"CVE-1"},"containers":{"cna":{"providerMetadata":{"shortName":"epss"},"title":"","metrics":[{"format":"EPSS"}]}}}`),
		json.RawMessage(`{"cveMetadata":{"cveId":"CVE-2"},"containers":{"cna":{}}}`),
		json.RawMessage(`{}`),
	})

	assert.Len(t, views, 3)
	assert.Equal(t, "epss", views[0].Feed)
	assert.Empty(t, views[0].Title)
	assert.Empty(t, views[0].CVSS)
	assert.Zero(t, views[0].Affected)
	assert.Equal(t, "CVE-2", views[1].CveID)
	assert.Empty(t, views[2].CveID)
}

// A title-less record falls back to its first non-empty description.
func TestParseAdvisoryRecordsFallsBackToDescription(t *testing.T) {
	views := ParseAdvisoryRecords([]json.RawMessage{
		json.RawMessage(`{"containers":{"cna":{"descriptions":[{"value":""},{"value":"a directory traversal in tvs.php"}]}}}`),
	})
	assert.Equal(t, "a directory traversal in tvs.php", views[0].Title)
}

func TestParseAdvisoryRecordsUpdatedFallbackChain(t *testing.T) {
	tests := []struct {
		name   string
		record string
		want   string
	}{
		{
			name:   "dateUpdated wins",
			record: `{"cveMetadata":{"dateUpdated":"2023-11-14T00:00:00Z","datePublished":"2019-01-01T00:00:00Z"},"containers":{"cna":{"providerMetadata":{"dateUpdated":"2020-01-01T00:00:00Z"}}}}`,
			want:   "2023-11-14",
		},
		{
			name:   "datePublished next",
			record: `{"cveMetadata":{"datePublished":"2019-01-01T00:00:00Z"},"containers":{"cna":{"providerMetadata":{"dateUpdated":"2020-01-01T00:00:00Z"}}}}`,
			want:   "2019-01-01",
		},
		{
			name:   "providerMetadata last",
			record: `{"containers":{"cna":{"providerMetadata":{"dateUpdated":"2020-01-01T00:00:00Z"}}}}`,
			want:   "2020-01-01",
		},
		{
			name:   "nothing at all",
			record: `{"containers":{"cna":{}}}`,
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			views := ParseAdvisoryRecords([]json.RawMessage{json.RawMessage(tt.record)})
			assert.Equal(t, tt.want, views[0].Updated)
		})
	}
}

// A record that will not decode is skipped rather than failing the page: the
// table is a convenience view and --json still carries everything.
func TestParseAdvisoryRecordsSkipsUndecodableRecords(t *testing.T) {
	views := ParseAdvisoryRecords([]json.RawMessage{
		json.RawMessage(`{"cveMetadata":{"cveId":"CVE-1"}}`),
		json.RawMessage(`not json`),
		json.RawMessage(`{"cveMetadata":{"cveId":"CVE-2"}}`),
	})
	assert.Len(t, views, 2)
	assert.Equal(t, "CVE-1", views[0].CveID)
	assert.Equal(t, "CVE-2", views[1].CveID)
}

// jvndb and bdu titles are not ASCII, so truncation has to land on a rune
// boundary or the cell contains a broken character.
func TestTruncateTitleCutsOnRuneBoundaries(t *testing.T) {
	long := strings.Repeat("脆", 100)
	got := truncateTitle(long, maxAdvisoryTitle)

	assert.Equal(t, maxAdvisoryTitle, len([]rune(got)))
	assert.True(t, strings.HasSuffix(got, "…"))
	assert.True(t, utf8.ValidString(got), "truncation must not split a rune")
}

func TestTruncateTitleCollapsesWhitespace(t *testing.T) {
	assert.Equal(t, "a b c", truncateTitle("  a\n\tb   c  ", maxAdvisoryTitle))
	assert.Equal(t, "short", truncateTitle("short", maxAdvisoryTitle))
}

func TestShortDate(t *testing.T) {
	assert.Equal(t, "2026-09-02", ShortDate("2026-09-02T16:17:49Z"))
	assert.Equal(t, "2026-09-02", ShortDate("2026-09-02T16:17:49.009938487Z"))
	// Feeds are inconsistent about precision, so anything unparseable is
	// passed through rather than dropped.
	assert.Equal(t, "2026-09-02", ShortDate("2026-09-02"))
	assert.Equal(t, "", ShortDate(""))
	assert.Equal(t, "nonsense", ShortDate("nonsense"))
}
