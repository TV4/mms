package titleservice

import "context"

type Series struct {
	SeriesCode   string
	Title        string
	SeasonNumber int
	Description  string
	GenreText    string
}

func (c *client) RegisterSeries(ctx context.Context, series Series) {}
