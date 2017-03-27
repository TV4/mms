package titleservice

import (
	"context"
	"net/http"
	"testing"
)

func TestRegisterEpisode(t *testing.T) {
	var (
		statusCode = http.StatusInternalServerError
		hf         = testHandlerFunc(statusCode, nil)
	)

	ts, c := testServerAndClient(testUser, testPass, hf)
	defer ts.Close()

	r, err := c.RegisterEpisode(context.Background(), MakeEpisode(
		"episode-title-code",
		"episode-series-code",
		"episode-title",

		123,
		Date(2017, 3, 27),
		TvProgram,

		func(e *Episode) {
			e.LiveTitle = "episode-live-title"
			e.LiveTvDay = Date(2017, 1, 2)
			e.LiveTime = Time(10, 20)
			e.LiveChannelID = TV4
		},
	))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := r.StatusCode, statusCode; got != want {
		t.Fatalf("r.StatusCode = %d, want %d", got, want)
	}
}
