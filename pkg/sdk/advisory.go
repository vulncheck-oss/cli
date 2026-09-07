package sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultAdvisoryLimit is the page size /v4/advisory applies when the
	// caller sends no limit. We mirror it locally so the result-window guard
	// below computes the same window the API would.
	DefaultAdvisoryLimit = 10

	// MaxAdvisoryLimit is the largest page size the API accepts; anything
	// higher is rejected with HTTP 400.
	MaxAdvisoryLimit = 100

	// maxAdvisoryWindow is OpenSearch's result window. Offset paging past it
	// fails with HTTP 503 wrapping an OpenSearch "Result window is too large"
	// 400, so we refuse the request rather than let that surface.
	maxAdvisoryWindow = 10000
)

// AdvisoryQueryParameters is the CLOSED set of parameters /v4/advisory accepts.
//
// The json tags name the wire parameters for reference only -- encoding goes
// through setAdvisoryQueryParameters, which is authoritative. Changing a tag
// here changes nothing that reaches the API.
//
// It is closed by construction on purpose: the API returns HTTP 200 with zero
// results for any parameter it does not recognise, rather than an error. Adding
// a field here without a matching API parameter silently empties every response
// that uses it.
type AdvisoryQueryParameters struct {
	// FILTERS
	Name            string `json:"name"` // advisory feed name; exposed as --feed
	CveID           string `json:"cve_id"`
	Vendor          string `json:"vendor"`
	Product         string `json:"product"`
	Platform        string `json:"platform"`
	Version         string `json:"version"`
	CPE             string `json:"cpe"`
	PackageName     string `json:"package_name"`
	PURL            string `json:"purl"`
	ReferenceURL    string `json:"reference_url"`
	ReferenceTag    string `json:"reference_tag"`
	DescriptionLang string `json:"description_lang"`
	UpdatedAfter    string `json:"updatedAfter"`
	UpdatedBefore   string `json:"updatedBefore"`

	// PAGINATION
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	StartCursor bool   `json:"start_cursor"`
	Cursor      string `json:"cursor"`
}

// filters returns every filter value, so callers can tell a filtered query from
// a whole-corpus one without enumerating the fields again.
func (q AdvisoryQueryParameters) filters() []string {
	return []string{
		q.Name, q.CveID, q.Vendor, q.Product, q.Platform, q.Version, q.CPE,
		q.PackageName, q.PURL, q.ReferenceURL, q.ReferenceTag,
		q.DescriptionLang, q.UpdatedAfter, q.UpdatedBefore,
	}
}

// HasFilter reports whether at least one filter is set.
func (q AdvisoryQueryParameters) HasFilter() bool {
	for _, f := range q.filters() {
		if f != "" {
			return true
		}
	}
	return false
}

// Validate rejects requests the API would answer misleadingly rather than with
// a useful error. Callers should run it before issuing a request; the request
// methods run it again so a bad query can never reach the wire.
func (q AdvisoryQueryParameters) Validate() error {
	if !q.HasFilter() {
		return fmt.Errorf("at least one filter is required; see 'vulncheck advisory list --help' " +
			"for the available filters, or 'vulncheck advisory feeds' for feed names")
	}
	if q.Limit > MaxAdvisoryLimit {
		return fmt.Errorf("limit must not exceed %d", MaxAdvisoryLimit)
	}
	if q.Limit < 0 || q.Page < 0 {
		return fmt.Errorf("limit and page must not be negative")
	}
	if q.Page > 0 && (q.StartCursor || q.Cursor != "") {
		return fmt.Errorf("page cannot be combined with cursor pagination; the API silently ignores page and returns the first page")
	}

	// The window guard has to substitute the API's own default, or the most
	// likely way to trip the ceiling -- a large --page with no --limit --
	// multiplies by zero and sails through to a 503.
	limit := q.Limit
	if limit == 0 {
		limit = DefaultAdvisoryLimit
	}
	page := q.Page
	if page == 0 {
		page = 1
	}
	if page*limit > maxAdvisoryWindow {
		return fmt.Errorf(
			"page %d at limit %d exceeds the API's %d result window; page beyond it with "+
				"--start-cursor and --cursor, or --all",
			page, limit, maxAdvisoryWindow)
	}
	return nil
}

