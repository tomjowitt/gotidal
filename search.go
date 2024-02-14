package gotidal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SearchTypeAlbums          = "ALBUMS"
	SearchTypeArtists         = "ARTISTS"
	SearchTypeTracks          = "TRACKS"
	SearchTypeVideos          = "VIDEOS"
	SearchPopularityWorldwide = "WORLDWIDE"
	SearchPopularityCountry   = "COUNTRY"
)

// SearchParams defines the request parameters used by the TIDAL search API endpoint.
// See: https://developer.tidal.com/apiref?spec=search&ref=search
type SearchParams struct {
	// Search query in plain text.
	// Example: Beyoncé
	Query string `json:"query"`

	// Target search type. Optional. Searches for all types if not specified.
	// Example: ARTISTS, ALBUMS, TRACKS, VIDEOS
	Type string `json:"type"`

	// Pagination offset (in number of items). Required if 'query' is provided.
	// Example: 0
	Offset int `json:"offset"`

	// Page size. Required if 'query' is provided.
	// Example: 10
	Limit int `json:"limit"`

	// ISO 3166-1 alpha-2 country code.
	// Example: AU
	CountryCode string `json:"countryCode"`

	// Specify which popularity type to apply for query result: either worldwide or country popularity.
	// Worldwide popularity is using by default if nothing is specified.
	// Example: WORLDWIDE, COUNTRY
	Popularity string `json:"popularity"`
}

type SearchResults struct {
	Albums []Album
}

func (c *Client) Search(ctx context.Context, params SearchParams) (*SearchResults, error) {
	response, err := c.Request(ctx, http.MethodGet, "/search", params)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the search endpoint: %w", err)
	}

	var results SearchResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the search response body: %w", err)
	}

	return &results, nil
}
