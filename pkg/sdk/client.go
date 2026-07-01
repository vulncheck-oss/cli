package sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// defaultHTTPTimeout guards every SDK request against a stuck / slow-read
// peer. Long enough for legitimate large-index metadata calls, short
// enough that a stalled connection can't hang the CLI indefinitely.
const defaultHTTPTimeout = 5 * time.Minute

// maxResponseBytes caps the size of any API response body the SDK will
// buffer into memory. Prevents a malicious or misbehaving upstream from
// OOM-killing the CLI by streaming a multi-GB response. Genuinely large
// payloads (backup archives) are downloaded via streamed io.Copy, not
// through this path — see pkg/cache/tasks.go and pkg/ui/download.go.
const maxResponseBytes = 512 * 1024 * 1024 // 512 MiB

// LimitedBody wraps resp.Body with an io.LimitReader honouring the
// maxResponseBytes ceiling. SDK methods use this before json-decoding.
func LimitedBody(body io.Reader) io.Reader {
	return io.LimitReader(body, maxResponseBytes)
}

// newHTTPClient builds an http.Client with the SDK's hardening defaults:
// explicit timeout so a slow peer cannot hang forever, and Go's default
// TLS verification (never disabled anywhere in the codebase).
func newHTTPClient() *http.Client {
	return &http.Client{Timeout: defaultHTTPTimeout}
}

// ResetQuery clears any accumulated query / form params from prior calls.
// SDK methods that build up state via c.Query / c.Form must call this
// first so a Client reused across multiple requests (e.g. inside a batch
// loop) doesn't leak params from earlier iterations into later ones.
func (c *Client) ResetQuery() *Client {
	c.Values = nil
	c.FormValues = nil
	return c
}

type Client struct {
	Url         string
	Token       string
	HttpClient  *http.Client
	HttpRequest *http.Request
	UserAgent   string
	Values      *url.Values
	FormValues  *url.Values
	ctx         context.Context
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

// WithContext attaches a context to subsequent HTTP requests issued by the
// client. Callers wire the command's cancellable context here so that
// SIGINT/SIGTERM and explicit cancellation propagate to in-flight calls.
func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

// context returns the attached context or context.Background when none is set.
func (c *Client) context() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}

// SetAuthHeader Sets the Authorization header for the request
func (c *Client) SetAuthHeader(req *http.Request) *Client {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	return c
}

func (c *Client) Request(method string, url string) (*http.Response, error) {
	if c.HttpClient == nil {
		c.HttpClient = newHTTPClient()
	}
	var err error
	if c.FormValues != nil {
		c.HttpRequest, err = http.NewRequestWithContext(c.context(), method, c.GetUrl()+url, strings.NewReader(c.FormValues.Encode()))
	} else {
		c.HttpRequest, err = http.NewRequestWithContext(c.context(), method, c.GetUrl()+url, nil)
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

func (c *Client) PostRequestWithBody(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(c.context(), http.MethodPost, c.GetUrl()+url, body)
	if err != nil {
		return nil, fmt.Errorf("making new post request: %w", err)
	}

	if c.HttpClient == nil {
		c.HttpClient = newHTTPClient()
	}
	c.SetAuthHeader(req)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.HttpClient.Do(req)
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
