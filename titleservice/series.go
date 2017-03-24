package titleservice

import (
	"fmt"
	"net/url"
)

// Series is a title linked to a series/season
type Series struct {
	SeriesCode   string // required
	Title        string // required
	SeasonNumber int
	Description  string
	GenreText    string
}

// Endpoint returns the endpoint to use for this request type
func (s *Series) Endpoint() Endpoint {
	return RegisterSeriesEndpoint
}

// Params validates and returns the parameters sent to the MMS TitleService API for this request type
func (s *Series) Params() (url.Values, error) {
	if err := s.validate(); err != nil {
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

func (s *Series) validate() error {
	if s.SeriesCode == "" {
		return ErrorWithMessage(ErrMissingParameter, "SeriesCode")
	}

	if s.Title == "" {
		return ErrorWithMessage(ErrMissingParameter, "Title")
	}

	return nil
}
