package titleservice

import (
	"context"
	"net/url"
)

// Series is a title linked to a series/season
type Series struct {
	SeriesCode   string
	Title        string
	SeasonNumber int
	Description  string
	GenreText    string
}

// Endpoint returns the endpoint to use for this request type
func (s Series) Endpoint() Endpoint {
	return RegisterSeriesEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (s Series) Params() (url.Values, error) {
	return url.Values{}, nil
}

func (c *client) RegisterSeries(ctx context.Context, series Series) (*Response, error) {
	return c.register(ctx, series)
}
