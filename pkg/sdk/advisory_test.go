package sdk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// knownAdvisoryParams is every parameter /v4/advisory documents. The API answers
// HTTP 200 with zero results for anything else, so a stray parameter is silent
// data loss rather than an error -- this list is the guard against that.
var knownAdvisoryParams = map[string]bool{
	"name": true, "cve_id": true, "vendor": true, "product": true,
	"platform": true, "version": true, "cpe": true, "package_name": true,
	"purl": true, "reference_url": true, "reference_tag": true,
	"description_lang": true, "updatedAfter": true, "updatedBefore": true,
	"page": true, "limit": true, "start_cursor": true, "cursor": true,
}

func advisoryStub(t *testing.T, body string, seen *[]url.Values) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if seen != nil {
			*seen = append(*seen, r.URL.Query())
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, body)
	}))
	t.Cleanup(srv.Close)
	return srv
}

const emptyAdvisoryBody = `{"data":[],"_meta":{"total":0,"page":1,"pages":0,"limit":10,"filtered":0}}`

func TestGetAdvisoriesSendsOnlyKnownParameters(t *testing.T) {
	var seen []url.Values
	srv := advisoryStub(t, emptyAdvisoryBody, &seen)

	_, err := Connect(srv.URL, "tok").GetAdvisories(AdvisoryQueryParameters{
		Name: "ghsa", CveID: "CVE-2019-12255", Vendor: "Siemens",
		Product: "SIPROTEC", Platform: "linux", Version: "1.0",
		CPE: "cpe:2.3:a:x:y", PackageName: "pkg", PURL: "pkg:npm/x",
		ReferenceURL: "https://example.com", ReferenceTag: "patch",
		DescriptionLang: "en", UpdatedAfter: "now-7d", UpdatedBefore: "now",
		Limit: 50, Page: 2,
	})
	assert.NoError(t, err)
	assert.Len(t, seen, 1)

	var unknown []string
	for key := range seen[0] {
		if !knownAdvisoryParams[key] {
			unknown = append(unknown, key)
		}
	}
	sort.Strings(unknown)
	assert.Empty(t, unknown, "unknown parameters silently zero the result set")

	assert.Equal(t, "ghsa", seen[0].Get("name"))
	assert.Equal(t, "CVE-2019-12255", seen[0].Get("cve_id"))
	assert.Equal(t, "pkg", seen[0].Get("package_name"))
	assert.Equal(t, "now-7d", seen[0].Get("updatedAfter"))
	assert.Equal(t, "50", seen[0].Get("limit"))
	assert.Equal(t, "2", seen[0].Get("page"))
}

// A reused client must not accumulate values: url.Values.Add appends, and the
// API keys off the first occurrence of a repeated parameter.
func TestGetAdvisoriesDoesNotAccumulateParametersAcrossCalls(t *testing.T) {
	var seen []url.Values
	srv := advisoryStub(t, emptyAdvisoryBody, &seen)
	client := Connect(srv.URL, "tok")

	_, err := client.GetAdvisories(AdvisoryQueryParameters{Name: "ghsa"})
	assert.NoError(t, err)
	_, err = client.GetAdvisories(AdvisoryQueryParameters{Name: "epss"})
	assert.NoError(t, err)

	assert.Len(t, seen, 2)
	assert.Equal(t, []string{"epss"}, seen[1]["name"])
}

// start_cursor activates on presence -- the API ignores the value, so sending
// start_cursor=false would still switch the response into cursor mode.
func TestStartCursorIsOmittedWhenFalse(t *testing.T) {
	var seen []url.Values
	srv := advisoryStub(t, emptyAdvisoryBody, &seen)
	client := Connect(srv.URL, "tok")

	_, err := client.GetAdvisories(AdvisoryQueryParameters{Name: "ghsa", StartCursor: false})
	assert.NoError(t, err)
	_, hasKey := seen[0]["start_cursor"]
	assert.False(t, hasKey, "start_cursor must be absent, not false")

	_, err = client.GetAdvisories(AdvisoryQueryParameters{Name: "ghsa", StartCursor: true})
	assert.NoError(t, err)
	assert.Equal(t, "true", seen[1].Get("start_cursor"))
}

func TestCursorSupersedesStartCursor(t *testing.T) {
	var seen []url.Values
	srv := advisoryStub(t, emptyAdvisoryBody, &seen)

	_, err := Connect(srv.URL, "tok").GetAdvisories(AdvisoryQueryParameters{
		Name: "ghsa", StartCursor: true, Cursor: "abc",
	})
	assert.NoError(t, err)
	_, hasKey := seen[0]["start_cursor"]
	assert.False(t, hasKey)
	assert.Equal(t, "abc", seen[0].Get("cursor"))
}

