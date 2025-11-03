package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vulncheck-oss/cli/pkg/utils"
)

type IndexQueryParameters struct {
	// INDEX RELATED
	Cve                string `json:"cve"`
	Alias              string `json:"alias"`
	Iava               string `json:"iava"`
	UpdatedAtStartDate string `json:"updatedAtStartDate"`
	UpdatedAtEndDate   string `json:"updatedAtEndDate"`
	LastModStartDate   string `json:"lastModStartDate"`
	LastModEndDate     string `json:"lastModEndDate"`
	PubStartDate       string `json:"pubStartDate"`
	PubEndDate         string `json:"pubEndDate"`
	ThreatActor        string `json:"threat_actor"`
	MitreId            string `json:"mitre_id"`
	MispId             string `json:"misp_id"`
	Ransomware         string `json:"ransomware"`
	Botnet             string `json:"botnet"`
	Hostname           string `json:"hostname"`
	ID                 string `json:"id"`
	Kind               string `json:"kind"`
	Country            string `json:"country"`
	CountryCode        string `json:"country_code"`
	Asn                string `json:"asn"`
	Cidr               string `json:"cidr"`
	Ilvn               string `json:"ilvn"`
	Jvndb              string `json:"jvndb"`
	SrcCountry         string `json:"src_country"`
	DstCountry         string `json:"dst_country"`

	// PAGINATION RELATED
	Limit       int    `json:"limit"`
	Sort        string `json:"sort"`
	Order       string `json:"order"`
	Page        int    `json:"page"`
	StartCursor bool   `json:"start_cursor"`
	Cursor      string `json:"cursor"`
}

type IndexMeta struct {
	Timestamp      string                `json:"timestamp"`
	Index          string                `json:"index"`
	Limit          int                   `json:"limit"`
	TotalDocuments int                   `json:"total_documents"`
	Sort           string                `json:"sort"`
	Parameters     []IndexMetaParameters `json:"parameters"`
	Order          string                `json:"order"`
	Page           int                   `json:"page"`
	TotalPages     int                   `json:"total_pages"`
	MaxPages       int                   `json:"max_pages"`
	FirstItem      int                   `json:"first_item"`
	LastItem       int                   `json:"last_item"`
	NextCursor     string                `json:"next_cursor"`
}

type IndexMetaParameters struct {
	Name   string `json:"name"`
	Format string `json:"format"`
}

type IndexResponse struct {
	Benchmark float64           `json:"_benchmark"`
	Meta      IndexMeta         `json:"_meta"`
	Data      []json.RawMessage `json:"data"`
}

// add method to set query parameters
func setIndexQueryParameters(query url.Values, queryParameters ...IndexQueryParameters) {
	for _, queryParameter := range queryParameters {
		// INDEX RELATED
		if queryParameter.Cve != "" {
			query.Add("cve", utils.FormatCVE(queryParameter.Cve))
		}
		if queryParameter.Alias != "" {
			query.Add("alias", queryParameter.Alias)
		}
		if queryParameter.Iava != "" {
			query.Add("iava", queryParameter.Iava)
		}
		if queryParameter.UpdatedAtStartDate != "" {
			query.Add("updatedAtStartDate", queryParameter.UpdatedAtStartDate)
		}
		if queryParameter.UpdatedAtEndDate != "" {
			query.Add("updatedAtEndDate", queryParameter.UpdatedAtEndDate)
		}
		if queryParameter.LastModStartDate != "" {
			query.Add("lastModStartDate", queryParameter.LastModStartDate)
		}
		if queryParameter.LastModEndDate != "" {
			query.Add("lastModEndDate", queryParameter.LastModEndDate)
		}
		if queryParameter.PubStartDate != "" {
			query.Add("pubStartDate", queryParameter.PubStartDate)
		}
		if queryParameter.PubEndDate != "" {
			query.Add("pubEndDate", queryParameter.PubEndDate)
		}
		if queryParameter.ThreatActor != "" {
			query.Add("threat_actor", queryParameter.ThreatActor)
		}
		if queryParameter.MitreId != "" {
			query.Add("mitre_id", queryParameter.MitreId)
		}
		if queryParameter.MispId != "" {
			query.Add("misp_id", queryParameter.MispId)
		}
		if queryParameter.Ransomware != "" {
			query.Add("ransomware", queryParameter.Ransomware)
		}
		if queryParameter.Botnet != "" {
			query.Add("botnet", queryParameter.Botnet)
		}
		if queryParameter.Hostname != "" {
			query.Add("hostname", queryParameter.Hostname)
		}
		if queryParameter.ID != "" {
			query.Add("id", queryParameter.ID)
		}
		if queryParameter.Kind != "" {
			query.Add("kind", queryParameter.Kind)
		}
		if queryParameter.Country != "" {
			query.Add("country", queryParameter.Country)
		}
		if queryParameter.CountryCode != "" {
			query.Add("country_code", queryParameter.CountryCode)
		}
		if queryParameter.Asn != "" {
			query.Add("asn", queryParameter.Asn)
		}
		if queryParameter.Cidr != "" {
			query.Add("cidr", queryParameter.Cidr)
		}
		if queryParameter.Ilvn != "" {
			query.Add("ilvn", queryParameter.Ilvn)
		}
		if queryParameter.Jvndb != "" {
			query.Add("jvndb", queryParameter.Jvndb)
		}
		if queryParameter.SrcCountry != "" {
			query.Add("src_country", queryParameter.SrcCountry)
		}
		if queryParameter.DstCountry != "" {
			query.Add("dst_country", queryParameter.DstCountry)
		}
		// PAGINATION RELATED
		if queryParameter.Limit != 0 {
			query.Add("limit", fmt.Sprintf("%d", queryParameter.Limit))
		}
		if queryParameter.Sort != "" {
			query.Add("sort", queryParameter.Sort)
		}
		if queryParameter.Order != "" {
			query.Add("order", queryParameter.Order)
		}
		if queryParameter.Page != 0 {
			query.Add("page", fmt.Sprintf("%d", queryParameter.Page))
		}
		if queryParameter.StartCursor {
			query.Add("start_cursor", fmt.Sprintf("%t", queryParameter.StartCursor))
		}
		if queryParameter.Cursor != "" {
			query.Add("cursor", queryParameter.Cursor)
			query.Del("start_cursor") // cursor and start_cursor are mutually exclusive
		}
	}
}

// https://docs.vulncheck.com/api/indice
func (c *Client) GetIndex(index string, queryParameters ...IndexQueryParameters) (responseJSON *IndexResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape(index), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

// Strings representation of the response
func (r IndexResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// GetData - Returns the data from the response
func (r IndexResponse) GetData() []json.RawMessage {
	return r.Data
}
