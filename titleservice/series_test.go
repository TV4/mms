package titleservice

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestRegisterSeries(t *testing.T) {
	var (
		statusCode = http.StatusTeapot
		hf         = testHandlerFunc(statusCode, nil)
	)

	ts, c := testServerAndClient(testUser, testPass, hf)
	defer ts.Close()

	r, err := c.RegisterSeries(context.Background(), MakeSeries(
		"series-code",
		"series-title",
	))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := r.StatusCode, statusCode; got != want {
		t.Fatalf("r.StatusCode = %d, want %d", got, want)
	}
}

func TestSeriesValidate(t *testing.T) {
	for _, tt := range []struct {
		s    *Series
		want string
	}{
		{&Series{}, "Series SeriesCode: missing parameter"},
		{&Series{SeriesCode: "S"}, "Series Title: missing parameter"},
		{&Series{SeriesCode: "S", Title: "T"}, "<nil>"},
		{&Series{SeriesCode: "S", Title: "T", Description: "Foo"}, "<nil>"},
		{&Series{SeriesCode: "S", Title: "T", Description: "<b>Foo</b>"}, "Series Description: invalid parameter"},
	} {
		if got := fmt.Sprintf("%v", tt.s.Validate()); got != tt.want {
			t.Fatalf("tt.s.Validate() = %q, want %q", got, tt.want)
		}
	}
}
