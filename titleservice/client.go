package titleservice

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	RegisterSeries(context.Context, Series)
	RegisterEpisode(context.Context, Episode)
	RegisterClip(context.Context, Clip)
}

type client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
	username   string
	password   string
}

// NewClient creates an MMS TitleService Client
func NewClient(username, password string, options ...func(*client)) Client {
	c := &client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "titleservice.mms.se",
		},
		userAgent: "mms/titleservice/client.go (https://github.com/TV4/mms)",
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

// Authentication changes the username and password used in requests sent by the *client
func Authentication(username, password string) func(*client) {
	return func(c *client) {
		c.username = username
		c.password = password
	}
}
