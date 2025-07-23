package sdk

import (
	"io"
	"net/http"
	"net/url"
)

// https://docs.vulncheck.com/api/tags
func (c *Client) GetTag(tag string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/tags/"+url.QueryEscape(tag), nil)
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
