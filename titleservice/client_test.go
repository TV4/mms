package titleservice

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Run("Defaults", func(t *testing.T) {
		username, password := "foo", "bar"

		c := NewClient(username, password).(*client)

		if got, want := c.httpClient.Timeout, 30*time.Second; got != want {
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
}
