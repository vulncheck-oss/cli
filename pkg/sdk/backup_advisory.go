package sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type AdvisoryBackupMeta struct {
	Name            string `json:"name"`
	Href            string `json:"href"`
	Available       bool   `json:"available"`
	BackupWrittenAt string `json:"backup_written_at"`
}

type AdvisoryBackupsResponse struct {
	Data []AdvisoryBackupMeta `json:"data"`
}

// AdvisoryBackup is the /v4/backup/{feed} payload.
//
// Unlike /v3/backup/{index} this is a flat object rather than a data array, it
// carries no filename or date_added (backup_written_at lives on the list
// endpoint instead), there is no plain "url" field, and url_ttl_minutes is a
// number here where v3 sends a string.
type AdvisoryBackup struct {
	Feed            string `json:"feed"`
	Available       bool   `json:"available"`
	URLMrap         string `json:"url_mrap"`
	URLUsEast1      string `json:"url_us-east-1"`
	URLEuWest2      string `json:"url_eu-west-2"`
	URLApSoutheast2 string `json:"url_ap-southeast-2"`
	URLExpires      string `json:"url_expires"`
	URLTTLMinutes   int    `json:"url_ttl_minutes"`
	SHA256          string `json:"sha256"`

	// URL is not sent by the API, which offers only the regional variants
	// above. It is resolved on decode so that `jq -r .url` works identically
	// against a v3 and a v4 backup payload.
	URL string `json:"url"`
}

// resolveURL picks the download URL, preferring the multi-region access point
// and falling back to a regional bucket if it is absent.
func (b AdvisoryBackup) resolveURL() string {
	for _, u := range []string{b.URLMrap, b.URLUsEast1, b.URLEuWest2, b.URLApSoutheast2} {
		if u != "" {
			return u
		}
	}
	return ""
}

// GetAdvisoryBackups https://docs.vulncheck.com/api/v4/backup
func (c *Client) GetAdvisoryBackups() (responseJSON *AdvisoryBackupsResponse, err error) {
	resp, err := c.ResetQuery().Request("GET", "/v4/backup")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&responseJSON)
	return responseJSON, nil
}

// GetAdvisoryBackup https://docs.vulncheck.com/api/v4/backup
//
// PathEscape, not QueryEscape: the feed is a path segment, and QueryEscape
// encodes a space as "+" rather than "%20".
func (c *Client) GetAdvisoryBackup(feed string) (responseJSON *AdvisoryBackup, err error) {
	resp, err := c.ResetQuery().Request("GET", "/v4/backup/"+url.PathEscape(feed))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&responseJSON)
	if responseJSON != nil {
		responseJSON.URL = responseJSON.resolveURL()
	}
	return responseJSON, nil
}

// Strings representation of the response
func (r AdvisoryBackupsResponse) String() string {
	return fmt.Sprintf("Data: %v\n", r.Data)
}

// GetData - Returns the data from the response
func (r AdvisoryBackupsResponse) GetData() []AdvisoryBackupMeta {
	return r.Data
}
