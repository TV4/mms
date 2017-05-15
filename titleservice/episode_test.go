package titleservice

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestRegisterEpisode(t *testing.T) {
	var (
		statusCode = http.StatusTeapot
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

func TestEpisodeValidate(t *testing.T) {
	for _, tt := range []struct {
		e    *Episode
		want string
	}{
		{&Episode{}, "Episode TitleCode: missing parameter"},
		{&Episode{TitleCode: "TC"}, "Episode SeriesCode: missing parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC"}, "Episode Title: missing parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T"}, "Episode Length: invalid parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T", Length: 1}, "Episode PublishedAt: invalid parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T", Length: 1, PublishedAt: "20070102"}, "Episode CategoryID: invalid parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T", Length: 1, PublishedAt: "20070102", CategoryID: Webisode}, "<nil>"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T", Length: 1, PublishedAt: "20070102", CategoryID: Webisode, Description: "<b>"}, "Episode Description: invalid parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T", Length: 1, PublishedAt: "20070102", CategoryID: TvProgram}, "Episode LiveTitle: missing parameter"},
		{&Episode{TitleCode: "TC", SeriesCode: "SC", Title: "T", Length: 1, PublishedAt: "20070102", CategoryID: TvProgram, LiveTitle: "LT"}, "Episode LiveTvDay: invalid parameter"},
	} {
		if got := fmt.Sprintf("%v", tt.e.Validate()); got != tt.want {
			t.Fatalf("tt.e.Validate() = %q, want %q", got, tt.want)
		}
	}
}