// setAdvisoryQueryParameters encodes q into query. Only the parameters the API
// documents are ever added.
func setAdvisoryQueryParameters(query url.Values, q AdvisoryQueryParameters) {
	for key, value := range map[string]string{
		"name":             q.Name,
		"cve_id":           q.CveID,
		"vendor":           q.Vendor,
		"product":          q.Product,
		"platform":         q.Platform,
		"version":          q.Version,
		"cpe":              q.CPE,
		"package_name":     q.PackageName,
		"purl":             q.PURL,
		"reference_url":    q.ReferenceURL,
		"reference_tag":    q.ReferenceTag,
		"description_lang": q.DescriptionLang,
		"updatedAfter":     q.UpdatedAfter,
		"updatedBefore":    q.UpdatedBefore,
	} {
		if value != "" {
			query.Add(key, value)
		}
	}

	if q.Limit != 0 {
		query.Add("limit", fmt.Sprintf("%d", q.Limit))
	}
	if q.Page != 0 {
		query.Add("page", fmt.Sprintf("%d", q.Page))
	}
	// start_cursor activates on PRESENCE -- the API ignores the value, so
	// sending start_cursor=false would still switch the response into cursor
	// mode. It must be omitted entirely when false.
	if q.StartCursor {
		query.Add("start_cursor", "true")
	}
	if q.Cursor != "" {
		query.Add("cursor", q.Cursor)
		query.Del("start_cursor") // mutually exclusive; the cursor wins
	}
}

type AdvisoryMeta struct {
	Total      int    `json:"total"`
	Page       int    `json:"page"`
	Pages      int    `json:"pages"`
	Limit      int    `json:"limit"`
	Filtered   int    `json:"filtered"`
	NextCursor string `json:"next_cursor"`
}

// AdvisoryResponse holds records as raw JSON.
//
// Every record is a CVE Record Format 5.2 document. The CLI emits them verbatim
// for --json and reads only a handful of paths for the table, so typing the
// payload would buy nothing and would couple the CLI to the API's OpenAPI
// codegen. IndexResponse takes the same approach for /v3/index.
type AdvisoryResponse struct {
	Meta AdvisoryMeta      `json:"_meta"`
	Data []json.RawMessage `json:"data"`
}

// GetAdvisories https://docs.vulncheck.com/api/v4/advisory
func (c *Client) GetAdvisories(q AdvisoryQueryParameters) (responseJSON *AdvisoryResponse, err error) {
	if err := q.Validate(); err != nil {
		return nil, err
	}

	// c.Request reads c.Values but never resets them, and Query/Add append
	// rather than replace -- a reused client would otherwise send the previous
	// call's values alongside this one, and the API keys off the first.
	c.Values = &url.Values{}
	setAdvisoryQueryParameters(*c.Values, q)

	resp, err := c.Request("GET", "/v4/advisory")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&responseJSON)
	return responseJSON, nil
}

