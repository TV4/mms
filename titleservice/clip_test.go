package titleservice

import (
	"context"
	"net/http"
	"testing"
)

func TestRegisterClip(t *testing.T) {
	var (
		statusCode = http.StatusInternalServerError
		hf         = testHandlerFunc(statusCode, nil)
	)

	ts, c := testServerAndClient(testUser, testPass, hf)
	defer ts.Close()

	r, err := c.RegisterClip(context.Background(), Clip{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := r.StatusCode, statusCode; got != want {
		t.Fatalf("r.StatusCode = %d, want %d", got, want)
	}
}
