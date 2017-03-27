package titleservice

import (
	"fmt"
	"net/url"
	"strings"
)

// MakeClip creates a Clip based on required parameters and optional parameters using options
func MakeClip(titleCode, title string, length int, publishedAt string, options ...func(*Clip)) Clip {
	c := &Clip{
		TitleCode:   titleCode,
		Title:       title,
		Length:      length,
		PublishedAt: publishedAt,
	}

	for _, f := range options {
		f(c)
	}

	return *c
}

// Clip is a “stand-alone” short clip not linked to a series/season or other title
type Clip struct {
	TitleCode      string // required
	Title          string // required
	Length         int    // required
	PublishedAt    string // required
	AvailableUntil string
	Description    string
	PlayURL        string
}

// Endpoint returns the endpoint to use for this request type
func (c *Clip) Endpoint() Endpoint {
	return RegisterClipEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (c *Clip) Params() (url.Values, error) {
	if err := c.validate(); err != nil {
		return url.Values{}, err
	}

	params := url.Values{
		"TitleCode":   {c.TitleCode},
		"Title":       {c.Title},
		"Length":      {fmt.Sprintf("%d", c.Length)},
		"PublishedAt": {c.PublishedAt},
	}

	// optional parameters

	if len(c.AvailableUntil) == 8 {
		params.Set("AvailableUntil", c.AvailableUntil)
	}

	if c.Description != "" {
		params.Set("Description", c.Description)
	}

	if strings.HasPrefix(c.PlayURL, "http") {
		params.Set("PlayUrl", c.PlayURL)
	}

	return params, nil
}

func (c *Clip) validate() error {
	if c.TitleCode == "" {
		return newErrorWithMessage(ErrMissingParameter, "TitleCode")
	}

	if c.Title == "" {
		return newErrorWithMessage(ErrMissingParameter, "Title")
	}

	if c.Length < 1 {
		return newErrorWithMessage(ErrMissingParameter, "Length")
	}

	if len(c.PublishedAt) != 8 {
		return newErrorWithMessage(ErrInvalidParameter, "PublishedAt")
	}

	return nil
}
