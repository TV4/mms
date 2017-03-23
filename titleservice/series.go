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

func (c *client) RegisterSeries(ctx context.Context, series Series) (*Response, error) {
	return c.post(ctx, "RegisterSeries", url.Values{})
}
