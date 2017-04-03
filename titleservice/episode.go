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
	TitleCode       string        `json:"title_code"`                  // required
	SeriesCode      string        `json:"series_code"`                 // required
	Title           string        `json:"title"`                       // required
	Length          int           `json:"length"`                      // required
	PublishedAt     string        `json:"published_at"`                // required (YYYYMMDD)
	AvailableUntil  string        `json:"available_until,omitempty"`   // optional (YYYYMMDD)
	CategoryID      CategoryID    `json:"category_id"`                 // required
	EpisodeNumber   int           `json:"episode_number,omitempty"`    // only applicable to Categories 1, 2, 4, 5
	Description     string        `json:"description,omitempty"`       // optional
	LinkedTitleCode string        `json:"linked_title_code,omitempty"` // can only be reported for categories 2, 3, 8 (only for updates)
	LiveTitle       string        `json:"live_title,omitempty"`        // obligatory for categories 1, 2, 3, 8
	LiveTvDay       string        `json:"live_tv_day,omitempty"`       // obligatory for categories 1, 2, 3, 8 (YYYYMMDD)
	LiveTime        string        `json:"live_time,omitempty"`         // obligatory for categories 1, 2, 3, 8 (HHMM, MMS-time: 23:45=2345, 01:45=2545, 02:00=0200)
	LiveChannelID   LiveChannelID `json:"live_channel_id,omitempty"`   // obligatory for categories 1, 2, 3, 8
	PlayURL         string        `json:"play_url,omitempty"`          // maximum of 150 characters
	TargetGroupCode string        `json:"target_group_code,omitempty"` // optional V = Vuxen (Adults) B = Barn (Children)
	TerritoryCode   string        `json:"territory_code,omitempty"`    // optional S = Svenskt (Swedish) U = Utländskt (Foreign)
	SuggestedGenre1 string        `json:"suggested_genre_1,omitempty"` // free text of maximum of 256 characters in length
	SuggestedGenre2 string        `json:"suggested_genre_2,omitempty"` // free text of maximum of 256 characters in length
	SuggestedGenre3 string        `json:"suggested_genre_3,omitempty"` // free text of maximum of 256 characters in length
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
		return newErrorWithMessage(ErrMissingParameter, "TitleCode")
	}

	if e.SeriesCode == "" {
		return newErrorWithMessage(ErrMissingParameter, "SeriesCode")
	}

	if e.Title == "" {
		return newErrorWithMessage(ErrMissingParameter, "Title")
	}

	if e.Length < 1 {
		return newErrorWithMessage(ErrInvalidParameter, "Length")
	}

	if len(e.PublishedAt) != 8 {
		return newErrorWithMessage(ErrInvalidParameter, "PublishedAt")
	}

	if !validCategoryID(e.CategoryID) {
		return newErrorWithMessage(ErrInvalidParameter, "CategoryID")
	}

	switch e.CategoryID {
	case TvProgram, TvSegment, TvExtra, Simulcast:
		if e.LiveTitle == "" {
			return newErrorWithMessage(ErrMissingParameter, "LiveTitle")
		}

		if len(e.LiveTvDay) != 8 {
			return newErrorWithMessage(ErrInvalidParameter, "LiveTvDay")
		}

		if len(e.LiveTime) != 4 {
			return newErrorWithMessage(ErrInvalidParameter, "LiveTime")
		}

		if !validLiveChannelID(e.LiveChannelID) {
			return newErrorWithMessage(ErrInvalidParameter, "LiveChannelID")
		}
	}

	return nil
}
