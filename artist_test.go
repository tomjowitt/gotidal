package gotidal

import (
	"context"
	"net/http"
	"testing"
)

func TestGetSingleArtist(t *testing.T) {
	t.Parallel()

	type args struct {
		httpClient HTTPClient
		id         string
	}

	type expected struct {
		ID           string
		Name         string
		PictureCount int
		URL          string
	}

	tests := []struct {
		name     string
		args     args
		expected expected
		wantErr  bool
	}{
		{
			"Single artist parses correctly",
			args{
				httpClient: &mockHTTPClient{FilePath: "testdata/single-artist.json", StatusCode: http.StatusOK},
				id:         "51584178",
			},
			expected{
				ID:           "5907",
				Name:         "Kronos Quartet",
				PictureCount: 10,
				URL:          "https://tidal.com/browse/artist/5907",
			},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client := &Client{httpClient: tt.args.httpClient}

			artist, err := client.GetSingleArtist(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetSingleArtist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if artist.ID != tt.expected.ID {
				t.Errorf("Client.GetSingleArtist() ID = %v, want %v", artist.ID, tt.expected.ID)
			}

			if artist.Name != tt.expected.Name {
				t.Errorf("Client.GetSingleArtist() Name = %v, want %v", artist.Name, tt.expected.Name)
			}

			if len(artist.Picture) != tt.expected.PictureCount {
				t.Errorf("Client.GetSingleArtist() Picture = %v, want %v", len(artist.Picture), tt.expected.PictureCount)
			}

			if artist.URL != tt.expected.URL {
				t.Errorf("Client.GetSingleArtist() URL = %v, want %v", artist.URL, tt.expected.URL)
			}
		})
	}
}

func TestGetAlbumsByArtist(t *testing.T) {
	t.Parallel()

	type args struct {
		httpClient HTTPClient
		id         string
	}

	type expected struct {
		AlbumCount int
	}

	tests := []struct {
		name     string
		args     args
		expected expected
		wantErr  bool
	}{
		{
			"Artist albums parses correctly",
			args{
				httpClient: &mockHTTPClient{FilePath: "testdata/albums-by-artist.json", StatusCode: http.StatusOK},
				id:         "5907",
			},
			expected{
				AlbumCount: 10,
			},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client := &Client{httpClient: tt.args.httpClient}

			albums, err := client.GetAlbumsByArtist(context.Background(), tt.args.id, PaginationParams{Limit: 10})
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAlbumsByArtist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(albums) != tt.expected.AlbumCount {
				t.Errorf("Client.GetAlbumsByArtist() AlbumCount = %v, want %v", len(albums), tt.expected.AlbumCount)
			}
		})
	}
}
