package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Url         string
	Token       string
	HttpClient  *http.Client
	HttpRequest *http.Request
	UserAgent   string
	Values      *url.Values
	FormValues  *url.Values
}

type MetaError struct {
	Error  bool     `json:"error"`
	Errors []string `json:"errors"`
}

type ReqError struct {
	StatusCode int
	Reason     MetaError
}

var ErrorUnauthorized = fmt.Errorf("unauthorized")

func Connect(url string, token string) *Client {
	return &Client{Url: url, Token: token}
}

func (c *Client) GetToken() string {
	return c.Token
}

func (c *Client) SetToken(token string) *Client {
	c.Token = token
	return c
}

func (c *Client) SetUrl(env string) *Client {
	c.Url = env
	return c
}

func (c *Client) GetUrl() string {
	return c.Url
}

func (c *Client) SetUserAgent(userAgent string) *Client {
	c.UserAgent = userAgent
	return c
}

// SetAuthHeader Sets the Authorization header for the request
func (c *Client) SetAuthHeader(req *http.Request) *Client {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	return c
}

func (c *Client) Request(method string, url string) (*http.Response, error) {
	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}
	}
	var err error
	if c.FormValues != nil {
		c.HttpRequest, err = http.NewRequest(method, c.GetUrl()+url, strings.NewReader(c.FormValues.Encode()))
	} else {
		c.HttpRequest, err = http.NewRequest(method, c.GetUrl()+url, nil)
	}
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(c.HttpRequest)

	if c.UserAgent != "" {
		c.HttpRequest.Header.Set("User-Agent", c.UserAgent)
	}

	if c.Values != nil {
		c.HttpRequest.URL.RawQuery = c.Values.Encode()
	}

	if c.FormValues != nil {
		c.HttpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := c.HttpClient.Do(c.HttpRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 401 {
		return nil, ErrorUnauthorized
	}

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	return resp, nil
}

func (c *Client) Query(key string, value string) *Client {
	if c.Values == nil {
		c.Values = &url.Values{}
	}
	c.Values.Add(key, value)
	return c
}

func (c *Client) Form(key string, value string) *Client {
	if c.FormValues == nil {
		c.FormValues = &url.Values{}
	}
	c.FormValues.Add(key, value)
	return c
}

func (e ReqError) Error() string {
	return fmt.Sprintf("error: %t, status code: %d, errors: %v", e.Reason.Error, e.StatusCode, e.Reason.Errors)
}
