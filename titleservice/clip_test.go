package titleservice

import (
	"context"
	"net/http"
	"testing"
)

func TestRegisterClip(t *testing.T) {
	var (
		statusCode = http.StatusTeapot
		hf         = testHandlerFunc(statusCode, nil)
	)

	ts, c := testServerAndClient(testUser, testPass, hf)
	defer ts.Close()

	r, err := c.RegisterClip(context.Background(), MakeClip(
		"clip-title-code",
		"clip-title",
		123,
		Date(2017, 3, 27),
	))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := r.StatusCode, statusCode; got != want {
		t.Fatalf("r.StatusCode = %d, want %d", got, want)
	}
}
