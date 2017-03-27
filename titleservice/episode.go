package titleservice

import (
	"fmt"
	"net/url"
	"strings"
)

// MakeEpisode creates an Episode based on required parameters and optional parameters using options
func MakeEpisode(titleCode, seriesCode, title string, length int, publishedAt string, categoryID CategoryID, options ...func(*Episode)) Episode {
	e := &Episode{
		TitleCode:   titleCode,
		SeriesCode:  seriesCode,
		Title:       title,
		Length:      length,
		PublishedAt: publishedAt,
		CategoryID:  categoryID,
	}

	for _, f := range options {
		f(e)
	}

	return *e
}

// Episode is a title linked to a series/season
type Episode struct {
	TitleCode       string        // required
	SeriesCode      string        // required
	Title           string        // required
	Length          int           // required
	PublishedAt     string        // required (YYYYMMDD)
	AvailableUntil  string        // optional (YYYYMMDD)
	CategoryID      CategoryID    // required
	EpisodeNumber   int           // only applicable to Categories 1, 2, 4, 5
	Description     string        // optional
	LinkedTitleCode string        // can only be reported for categories 2, 3, 8 (only for updates)
	LiveTitle       string        // obligatory for categories 1, 2, 3, 8
	LiveTvDay       string        // obligatory for categories 1, 2, 3, 8 (YYYYMMDD)
	LiveTime        string        // obligatory for categories 1, 2, 3, 8 (HHMM, MMS-time: 23:45=2345, 01:45=2545, 02:00=0200)
	LiveChannelID   LiveChannelID // obligatory for categories 1, 2, 3, 8
	PlayURL         string        // maximum of 150 characters
	TargetGroupCode string        // optional V = Vuxen (Adults) B = Barn (Children)
	TerritoryCode   string        // optional S = Svenskt (Swedish) U = Utländskt (Foreign)
	SuggestedGenre1 string        // free text of maximum of 256 characters in length
	SuggestedGenre2 string        // free text of maximum of 256 characters in length
	SuggestedGenre3 string        // free text of maximum of 256 characters in length
}

// Endpoint returns the endpoint to use for this request type
func (e *Episode) Endpoint() Endpoint {
	return RegisterEpisodeEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (e *Episode) Params() (url.Values, error) {
	if err := e.validate(); err != nil {
		return url.Values{}, err
	}

	params := url.Values{
		"TitleCode":   {e.TitleCode},
		"SeriesCode":  {e.SeriesCode},
		"Title":       {e.Title},
		"Length":      {fmt.Sprintf("%d", e.Length)},
		"PublishedAt": {e.PublishedAt},
		"CategoryID":  {fmt.Sprintf("%d", e.CategoryID)},
	}

	if len(e.AvailableUntil) == 8 {
		params.Set("AvailableUntil", e.AvailableUntil)
	}

	// only applicable to Categories 1, 2, 4, 5
	switch e.CategoryID {
	case TvProgram, TvSegment, Webisode, WebSegment:
		if e.EpisodeNumber > 0 {
			params.Set("EpisodeNumber", fmt.Sprintf("%d", e.EpisodeNumber))
		}
	}

	if e.Description != "" {
		params.Set("Description", e.Description)
	}

	// can only be reported for Categories 2, 3, 8
	switch e.CategoryID {
	case TvSegment, TvExtra, Simulcast:
		if e.LinkedTitleCode != "" {
			params.Set("LinkedTitleCode", e.LinkedTitleCode)
		}
	}

	// obligatory params for categories 1, 2, 3, 8
	switch e.CategoryID {
	case TvProgram, TvSegment, TvExtra, Simulcast:
		params.Set("LiveTitle", e.LiveTitle)
		params.Set("LiveTvDay", e.LiveTvDay)
		params.Set("LiveTime", e.LiveTime)
		params.Set("LiveChannelID", fmt.Sprintf("%d", e.LiveChannelID))
	}

	// optional parameters

	if strings.HasPrefix(e.PlayURL, "http") {
		params.Set("PlayUrl", e.PlayURL)
	}

	switch e.TargetGroupCode {
	case Adults, Children:
		params.Set("TargetGroupCode", e.TargetGroupCode)
	}

	switch e.TerritoryCode {
	case Swedish, Foreign:
		params.Set("TargetGroupCode", e.TargetGroupCode)
	}

	if e.SuggestedGenre1 != "" {
		params.Set("SuggestedGenre1", e.SuggestedGenre1)
	}

	if e.SuggestedGenre2 != "" {
		params.Set("SuggestedGenre2", e.SuggestedGenre2)
	}

	if e.SuggestedGenre3 != "" {
		params.Set("SuggestedGenre3", e.SuggestedGenre3)
	}

	return params, nil
}

func (e *Episode) validate() error {
	if e.TitleCode == "" {
		return ErrorWithMessage(ErrMissingParameter, "TitleCode")
	}

	if e.SeriesCode == "" {
		return ErrorWithMessage(ErrMissingParameter, "SeriesCode")
	}

	if e.Title == "" {
		return ErrorWithMessage(ErrMissingParameter, "Title")
	}

	if e.Length < 1 {
		return ErrorWithMessage(ErrInvalidParameter, "Length")
	}

	if len(e.PublishedAt) != 8 {
		return ErrorWithMessage(ErrInvalidParameter, "PublishedAt")
	}

	if !validCategoryID(e.CategoryID) {
		return ErrorWithMessage(ErrInvalidParameter, "CategoryID")
	}

	switch e.CategoryID {
	case TvProgram, TvSegment, TvExtra, Simulcast:
		if e.LiveTitle == "" {
			return ErrorWithMessage(ErrMissingParameter, "LiveTitle")
		}

		if len(e.LiveTvDay) != 8 {
			return ErrorWithMessage(ErrInvalidParameter, "LiveTvDay")
		}

		if len(e.LiveTime) != 4 {
			return ErrorWithMessage(ErrInvalidParameter, "LiveTime")
		}

		if !validLiveChannelID(e.LiveChannelID) {
			return ErrorWithMessage(ErrInvalidParameter, "LiveChannelID")
		}
	}

	return nil
}
