package gotidal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Album represents an individual release.
type Album struct {
	AlbumResource `json:"resource"`
}

type AlbumResource struct {
	ID              string           `json:"id"`
	BarcodeID       string           `json:"barcodeID"`
	Title           string           `json:"title"`
	Artists         []artistResource `json:"artists"`
	Duration        int              `json:"duration"`
	ReleaseDate     string           `json:"releaseDate"`
	ImageCover      []Image          `json:"imageCover"`
	VideoCover      []Image          `json:"videoCover"`
	NumberOfVolumes int              `json:"numberOfVolumes"`
	NumberOfTracks  int              `json:"numberOfTracks"`
	NumberOfVideos  int              `json:"numberOfVideos"`
	Type            string           `json:"type"`
	Copyright       string           `json:"copyright"`
	MediaMetaData   MediaMetaData    `json:"mediaMetadata"`
	Properties      AlbumProperties  `json:"properties"`
	TidalURL        string           `json:"tidalUrl"`
	ProviderInfo    ProviderInfo     `json:"providerInfo"`
}

// MediaMetaData represents the metadata of an album.
type MediaMetaData struct {
	Tags []string `json:"tags"`
}

// AlbumProperties represents the properties of an album.
type AlbumProperties struct {
	Content []string `json:"content"`
}

// ProviderInfo represents the provider of an album.
type ProviderInfo struct {
	ID   string `json:"providerId"`
	Name string `json:"providerName"`
}

// Track represents an individual track on an album.
type Track struct {
	trackResource `json:"resource"`
}

type trackResource struct {
	ID            string           `json:"id"`
	ArtifactType  string           `json:"artifactType"`
	Title         string           `json:"title"`
	ISRC          string           `json:"isrc"`
	Copyright     string           `json:"copyright"`
	Version       string           `json:"version"`
	Artists       []artistResource `json:"artists"`
	Album         AlbumResource    `json:"album"`
	TrackNumber   int              `json:"trackNumber"`
	VolumeNumber  int              `json:"volumeNumber"`
	MediaMetaData MediaMetaData    `json:"mediaMetadata"`
	Properties    AlbumProperties  `json:"properties"`
	TidalURL      string           `json:"tidalUrl"`
	ProviderInfo  ProviderInfo     `json:"providerInfo"`
}

type albumResults struct {
	Data []Album `json:"data"`
}

type ItemMetaData struct {
	Total int `json:"total"`
}

// GetSingleAlbum returns an album that matches an ID.
func (c *Client) GetSingleAlbum(ctx context.Context, id string) (*Album, error) {
	if id == "" {
		return nil, ErrMissingRequiredParameters
	}

	response, err := c.request(ctx, http.MethodGet, concat("/albums/", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the albums endpoint: %w", err)
	}

	var result Album

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the albums response body: %w", err)
	}

	return &result, nil
}

type trackResults struct {
	Data     []Track      `json:"data"`
	MetaData ItemMetaData `json:"metadata"`
}

const paginationLimit = 100

// GetAlbumTracks returns a list of album tracks.
//
// The items endpoint is paginated so we set a fairly high limit (100 items) in the hope to catch most cases in
// one round-trip. If the metadata reports a higher total then we make susequent API calls until all the tracks are
// returned.
//
// This endpoint also supports videos but it was hard to find any examples of this, so for the moment this is tracks
// only.
func (c *Client) GetAlbumTracks(ctx context.Context, id string) ([]Track, error) {
	if id == "" {
		return nil, ErrMissingRequiredParameters
	}

	params := PaginationParams{
		Limit:  paginationLimit,
		Offset: 0,
	}

	total := 0
	runningTotal := 0

	var tracks []Track

	for total >= runningTotal {
		response, err := c.request(ctx, http.MethodGet, concat("/albums/", id, "/items"), params)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to the albums endpoint: %w", err)
		}

		var results trackResults

		err = json.Unmarshal(response, &results)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal the albums response body: %w", err)
		}

		tracks = append(tracks, results.Data...)

		params.Offset += params.Limit
		runningTotal += params.Limit
		total = results.MetaData.Total
	}

	return tracks, nil
}

// GetAlbumByBarcodeID returns a list of albums that match a barcode ID.
func (c *Client) GetAlbumByBarcodeID(ctx context.Context, barcodeID string) ([]Album, error) {
	if barcodeID == "" {
		return nil, ErrMissingRequiredParameters
	}

	type barcodeParams struct {
		barcodeId string // nolint:revive // This variable is directly referenced in the query string.
	}

	params := barcodeParams{
		barcodeId: barcodeID,
	}

	response, err := c.request(ctx, http.MethodGet, "/albums/byBarcodeId", params)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the albums endpoint: %w", err)
	}

	var results albumResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the albums response body: %w", err)
	}

	return results.Data, nil
}

// GetMultipleAlbums returns a list of albums filtered by their IDs.
func (c *Client) GetMultipleAlbums(ctx context.Context, ids []string) ([]Album, error) {
	type multiAlbumParams struct {
		ids string
	}

	params := multiAlbumParams{
		ids: strings.Join(ids, ","),
	}

	response, err := c.request(ctx, http.MethodGet, "/albums/byIds", params)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the multiple albums endpoint: %w", err)
	}

	var results albumResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the multiple albums response body: %w", err)
	}

	return results.Data, nil
}

type similarAlbum struct {
	Resource struct {
		ID string `json:"id"`
	}
}

type similarAlbumResults struct {
	Data     []similarAlbum `json:"data"`
	MetaData ItemMetaData   `json:"metadata"`
}

// GetSimilarAlbums returns a slice of album IDs that can be used as a parameter in the GetMultipleAlbums function.
func (c *Client) GetSimilarAlbums(ctx context.Context, id string, params PaginationParams) ([]string, error) {
	response, err := c.request(ctx, http.MethodGet, concat("/albums/", id, "/similar"), params)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the similar albums endpoint: %w", err)
	}

	var results similarAlbumResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the similar albums response body: %w", err)
	}

	var albumIDs []string
	for _, albumID := range results.Data {
		albumIDs = append(albumIDs, albumID.Resource.ID)
	}

	return albumIDs, nil
}

// GetAlbumsByArtist returns a list of albums that match an artist ID.
func (c *Client) GetSingleTrack(ctx context.Context, id string) (*Track, error) {
	response, err := c.request(ctx, http.MethodGet, concat("/tracks/", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the tracks endpoint: %w", err)
	}

	var result Track

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the tracks response body: %w", err)
	}

	return &result, nil
}

// GetTracksByISRC returns a list of tracks that match an ISRC.
//
// ISRC lookup can be found here. This is a useful tool for finding ISRCs for testing purposes:
// https://isrcsearch.ifpi.org/
func (c *Client) GetTracksByISRC(ctx context.Context, isrc string, params PaginationParams) ([]Track, error) {
	type isrcParams struct {
		isrc   string
		Limit  int
		Offset int
	}

	response, err := c.request(ctx, http.MethodGet, "/tracks/byIsrc", isrcParams{
		isrc:   isrc,
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the tracks endpoint: %w", err)
	}

	var result trackResults

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the tracks response body: %w", err)
	}

	return result.Data, nil
}
