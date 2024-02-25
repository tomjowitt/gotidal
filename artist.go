package gotidal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
