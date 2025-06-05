package sdk

import (
	"encoding/json"
	"fmt"
)

type BackupsMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Href        string `json:"href"`
}

type BackupsResponse struct {
	Benchmark float64       `json:"_benchmark"`
	Data      []BackupsMeta `json:"data"`
}

// GetBackups https://docs.vulncheck.com/api/backups
func (c *Client) GetBackups() (responseJSON *BackupsResponse, err error) {

	resp, err := c.Request("GET", "/v3/backup")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r BackupsResponse) String() string {
	return fmt.Sprintf("Benchmark: %v\nData: %v\n", r.Benchmark, r.Data)
}

// GetData - Returns the data from the response
func (r BackupsResponse) GetData() []BackupsMeta {
	return r.Data
}
