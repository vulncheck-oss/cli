package sdk

import (
	"io"
	"net/http"
	"net/url"
)

// https://docs.vulncheck.com/api/rules
func (c *Client) GetRule(rule string) (string, error) {

	// in the future, we can add more indexes. For now, we only have initial-access
	index := "initial-access"

	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/rules/"+index+"/"+url.QueryEscape(rule), nil)
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
