package sdk

import (
	"encoding/json"
	"fmt"
)

type BackupFile struct {
	Filename        string `json:"filename"`
	Sha256          string `json:"sha256"`
	DateAdded       string `json:"date_added"`
	URLTtlMinutes   int    `json:"url_ttl_minutes"`
	URLExpires      string `json:"url_expires"`
	URL             string `json:"url"`
	URLMrap         string `json:"url_mrap"`
	URLUsEast1      string `json:"url_us-east-1"`
	URLUsWest2      string `json:"url_us-west-2"`
	URLEuWest2      string `json:"url_eu-west-2"`
	URLApSoutheast2 string `json:"url_ap-southeast-2"`
}

type BackupResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		Timestamp string `json:"timestamp"`
		Index     string `json:"index"`
	} `json:"_meta"`
	Data []BackupFile `json:"data"`
}

// https://docs.vulncheck.com/api/backup
func (c *Client) GetIndexBackup(index string) (responseJSON *BackupResponse, err error) {
	resp, err := c.Request("GET", "/v3/backup/"+index)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r BackupResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// Returns the data from the response
func (r BackupResponse) GetData() []BackupFile {
	return r.Data
}

// Returns the list of filenames associated with the backup
func (r BackupResponse) Filenames() []string {
	var filenames []string
	for _, v := range r.Data {
		filenames = append(filenames, v.Filename)
	}
	return filenames
}

// Returns the list of URLs associated with the backup
func (r BackupResponse) Urls() []string {
	var urls []string
	for _, v := range r.Data {
		urls = append(urls, v.URL)
	}
	return urls
}
