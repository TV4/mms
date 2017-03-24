package titleservice

import (
	"context"
	"net/url"
)

// Clip is a “stand-alone” short clip not linked to a series/season or other title
type Clip struct {
	TitleCode      string
	Title          string
	Length         int
	PublishedAt    string
	AvailableUntil string
	Description    string
	PlayURL        string
}

// Endpoint returns the endpoint to use for this request type
func (c Clip) Endpoint() Endpoint {
	return RegisterClipEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (c Clip) Params() (url.Values, error) {
	return url.Values{}, nil
}

func (c *client) RegisterClip(ctx context.Context, clip Clip) (*Response, error) {
	return c.register(ctx, clip)
}