// Every rejected query must be rejected before a request is issued.
func TestValidationRejectsBeforeAnyRequest(t *testing.T) {
	tests := []struct {
		name  string
		query AdvisoryQueryParameters
		want  string
	}{
		{
			name:  "no filter would walk the whole corpus",
			query: AdvisoryQueryParameters{Limit: 10},
			want:  "at least one filter is required",
		},
		{
			name:  "page with start_cursor is silently ignored by the API",
			query: AdvisoryQueryParameters{Name: "ghsa", Page: 3, StartCursor: true},
			want:  "cannot be combined with cursor pagination",
		},
		{
			name:  "page with an explicit cursor",
			query: AdvisoryQueryParameters{Name: "ghsa", Page: 3, Cursor: "abc"},
			want:  "cannot be combined with cursor pagination",
		},
		{
			name:  "limit above the API maximum",
			query: AdvisoryQueryParameters{Name: "ghsa", Limit: 101},
			want:  "must not exceed 100",
		},
		{
			name:  "result window, explicit limit",
			query: AdvisoryQueryParameters{Name: "ghsa", Limit: 100, Page: 101},
			want:  "result window",
		},
		{
			name:  "result window, small limit",
			query: AdvisoryQueryParameters{Name: "ghsa", Limit: 1, Page: 10001},
			want:  "result window",
		},
		{
			// The likeliest way to trip the ceiling: without substituting the
			// API's default of 10, page*limit is 0 and this sails through.
			name:  "result window with limit unset",
			query: AdvisoryQueryParameters{Name: "ghsa", Page: 1001},
			want:  "result window",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requests := 0
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				requests++
				_, _ = fmt.Fprint(w, emptyAdvisoryBody)
			}))
			defer srv.Close()

			_, err := Connect(srv.URL, "tok").GetAdvisories(tt.query)
			assert.ErrorContains(t, err, tt.want)
			assert.Zero(t, requests, "the guard must fire before any HTTP request")
		})
	}
}

func TestValidationAcceptsTheBoundaries(t *testing.T) {
	assert.NoError(t, AdvisoryQueryParameters{Name: "ghsa", Limit: 100, Page: 100}.Validate())
	assert.NoError(t, AdvisoryQueryParameters{Name: "ghsa", Limit: 10, Page: 1000}.Validate())
	assert.NoError(t, AdvisoryQueryParameters{Name: "ghsa", Page: 1000}.Validate())
	assert.NoError(t, AdvisoryQueryParameters{CveID: "CVE-2019-12255"}.Validate())
}

// next_cursor stays populated on the final page of content, so termination has
// to come from an empty page. This mirrors the live 50/50/50/9/0 shape.
func TestGetAllAdvisoriesWalksUntilAnEmptyPage(t *testing.T) {
	pages := []string{
		`{"data":[{"a":1},{"a":2}],"_meta":{"total":5,"next_cursor":"c1"}}`,
		`{"data":[{"a":3},{"a":4}],"_meta":{"total":5,"next_cursor":"c2"}}`,
		`{"data":[{"a":5}],"_meta":{"total":5,"next_cursor":"c3"}}`,
		`{"data":[],"_meta":{"total":5,"next_cursor":"c3"}}`,
	}
	var seen []url.Values
	call := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.URL.Query())
		body := `{"data":[],"_meta":{}}`
		if call < len(pages) {
			body = pages[call]
		}
		call++
		_, _ = fmt.Fprint(w, body)
	}))
	defer srv.Close()

	records, walked, err := Connect(srv.URL, "tok").GetAllAdvisories(AdvisoryQueryParameters{Name: "sigmahq-sigma-rules", Limit: 2})
	assert.NoError(t, err)
	assert.Len(t, records, 5)
	// Every request counts, terminator included: the point of the number is
	// what the walk cost.
	assert.Equal(t, 4, walked, "three content pages plus the empty terminator")
	assert.Len(t, seen, 4)

	assert.Equal(t, "true", seen[0].Get("start_cursor"))
	assert.Equal(t, "c1", seen[1].Get("cursor"))
	_, hasKey := seen[1]["start_cursor"]
	assert.False(t, hasKey, "start_cursor must not ride along with a cursor")
}

// The page size is invisible in the combined output, so a walk must always use
// the maximum. Inheriting the API default of 10 costs an order of magnitude in
// requests -- measured at 57.5s vs 6.2s over 3,167 records.
func TestGetAllAdvisoriesAlwaysWalksAtTheMaximumPageSize(t *testing.T) {
	for _, requested := range []int{0, 2, 100} {
		var seen []url.Values
		srv := advisoryStub(t, `{"data":[],"_meta":{}}`, &seen)

		_, _, err := Connect(srv.URL, "tok").GetAllAdvisories(AdvisoryQueryParameters{
			Name: "ghsa", Limit: requested,
		})
		assert.NoError(t, err)
		assert.Equal(t, "100", seen[0].Get("limit"), "requested limit %d", requested)
	}
}

