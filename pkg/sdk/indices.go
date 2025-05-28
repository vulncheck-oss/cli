package sdk

import (
	"encoding/json"
	"fmt"
)

type IndicesMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Href        string `json:"href"`
}

type IndicesResponse struct {
	Benchmark float64       `json:"_benchmark"`
	Data      []IndicesMeta `json:"data"`
}

// GetIndices https://docs.vulncheck.com/api/indexes
func (c *Client) GetIndices() (responseJSON *IndicesResponse, err error) {

	resp, err := c.Request("GET", "/v3/index")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r IndicesResponse) String() string {
	return fmt.Sprintf("Benchmark: %v\nData: %v\n", r.Benchmark, r.Data)
}

// GetData - Returns the data from the response
func (r IndicesResponse) GetData() []IndicesMeta {
	return r.Data
}
