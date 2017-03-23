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

func (c *client) RegisterClip(ctx context.Context, clip Clip) (*Response, error) {
	return c.post(ctx, "RegisterClip", url.Values{})
}
