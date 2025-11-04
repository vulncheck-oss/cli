package sdk

import (
	"io"
	"net/http"
	"net/url"
)

// https://docs.vulncheck.com/api/pdns
func (c *Client) GetPdns(list string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/pdns/"+url.QueryEscape(list), nil)
	if err != nil {
		return "", err
	}

	c.SetAuthHeader(req)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func() { _ = res.Body.Close() }()
	body, _ := io.ReadAll(res.Body)

	return string(body), nil
}
