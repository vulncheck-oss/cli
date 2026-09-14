package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type PurlMeta struct {
	Namespace  string   `json:"namespace"`
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Qualifiers []string `json:"qualifiers"`
	Subpath    string   `json:"subpath"`
	Type       string   `json:"type"`
}

type PurlVulnerability struct {
	Detection    string `json:"detection"`
	FixedVersion string `json:"fixed_version"`
}

type PurlData struct {
	Cves            []string            `json:"cves"`
	Vulnerabilities []PurlVulnerability `json:"vulnerabilities"`
}

type PurlResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		PurlMeta       PurlMeta `json:"purl_struct"`
		Timestamp      string   `json:"timestamp"`
		TotalDocuments float64  `json:"total_documents"`
	} `json:"_meta"`
	Data PurlData `json:"data"`
}

type BatchPurlData struct {
	Purl            string              `json:"purl"`
	PurlMeta        PurlMeta            `json:"purl_struct"`
	Cves            []string            `json:"cves"`
	Vulnerabilities []PurlVulnerability `json:"vulnerabilities"`
}

// UnprocessedPurl is a purl the API could not look up. Reason is passed through
// verbatim and should be treated as an open set; as of writing it is
// "unsupported_type" (a valid purl for an ecosystem VulnCheck does not index),
// "unparseable" (not a valid purl), or "unsupported_distro" (a distro-scoped
// purl whose distro qualifier is missing or unrecognised).
type UnprocessedPurl struct {
	Purl   string `json:"purl"`
	Reason string `json:"reason"`
}

type PurlsResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		Timestamp      string  `json:"timestamp"`
		TotalDocuments float64 `json:"total_documents"`
		TotalSubmitted float64 `json:"total_submitted"`

		// Absent on API versions predating partial results, where an unusable
		// purl failed the whole batch instead of being reported.
		Unprocessed []UnprocessedPurl `json:"unprocessed"`
	} `json:"_meta"`
	PurlData []BatchPurlData `json:"data"`
}

// GetPurl https://docs.vulncheck.com/api/purl
func (c *Client) GetPurl(purl string) (responseJSON *PurlResponse, err error) {
	// See GetCpe — reset before adding this call's query param so batch
	// callers don't accumulate stale `purl` values across iterations.
	resp, err := c.ResetQuery().Query("purl", purl).Request("GET", "/v3/purl")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&responseJSON)
	return responseJSON, nil
}

// GetPurls https://docs.vulncheck.com/api/purls
func (c *Client) GetPurls(purls []string) (responseJSON *PurlsResponse, err error) {
	purlBytes, err := json.Marshal(purls)
	if err != nil {
		return nil, err
	}
	resp, err := c.PostRequestWithBody("/v3/purls", bytes.NewReader(purlBytes))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r PurlResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// GetData Returns the data from the response
func (r PurlResponse) GetData() PurlData {
	return r.Data
}

// PurlMeta Returns the PurlMeta from the Metadata
func (r PurlResponse) PurlMeta() PurlMeta {
	return r.Meta.PurlMeta
}

// Cves Cves Returns the list of CVEs associated with the purl
func (r PurlResponse) Cves() []string {
	return r.Data.Cves
}

// Vulnerabilities Returns the list of vulnerabilities associated with the purl
func (r PurlResponse) Vulnerabilities() []PurlVulnerability {
	return r.Data.Vulnerabilities
}
