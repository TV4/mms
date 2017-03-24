package titleservice

import (
	"context"
	"net/url"
)

// Episode is a title linked to a series/season
type Episode struct {
	TitleCode       string
	SeriesCode      string
	Title           string
	Length          string
	PublishedAt     string
	AvailableUntil  string
	CategoryID      CategoryID
	EpisodeNumber   string
	Description     string
	LinkedTitleCode string
	LiveTitle       string
	LiveTvDay       string
	LiveTime        string
	LiveChannelID   LiveChannelID
	PlayURL         string
	TargetGroupCode string
	TerritoryCode   string
	SuggestedGenre1 string
	SuggestedGenre2 string
	SuggestedGenre3 string
}

// Endpoint returns the endpoint to use for this request type
func (e Episode) Endpoint() Endpoint {
	return RegisterEpisodeEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (e Episode) Params() (url.Values, error) {
	return url.Values{}, nil
}

func (c *client) RegisterEpisode(ctx context.Context, episode Episode) (*Response, error) {
	return c.register(ctx, episode)
}
