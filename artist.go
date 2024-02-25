package gotidal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Artist represents an individual artist.
type Artist struct {
	artistResource `json:"resource"`
}

type artistResource struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Picture []Image `json:"picture"`
	URL     string  `json:"tidalURL"`
	Main    bool    `json:"main"`
}

// GetSingleArtist returns an artist that matches an ID.
func (c *Client) GetSingleArtist(ctx context.Context, id string) (*Artist, error) {
	if id == "" {
		return nil, ErrMissingRequiredParameters
	}

	response, err := c.request(ctx, http.MethodGet, concat("/artists/", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the artists endpoint: %w", err)
	}

	var result Artist

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the artist response body: %w", err)
	}

	return &result, nil
}

// GetAlbumsByArtist returns a paginated list of albums for an artist.
func (c *Client) GetAlbumsByArtist(ctx context.Context, id string, params PaginationParams) ([]Album, error) {
	if id == "" {
		return nil, ErrMissingRequiredParameters
	}

	response, err := c.request(ctx, http.MethodGet, concat("/artists/", id, "/albums"), params)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the artist albums endpoint: %w", err)
	}

	var results albumResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the artist albums response body: %w", err)
	}

	return results.Data, nil
}

type artistResults struct {
	Data []Artist `json:"data"`
}

// GetMultipleArtists returns a list of artists filtered by their IDs.
func (c *Client) GetMultipleArtists(ctx context.Context, ids []string) ([]Artist, error) {
	type multiArtistParams struct {
		ids string
	}

	params := multiArtistParams{
		ids: strings.Join(ids, ","),
	}

	response, err := c.request(ctx, http.MethodGet, "/artists", params)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the multiple artists endpoint: %w", err)
	}

	var results artistResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the multiple artists response body: %w", err)
	}

	return results.Data, nil
}
