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

// Client for the MMS TitleService API
type Client interface {
	RegisterSeries(context.Context, Series) (*Response, error)
	RegisterEpisode(context.Context, Episode) (*Response, error)
	RegisterClip(context.Context, Clip) (*Response, error)
}

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

type client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
	username   string
	password   string
	simulate   bool
}

// NewClient creates a MMS TitleService Client
func NewClient(username, password string, options ...func(*client)) Client {
	c := &client{
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
func HTTPClient(hc *http.Client) func(*client) {
	return func(c *client) {
		c.httpClient = hc
	}
}

// BaseURL changes the *client base URL based on the provided rawurl
func BaseURL(rawurl string) func(*client) {
	return func(c *client) {
		if u, err := url.Parse(rawurl); err == nil {
			c.baseURL = u
		}
	}
}

// UserAgent changes the User-Agent used in requests sent by the *client
func UserAgent(ua string) func(*client) {
	return func(c *client) {
		c.userAgent = ua
	}
}

// Simulate configures the client to make simulated requests (nothing will be saved to MMS' databases)
func Simulate(c *client) {
	c.simulate = true
}

func (c *client) register(ctx context.Context, req Request) (*Response, error) {
	params, err := req.Params()
	if err != nil {
		return nil, ErrorWithMessage(err, string(req.Endpoint()))
	}

	return c.post(ctx, req.Endpoint(), params)
}

func (c *client) post(ctx context.Context, endpoint Endpoint, params url.Values) (*Response, error) {
	req, err := c.request(ctx, string(endpoint), params)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *client) request(ctx context.Context, path string, params url.Values) (*http.Request, error) {
	params.Set("user", c.username)
	params.Set("pass", c.password)

	if c.simulate {
		params.Set("simulate", "")
	}

	rel, err := url.Parse(path)
	if err != nil {
		return nil, ErrorWithMessage(err, "unable to parse path")
	}

	rawurl := c.baseURL.ResolveReference(rel).String()

	req, err := http.NewRequest("POST", rawurl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, ErrorWithMessage(err, "unable to create POST request")
	}

	req = req.WithContext(ctx)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *client) do(req *http.Request) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, ErrorWithMessage(err, "error sending the request")
	}
	defer func() {
		_, _ = io.CopyN(ioutil.Discard, resp.Body, 64)
		_ = resp.Body.Close()
	}()

	var r Response

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, ErrorWithMessage(err, "unable to decode the response body as JSON")
	}

	return &r, nil
}