// GetAllAdvisories walks every page via cursor pagination, returning the
// combined records and the number of pages fetched.
//
// Cursor mode is the only way past the API's result window, and termination is
// not signalled by an empty next_cursor -- it stays populated on the final page
// of content. We stop on an empty page or a cursor that stops advancing.
//
// The page count is returned because it is the one fact about the walk that the
// combined result cannot express; the per-page metadata is deliberately not,
// since it describes a page the caller never sees.
func (c *Client) GetAllAdvisories(q AdvisoryQueryParameters) ([]json.RawMessage, int, error) {
	if err := q.Validate(); err != nil {
		return nil, 0, err
	}

	q.Page = 0
	q.Cursor = ""
	q.StartCursor = true

	// The page size is not observable in the combined result, so always walk at
	// the maximum. Inheriting the API's default costs an order of magnitude in
	// requests: when this was written, walking metasploit (~3k records) took
	// 57.5s at limit 10 against 6.2s at limit 100.
	q.Limit = MaxAdvisoryLimit

	response, err := c.GetAdvisories(q)
	if err != nil {
		return nil, 0, err
	}
	if response == nil {
		return []json.RawMessage{}, 0, nil
	}

	combined := append([]json.RawMessage{}, response.Data...)
	pages := 1
	cursor := response.Meta.NextCursor

	for cursor != "" {
		next := q
		next.StartCursor = false
		next.Cursor = cursor

		page, err := c.GetAdvisories(next)
		if err != nil {
			return nil, 0, err
		}
		pages++
		if page == nil || len(page.Data) == 0 {
			break
		}
		combined = append(combined, page.Data...)
		if page.Meta.NextCursor == cursor {
			break
		}
		cursor = page.Meta.NextCursor
	}

	return combined, pages, nil
}

type AdvisoryFeedMeta struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type AdvisoryFeedsResponse struct {
	Data []AdvisoryFeedMeta `json:"data"`
}

// GetAdvisoryFeeds https://docs.vulncheck.com/api/v4/advisory
//
// Note the endpoint is /v4/advisory/list: in the API "list" means the catalogue
// of feeds, not the records. The CLI spells this 'advisory feeds' so that
// 'advisory list' can keep meaning what 'index list' means.
func (c *Client) GetAdvisoryFeeds() (responseJSON *AdvisoryFeedsResponse, err error) {
	resp, err := c.ResetQuery().Request("GET", "/v4/advisory/list")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r AdvisoryResponse) String() string {
	return fmt.Sprintf("Meta: %v\nData: %v\n", r.Meta, r.Data)
}

// GetData returns the records, never nil, so an empty page marshals as [] and
// not null. Callers pipe this straight to --json.
func (r AdvisoryResponse) GetData() []json.RawMessage {
	if r.Data == nil {
		return []json.RawMessage{}
	}
	return r.Data
}

// Strings representation of the response
func (r AdvisoryFeedsResponse) String() string {
	return fmt.Sprintf("Data: %v\n", r.Data)
}

// GetData - Returns the data from the response
func (r AdvisoryFeedsResponse) GetData() []AdvisoryFeedMeta {
	return r.Data
}

// maxAdvisoryTitle bounds the title column so one verbose feed cannot wreck the
// table layout for every other row.
const maxAdvisoryTitle = 72

// AdvisoryRecordView is the flattened projection of a CVE Record Format 5.2
// document that the table renderer needs.
//
// It is decoded separately from the raw record so that --json keeps emitting the
// API's payload untouched. Every field under containers.cna is optional -- epss
// records carry an empty title and nothing but metrics, ncsc-cves carry only a
// title and references -- so each value falls back to a zero rather than
// assuming a shape.
type AdvisoryRecordView struct {
	Feed       string
	CveID      string
	Title      string
	CVSS       string
	Affected   int
	References int
	Updated    string
}

type advisoryCVSS struct {
	BaseScore float64 `json:"baseScore"`
}

// advisoryMetric is one entry in containers.cna.metrics. Non-CVSS metrics --
// EPSS and other vendor blobs arrive under "other" -- decode to a zero value
// and carry no base score.
type advisoryMetric struct {
	CvssV40 *advisoryCVSS `json:"cvssV4_0"`
	CvssV31 *advisoryCVSS `json:"cvssV3_1"`
	CvssV30 *advisoryCVSS `json:"cvssV3_0"`
	CvssV20 *advisoryCVSS `json:"cvssV2_0"`
}

