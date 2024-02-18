package gotidal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Album represents an individual release.
type Album struct {
	albumResource `json:"resource"`
}

type albumResource struct {
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
}

type MediaMetaData struct {
	Tags []string `json:"tags"`
}

type AlbumProperties struct {
	Content []string `json:"content"`
}

type ProviderInfo struct {
	ID   string `json:"providerId"`
	Name string `json:"providerName"`
}

// Track represents an individual track on an album.
type Track struct {
	trackResource `json:"resource"`
}

type trackResource struct {
	ID      string        `json:"id"`
	Title   string        `json:"title"`
	Version string        `json:"version"`
	Artists []Artist      `json:"artists"`
	Album   albumResource `json:"album"`
}

type albumResults struct {
	Data []Album `json:"data"`
}

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
		return nil, fmt.Errorf("failed to connect to the album endpoint: %w", err)
	}

	var results albumResults

	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the album response body: %w", err)
	}

	return results.Data, nil
}
