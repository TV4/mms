package titleservice

import (
	"context"
	"fmt"
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

func TestClipValidate(t *testing.T) {
	for _, tt := range []struct {
		c    *Clip
		want string
	}{
		{&Clip{}, "TitleCode: missing parameter"},
		{&Clip{TitleCode: "TC"}, "Title: missing parameter"},
		{&Clip{TitleCode: "TC", Title: "T"}, "Length: missing parameter"},
		{&Clip{TitleCode: "TC", Title: "T", Length: 1}, "PublishedAt: invalid parameter"},
		{&Clip{TitleCode: "TC", Title: "T", Length: 1, PublishedAt: "invalid"}, "PublishedAt: invalid parameter"},
		{&Clip{TitleCode: "TC", Title: "T", Length: 1, PublishedAt: "20070102"}, "<nil>"},
		{&Clip{TitleCode: "TC", Title: "T", Length: 1, PublishedAt: "20070102", Description: "<b>"}, "Description: invalid parameter"},
	} {
		if got := fmt.Sprintf("%v", tt.c.Validate()); got != tt.want {
			t.Fatalf("tt.c.Validate() = %q, want %q", got, tt.want)
		}
	}
}
