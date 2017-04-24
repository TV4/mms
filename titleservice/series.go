package titleservice

import (
	"fmt"
	"net/url"
)

// MakeSeries creates a Series based on required parameters and optional parameters using options
func MakeSeries(seriesCode, title string, options ...func(*Series)) Series {
	s := &Series{
		SeriesCode: seriesCode,
		Title:      title,
	}

	for _, f := range options {
		f(s)
	}

	return *s
}

// Series is a title linked to a series/season
type Series struct {
	SeriesCode   string `json:"series_code"`             // required
	Title        string `json:"title"`                   // required
	SeasonNumber int    `json:"season_number,omitempty"` // optional
	Description  string `json:"description,omitempty"`   // optional
	GenreText    string `json:"genre_text,omitempty"`    // optional
}

// Endpoint returns the endpoint to use for this request type
func (s *Series) Endpoint() Endpoint {
	return RegisterSeriesEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (s *Series) Params() (url.Values, error) {
	if err := s.Validate(); err != nil {
		return url.Values{}, err
	}

	params := url.Values{
		"SeriesCode": {s.SeriesCode},
		"Title":      {s.Title},
	}

	// optional parameters

	if s.SeasonNumber > 0 {
		params.Set("SeasonNumber", fmt.Sprintf("%d", s.SeasonNumber))
	}

	if s.Description != "" {
		params.Set("Description", s.Description)
	}

	if s.GenreText != "" {
		params.Set("GenreText", s.GenreText)
	}

	return params, nil
}

// Validate all fields
func (s *Series) Validate() error {
	if s.SeriesCode == "" {
		return newErrorWithMessage(ErrMissingParameter, "SeriesCode")
	}

	if s.Title == "" {
		return newErrorWithMessage(ErrMissingParameter, "Title")
	}

	return nil
}
