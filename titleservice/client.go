package titleservice

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultScheme    = "https"
	defaultHost      = "titleservice.mms.se"
	defaultUserAgent = "titleservice/client.go (godoc.org/github.com/TV4/mms/titleservice)"
	defaultTimeout   = 30 * time.Second
)

// Request interface used in requests to the MMS TitleService API
type Request interface {
	Endpoint() Endpoint
	Params() (url.Values, error)
}

// Response is used for responses from the MMS TitleService API
type Response struct {
	StatusCode        int      `json:"StatusCode"`
	StatusDescription string   `json:"StatusDescription"`
	Errors            []string `json:"Errors"`
}

// Client for the MMS TitleService API
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
	username   string
	password   string
	simulate   bool
}

// NewClient creates a MMS TitleService Client
func NewClient(username, password string, options ...func(*Client)) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: &url.URL{
			Scheme: defaultScheme,
			Host:   defaultHost,
		},
		userAgent: defaultUserAgent,
		username:  username,
		password:  password,
	}

	for _, f := range options {
		f(c)
	}

	return c
}

// HTTPClient changes the *client HTTP client to the provided *http.Client
func HTTPClient(hc *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// BaseURL changes the *client base URL based on the provided rawurl
func BaseURL(rawurl string) func(*Client) {
	return func(c *Client) {
		if u, err := url.Parse(rawurl); err == nil {
			c.baseURL = u
		}
	}
}

// UserAgent changes the User-Agent used in requests sent by the *client
func UserAgent(ua string) func(*Client) {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// Simulate configures the client to make simulated requests (nothing will be saved to MMS' databases)
func Simulate(b bool) func(c *Client) {
	return func(c *Client) {
		c.simulate = b
	}
}

// Simulated returns true if the client is configured to send simulated requests
func (c *Client) Simulated() bool {
	return c.simulate
}

// RegisterSeries registers a Series
func (c *Client) RegisterSeries(ctx context.Context, series Series) (*Response, error) {
	return c.register(ctx, &series)
}

// RegisterEpisode registers an Episode
func (c *Client) RegisterEpisode(ctx context.Context, episode Episode) (*Response, error) {
	return c.register(ctx, &episode)
}

// RegisterClip registers a Clip
func (c *Client) RegisterClip(ctx context.Context, clip Clip) (*Response, error) {
	return c.register(ctx, &clip)
}

func (c *Client) register(ctx context.Context, req Request) (*Response, error) {
	params, err := req.Params()
	if err != nil {
		return nil, newErrorWithMessage(err, string(req.Endpoint()))
	}

	return c.post(ctx, req.Endpoint(), params)
}

func (c *Client) post(ctx context.Context, endpoint Endpoint, params url.Values) (*Response, error) {
	req, err := c.request(ctx, string(endpoint), params)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) request(ctx context.Context, path string, params url.Values) (*http.Request, error) {
	if err := c.validateCredentials(); err != nil {
		return nil, err
	}

	params.Set("user", c.username)
	params.Set("pass", c.password)

	if c.simulate {
		params.Set("simulate", "")
	}

	rel, err := url.Parse(path)
	if err != nil {
		return nil, newErrorWithMessage(err, "unable to parse path")
	}

	rawurl := c.baseURL.ResolveReference(rel).String()

	req, err := http.NewRequest("POST", rawurl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, newErrorWithMessage(err, "unable to create POST request")
	}

	req = req.WithContext(ctx)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) do(req *http.Request) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, newErrorWithMessage(err, "error sending the request")
	}
	defer func() {
		_, _ = io.CopyN(ioutil.Discard, resp.Body, 64)
		_ = resp.Body.Close()
	}()

	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
		err := newErrorWithMessage(ErrUnexpectedContentType, ct)

		return &Response{
			StatusCode:        resp.StatusCode,
			StatusDescription: resp.Status,
			Errors:            []string{err.Error()},
		}, err
	}

	var r Response

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, newErrorWithMessage(err, "unable to decode the response body as JSON")
	}

	return &r, nil
}

func (c *Client) validateCredentials() error {
	if c.username == "" {
		return ErrNoUsername
	}

	if c.password == "" {
		return ErrNoPassword
	}

	return nil
}
