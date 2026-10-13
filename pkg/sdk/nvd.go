package sdk

import (
	"encoding/json"
	"fmt"
	"io"
)

// Paths of the NVD API 2.0 compatible endpoints. They mirror NVD's own paths, under the
// /rest/json prefix NVD uses.
const (
	NvdCvesPath       = "/rest/json/cves/2.0"
	NvdCveHistoryPath = "/rest/json/cvehistory/2.0"
	NvdCpesPath       = "/rest/json/cpes/2.0"
	NvdCpeMatchPath   = "/rest/json/cpematch/2.0"
	NvdSourcePath     = "/rest/json/source/2.0"
)

// NvdParam is one query parameter for an NVD endpoint, sent exactly as given.
//
// The parameters are an ordered list rather than a struct because the CLI passes them
// through without interpreting them: the API validates every value and answers a bad one
// with a 400 naming the parameter. An empty Value is still sent, which matters for
// `cursor`: an empty cursor is how a caller opts in to cursor paging.
type NvdParam struct {
	Name  string
	Value string
}

// NvdResponse is one NVD API 2.0 response, kept byte for byte as the server sent it.
//
// The record array is under a different key per endpoint (vulnerabilities, cveChanges,
// products, matchStrings, sources), so it is looked up by name rather than decoded into
// a field.
type NvdResponse struct {
	raw    json.RawMessage
	fields map[string]json.RawMessage

	ResultsPerPage int    `json:"resultsPerPage"`
	StartIndex     int    `json:"startIndex"`
	TotalResults   int    `json:"totalResults"`
	NextCursor     string `json:"nextCursor"`
}

// GetNvd queries one NVD API 2.0 compatible endpoint.
//
// ResetQuery is mandatory, as for every other SDK method: c.Query appends, so a reused
// Client would otherwise carry the previous call's parameters into this one.
func (c *Client) GetNvd(path string, params []NvdParam) (*NvdResponse, error) {
	c.ResetQuery()
	for _, p := range params {
		c.Query(p.Name, p.Value)
	}

	resp, err := c.Request("GET", path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(LimitedBody(resp.Body))
	if err != nil {
		return nil, fmt.Errorf("reading %s response: %w", path, err)
	}

	response := &NvdResponse{raw: body}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, fmt.Errorf("decoding %s response: %w", path, err)
	}
	if err := json.Unmarshal(body, &response.fields); err != nil {
		return nil, fmt.Errorf("decoding %s response: %w", path, err)
	}
	return response, nil
}

// Raw returns the whole response as the server sent it.
func (r *NvdResponse) Raw() json.RawMessage { return r.raw }

// Records returns the record array under key, as the server sent it, or an empty array
// when the response has none.
func (r *NvdResponse) Records(key string) json.RawMessage {
	if records, ok := r.fields[key]; ok && len(records) > 0 && string(records) != "null" {
		return records
	}
	return json.RawMessage("[]")
}
