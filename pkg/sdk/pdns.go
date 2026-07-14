package sdk

import (
	"io"
	"net/url"
)

// GetPdns https://docs.vulncheck.com/api/pdns
func (c *Client) GetPdns(list string) (string, error) {
	resp, err := c.ResetQuery().Request("GET", "/v3/pdns/"+url.QueryEscape(list))
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