// A cursor that stops advancing must terminate the walk rather than loop.
func TestGetAllAdvisoriesStopsOnARepeatedCursor(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls++
		_, _ = fmt.Fprint(w, `{"data":[{"a":1}],"_meta":{"next_cursor":"stuck"}}`)
	}))
	defer srv.Close()

	records, _, err := Connect(srv.URL, "tok").GetAllAdvisories(AdvisoryQueryParameters{Name: "ghsa"})
	assert.NoError(t, err)
	assert.Len(t, records, 2)
	assert.Equal(t, 2, calls)
}

// A body that does not decode is the only evidence that the response was not
// the one we asked for. Returning it as a nil result with a nil error would
// make a truncated page indistinguishable from "no matches".
func TestGetAdvisoriesSurfacesADecodeFailure(t *testing.T) {
	srv := advisoryStub(t, `{"data":[{"a":1}`, nil)

	response, err := Connect(srv.URL, "tok").GetAdvisories(AdvisoryQueryParameters{Name: "ghsa"})
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "decoding /v4/advisory response")
}

// --all is the one path where a broken response used to be invisible: the walk
// read a nil page as end-of-walk and reported a clean run of zero records.
func TestGetAllAdvisoriesFailsOnATruncatedFirstPage(t *testing.T) {
	srv := advisoryStub(t, `{"data":[{"a":1}`, nil)

	records, walked, err := Connect(srv.URL, "tok").GetAllAdvisories(AdvisoryQueryParameters{Name: "ghsa"})
	assert.ErrorContains(t, err, "decoding /v4/advisory response")
	assert.Nil(t, records)
	assert.Zero(t, walked)
}

// Mid-walk the stakes are higher still: page 1 decoded, so without this the
// caller receives a genuine but silently incomplete slice.
func TestGetAllAdvisoriesFailsOnATruncatedPageMidWalk(t *testing.T) {
	call := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		call++
		if call == 1 {
			_, _ = fmt.Fprint(w, `{"data":[{"a":1}],"_meta":{"next_cursor":"c1"}}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"data":[{"a":2}`)
	}))
	defer srv.Close()

	records, walked, err := Connect(srv.URL, "tok").GetAllAdvisories(AdvisoryQueryParameters{Name: "ghsa"})
	assert.ErrorContains(t, err, "decoding /v4/advisory response")
	assert.Nil(t, records, "a partial walk must not be returned as a complete one")
	assert.Zero(t, walked)
}

// A literal `null` decodes cleanly to no response at all. It is never a real
// empty result -- the API spells that {"data":[],"_meta":{...}}.
func TestGetAllAdvisoriesFailsOnANullBody(t *testing.T) {
	for _, page := range []string{"first", "mid-walk"} {
		call := 0
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			call++
			if page == "mid-walk" && call == 1 {
				_, _ = fmt.Fprint(w, `{"data":[{"a":1}],"_meta":{"next_cursor":"c1"}}`)
				return
			}
			_, _ = fmt.Fprint(w, `null`)
		}))

		records, walked, err := Connect(srv.URL, "tok").GetAllAdvisories(AdvisoryQueryParameters{Name: "ghsa"})
		assert.ErrorIs(t, err, ErrEmptyAdvisoryResponse, "%s page", page)
		assert.Nil(t, records)
		assert.Zero(t, walked)
		srv.Close()
	}
}

func TestGetAdvisoryFeedsDecodesTheCatalogue(t *testing.T) {
	srv := advisoryStub(t, `{"data":[{"name":"epss","href":"http://api.vulncheck.com/v4/advisory?name=epss"},{"name":"ghsa","href":"x"}]}`, nil)

	response, err := Connect(srv.URL, "tok").GetAdvisoryFeeds()
	assert.NoError(t, err)
	assert.Len(t, response.GetData(), 2)
	assert.Equal(t, "epss", response.GetData()[0].Name)
}

func TestAdvisoryResponseDecodesMeta(t *testing.T) {
	srv := advisoryStub(t, `{"data":[{"dataType":"CVE_RECORD"}],"_meta":{"total":39,"page":1,"pages":1,"limit":100,"filtered":39,"next_cursor":"abc"}}`, nil)

	response, err := Connect(srv.URL, "tok").GetAdvisories(AdvisoryQueryParameters{CveID: "CVE-2019-12255"})
	assert.NoError(t, err)
	assert.Equal(t, 39, response.Meta.Total)
	assert.Equal(t, "abc", response.Meta.NextCursor)
	assert.Len(t, response.GetData(), 1)
}
