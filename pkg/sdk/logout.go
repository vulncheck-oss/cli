package sdk

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Benchmark float64  `json:"_benchmark"`
	Data      UserData `json:"data"`
	Type      string   `json:"type"`
	Error     bool     `json:"error"`
	Errors    []string `json:"errors"`
	Success   bool     `json:"success"`
}

// https://docs.vulncheck.com/api/logout
func (c *Client) Logout() (responseJSON *Response, err error) {
	resp, err := c.Request("GET", "/logout")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

// Strings representation of the response
func (r Response) String() string {
	return fmt.Sprintf("Benchmark: %v\nResponse: %v\n", r.Benchmark, r.Success)
}