// bestCVSS returns the base score from the newest CVSS version present anywhere
// in the slice.
//
// The version preference has to be applied across the whole slice, not within
// one entry: CVE 5.x puts each version in its own metric object, so a record
// carrying both v3.0 and v3.1 has two entries. Taking the maximum instead
// renders a superseded score: CVE-2019-13548 in cisa-csaf was observed carrying
// v3.0 10 alongside v3.1 9.8, where 9.8 is the authoritative one. Within a
// single version, the highest score wins.
func bestCVSS(metrics []advisoryMetric) (float64, bool) {
	for _, ofVersion := range []func(advisoryMetric) *advisoryCVSS{
		func(m advisoryMetric) *advisoryCVSS { return m.CvssV40 },
		func(m advisoryMetric) *advisoryCVSS { return m.CvssV31 },
		func(m advisoryMetric) *advisoryCVSS { return m.CvssV30 },
		func(m advisoryMetric) *advisoryCVSS { return m.CvssV20 },
	} {
		best, found := 0.0, false
		for _, m := range metrics {
			if c := ofVersion(m); c != nil && (!found || c.BaseScore > best) {
				best, found = c.BaseScore, true
			}
		}
		if found {
			return best, true
		}
	}
	return 0, false
}

type advisoryRecordWire struct {
	CveMetadata struct {
		CveID         string `json:"cveId"`
		DatePublished string `json:"datePublished"`
		DateUpdated   string `json:"dateUpdated"`
	} `json:"cveMetadata"`
	Containers struct {
		Cna struct {
			ProviderMetadata struct {
				ShortName   string `json:"shortName"`
				DateUpdated string `json:"dateUpdated"`
			} `json:"providerMetadata"`
			Title        string `json:"title"`
			Descriptions []struct {
				Value string `json:"value"`
			} `json:"descriptions"`
			Affected   []json.RawMessage `json:"affected"`
			References []json.RawMessage `json:"references"`
			Metrics    []advisoryMetric  `json:"metrics"`
		} `json:"cna"`
	} `json:"containers"`
}

// ParseAdvisoryRecords projects raw records for display. Records that cannot be
// decoded are skipped rather than failing the whole page: the table is a
// convenience view, and --json still carries everything.
func ParseAdvisoryRecords(records []json.RawMessage) []AdvisoryRecordView {
	views := make([]AdvisoryRecordView, 0, len(records))
	for _, raw := range records {
		var wire advisoryRecordWire
		if err := json.Unmarshal(raw, &wire); err != nil {
			continue
		}
		cna := wire.Containers.Cna

		title := cna.Title
		if title == "" {
			for _, d := range cna.Descriptions {
				if d.Value != "" {
					title = d.Value
					break
				}
			}
		}

		updated := wire.CveMetadata.DateUpdated
		if updated == "" {
			updated = wire.CveMetadata.DatePublished
		}
		if updated == "" {
			updated = cna.ProviderMetadata.DateUpdated
		}

		cvss := ""
		if score, ok := bestCVSS(cna.Metrics); ok {
			cvss = strconv.FormatFloat(score, 'f', -1, 64)
		}

		views = append(views, AdvisoryRecordView{
			Feed:       cna.ProviderMetadata.ShortName,
			CveID:      wire.CveMetadata.CveID,
			Title:      truncateTitle(title, maxAdvisoryTitle),
			CVSS:       cvss,
			Affected:   len(cna.Affected),
			References: len(cna.References),
			Updated:    ShortDate(updated),
		})
	}
	return views
}

// truncateTitle shortens on rune boundaries so multi-byte titles (jvndb and bdu
// are not ASCII) are not cut mid-character.
func truncateTitle(s string, limit int) string {
	s = strings.Join(strings.Fields(s), " ")
	runes := []rune(s)
	if len(runes) <= limit {
		return s
	}
	return string(runes[:limit-1]) + "…"
}

// ShortDate trims an RFC3339 timestamp to its date. Feeds are inconsistent
// about precision and offset, so anything unparseable is passed through
// untouched.
func ShortDate(value string) string {
	if value == "" {
		return ""
	}
	if t, err := time.Parse(time.RFC3339, value); err == nil {
		return t.Format("2006-01-02")
	}
	if len(value) >= 10 {
		return value[:10]
	}
	return value
}
