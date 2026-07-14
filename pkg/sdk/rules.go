package sdk

import (
	"io"
	"net/url"
)

// GetRule https://docs.vulncheck.com/api/rules
func (c *Client) GetRule(rule string) (string, error) {
	// Only the initial-access rule index is exposed today; the path is
	// hard-coded here rather than plumbed through the API.
	const index = "initial-access"

	resp, err := c.ResetQuery().Request("GET", "/v3/rules/"+index+"/"+url.QueryEscape(rule))
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(LimitedBody(resp.Body))
	if err != nil {
		return "", err
	}
	return string(body), nil
}
