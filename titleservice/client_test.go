package titleservice

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Run("Defaults", func(t *testing.T) {
		username, password := "foo", "bar"

		c := NewClient(username, password).(*client)

		if got, want := c.httpClient.Timeout, defaultTimeout; got != want {
			t.Fatalf("c.httpClient.Timeout = %v, want %v", got, want)
		}

		if got, want := c.baseURL.Scheme, defaultScheme; got != want {
			t.Fatalf("c.baseURL.Scheme = %q, want %q", got, want)
		}

		if got, want := c.baseURL.Host, defaultHost; got != want {
			t.Fatalf("c.baseURL.Host = %q, want %q", got, want)
		}

		if got, want := c.userAgent, defaultUserAgent; got != want {
			t.Fatalf("c.userAgent = %q, want %q", got, want)
		}

		if got, want := c.username, "foo"; got != want {
			t.Fatalf("c.username = %q, want %q", got, want)
		}

		if got, want := c.password, "bar"; got != want {
			t.Fatalf("c.password = %q, want %q", got, want)
		}

		if got, want := c.simulate, false; got != want {
			t.Fatalf("c.simulate = %v, want %v", got, want)
		}
	})

	t.Run("HTTPClient", func(t *testing.T) {
		timeout := 123 * time.Second

		c := NewClient("", "", HTTPClient(&http.Client{Timeout: timeout})).(*client)

		if got, want := c.httpClient.Timeout, timeout; got != want {
			t.Fatalf("c.httpClient.Timeout = %v, want %v", got, want)
		}
	})

	t.Run("BaseURL", func(t *testing.T) {
		rawurl := "http://example.com"

		c := NewClient("", "", BaseURL(rawurl)).(*client)

		if got, want := c.baseURL.String(), rawurl; got != want {
			t.Fatalf("c.baseURL.String() = %q, want %q", got, want)
		}
	})

	t.Run("UserAgent", func(t *testing.T) {
		ua := "user-agent-test"

		c := NewClient("", "", UserAgent(ua)).(*client)

		if got, want := c.userAgent, ua; got != want {
			t.Fatalf("c.userAgent = %q, want %q", got, want)
		}
	})

	t.Run("Simulate", func(t *testing.T) {
		c := NewClient("", "", Simulate).(*client)

		if got, want := c.simulate, true; got != want {
			t.Fatalf("c.simulate = %v, want %v", got, want)
		}
	})
}

func TestClientRequest(t *testing.T) {
	for _, tt := range []struct {
		simulate bool
		path     string
		username string
		password string
		v        url.Values
		encoded  string
	}{
		{false, "/FooBar", "fooUser", "barPass", url.Values{}, "pass=barPass&user=fooUser"},
		{false, "/BarBaz", "barUser", "bazPass", url.Values{}, "pass=bazPass&user=barUser"},
		{true, "/BazQux", "bazUser", "quxPass", url.Values{"quux": {"corge"}}, "pass=quxPass&quux=corge&simulate=&user=bazUser"},
	} {
		c := testClient(func(c *client) {
			c.username = tt.username
			c.password = tt.password
			c.simulate = tt.simulate
		})

		req, err := c.request(context.Background(), tt.path, tt.v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.URL.Path, tt.path; got != want {
			t.Fatalf("req.URL.Path = %q, want %q", got, want)
		}

		if got, want := req.Header.Get("Accept"), "application/json"; got != want {
			t.Fatalf(`req.Header.Get("Accept") = %q, want %q`, got, want)
		}

		if got, want := req.Header.Get("User-Agent"), defaultUserAgent; got != want {
			t.Fatalf(`req.Header.Get("Content-Type") = %q, want %q`, got, want)
		}

		if got, want := req.Header.Get("Content-Type"), "application/x-www-form-urlencoded"; got != want {
			t.Fatalf(`req.Header.Get("Content-Type") = %q, want %q`, got, want)
		}

		req.ParseForm()

		if got, want := req.Form.Encode(), tt.encoded; got != want {
			t.Fatalf("req.Form.Encode() = %q, want %q", got, want)
		}
	}
}

const (
	testUser = "testUser-123"
	testPass = "testPass-XYZ"
	testHost = "http://example.com"
)

func testClient(options ...func(*client)) *client {
	c := NewClient(testUser, testPass, BaseURL(testHost)).(*client)

	for _, f := range options {
		f(c)
	}

	return c
}

func testServerAndClient(username, password string, hf http.HandlerFunc) (*httptest.Server, *client) {
	ts := httptest.NewServer(http.HandlerFunc(hf))

	return ts, NewClient(username, password, BaseURL(ts.URL)).(*client)
}

func testHandlerFunc(statusCode int, errors []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		enc := json.NewEncoder(w)

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			enc.Encode(testResponse(http.StatusInternalServerError, nil))
			return
		}

		if r.FormValue("user") != testUser || r.FormValue("pass") != testPass {
			w.WriteHeader(http.StatusForbidden)
			enc.Encode(testResponse(http.StatusForbidden, nil))
			return
		}

		w.WriteHeader(statusCode)

		enc.Encode(testResponse(statusCode, errors))
	}
}

func testResponse(code int, errors []string) *Response {
	if errors == nil {
		errors = []string{}
	}

	return &Response{
		StatusCode:        code,
		StatusDescription: http.StatusText(code),
		Errors:            errors,
	}
}
