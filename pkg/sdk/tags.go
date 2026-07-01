package sdk

import (
	"io"
	"net/url"
)

// GetTag https://docs.vulncheck.com/api/tags
//
// Routes through c.Request so ctx propagation, timeouts, and auth are
// consistent with every other SDK method. Response body is size-limited
// to defeat a malicious/misbehaving upstream returning multi-GB payloads.
func (c *Client) GetTag(tag string) (string, error) {
	resp, err := c.ResetQuery().Request("GET", "/v3/tags/"+url.QueryEscape(tag))
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
