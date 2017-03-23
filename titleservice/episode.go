package titleservice

import "context"

type Episode struct{}

func (c *client) RegisterEpisode(ctx context.Context, episode Episode) {}
