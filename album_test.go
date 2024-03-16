package gotidal

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestGetSingleAlbum(t *testing.T) {
	t.Parallel()

	type args struct {
		httpClient HTTPClient
		ID         string
	}

	type expected struct {
		ID                 string
		Title              string
		ArtistCount        int
		ArtistID           string
		ArtistName         string
		ArtistPictureCount int
		IsMain             bool
		Duration           int
		ReleaseDate        string
		CoverImageCount    int
		VideoCoverCount    int
		Volumes            int
		Tracks             int
		Videos             int
		Type               string
		Copyright          string
		MetadataTags       []string
		TidalURL           string
	}

	tests := []struct {
		name     string
		args     args
		expected expected
		wantErr  bool
	}{
		{
			"Single album parses correctly",
			args{
				httpClient: &mockHTTPClient{FilePath: "testdata/single-album.json", StatusCode: http.StatusOK},
				ID:         "51584178",
			},
			expected{
				ID:                 "51584178",
				Title:              "Power Corruption and Lies",
				ArtistCount:        1,
				ArtistID:           "11950",
				ArtistName:         "New Order",
				ArtistPictureCount: 10,
				IsMain:             true,
				Duration:           2555,
				ReleaseDate:        "1983-01-01",
				CoverImageCount:    7,
				VideoCoverCount:    0,
				Volumes:            1,
				Tracks:             8,
				Videos:             0,
				Type:               "ALBUM",
				Copyright:          "© 2015 Warner Records 90 Ltd",
				MetadataTags:       []string{"LOSSLESS", "MQA"},
				TidalURL:           "https://tidal.com/browse/album/51584178",
			},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client := &Client{httpClient: tt.args.httpClient}

			album, err := client.GetSingleAlbum(context.Background(), tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetSingleAlbum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if album.ID != tt.expected.ID {
				t.Errorf("Album ID error. Want %v, Got %v", tt.expected.ID, album.ID)
			}

			if album.Title != tt.expected.Title {
				t.Errorf("Album Title error. Want %v, Got %v", tt.expected.Title, album.Title)
			}

			if len(album.Artists) != tt.expected.ArtistCount {
				t.Errorf("Album Artists error. Want %v, Got %v", tt.expected.ArtistCount, len(album.Artists))
			}

			if album.Artists[0].ID != tt.expected.ArtistID {
				t.Errorf("Album Artist ID error. Want %v, Got %v", tt.expected.ArtistID, album.Artists[0].ID)
			}

			if album.Artists[0].Name != tt.expected.ArtistName {
				t.Errorf("Album Artist Name error. Want %v, Got %v", tt.expected.ArtistName, album.Artists[0].Name)
			}

			picCount := len(album.Artists[0].Picture)
			if picCount != tt.expected.ArtistPictureCount {
				t.Errorf("Album Artist Picture error. Want %v, Got %v", tt.expected.ArtistPictureCount, picCount)
			}

			if album.Artists[0].Main != tt.expected.IsMain {
				t.Errorf("Album Artist Main error. Want %v, Got %v", tt.expected.IsMain, album.Artists[0].Main)
			}

			if album.Duration != tt.expected.Duration {
				t.Errorf("Album Duration error. Want %v, Got %v", tt.expected.Duration, album.Duration)
			}

			if album.ReleaseDate != tt.expected.ReleaseDate {
				t.Errorf("Album ReleaseDate error. Want %v, Got %v", tt.expected.ReleaseDate, album.ReleaseDate)
			}

			coverCount := len(album.ImageCover)
			if coverCount != tt.expected.CoverImageCount {
				t.Errorf("Album Cover error. Want %v, Got %v", tt.expected.CoverImageCount, coverCount)
			}

			videoCount := len(album.VideoCover)
			if videoCount != tt.expected.VideoCoverCount {
				t.Errorf("Album Video Cover error. Want %v, Got %v", tt.expected.VideoCoverCount, videoCount)
			}

			if album.NumberOfVolumes != tt.expected.Volumes {
				t.Errorf("Album Volumes error. Want %v, Got %v", tt.expected.Volumes, album.NumberOfVolumes)
			}

			if album.NumberOfTracks != tt.expected.Tracks {
				t.Errorf("Album Tracks error. Want %v, Got %v", tt.expected.Tracks, album.NumberOfTracks)
			}

			if album.NumberOfVideos != tt.expected.Videos {
				t.Errorf("Album Videos error. Want %v, Got %v", tt.expected.Videos, album.NumberOfVideos)
			}

			if album.Type != tt.expected.Type {
				t.Errorf("Album Type error. Want %v, Got %v", tt.expected.Type, album.Type)
			}

			if album.Copyright != tt.expected.Copyright {
				t.Errorf("Album Copyright error. Want %v, Got %v", tt.expected.Copyright, album.Copyright)
			}

			if !reflect.DeepEqual(album.MediaMetaData.Tags, tt.expected.MetadataTags) {
				t.Errorf("Album Metadata error. Want %v, Got %v", tt.expected.MetadataTags, album.MediaMetaData.Tags)
			}

			if album.TidalURL != tt.expected.TidalURL {
				t.Errorf("Album TidalURL error. Want %v, Got %v", tt.expected.TidalURL, album.TidalURL)
			}
		})
	}
}

func TestGetAlbumTracks(t *testing.T) {
	t.Parallel()

	type args struct {
		httpClient HTTPClient
		id         string
	}

	type expected struct {
		trackCount int
	}

	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			"Count of album tracks",
			args{
				httpClient: &mockHTTPClient{FilePath: "testdata/album-items.json", StatusCode: http.StatusOK},
				id:         "51584178",
			},
			expected{
				trackCount: 8,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Client{
				httpClient: tt.args.httpClient,
			}

			tracks, err := c.GetAlbumTracks(context.Background(), tt.args.id)
			if err != nil {
				t.Errorf("Client.GetAlbumTracks() error = %v", err)
				return
			}

			if len(tracks) != tt.expected.trackCount {
				t.Errorf("Client.GetAlbumTracks() track count = %v, want %v", len(tracks), tt.expected.trackCount)
			}
		})
	}
}

func TestGetSingleTrack(t *testing.T) {
	t.Parallel()

	type args struct {
		httpClient HTTPClient
		id         string
	}

	type expected struct {
		artist string
		title  string
		isrc   string
	}

	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			"Count of album tracks",
			args{
				httpClient: &mockHTTPClient{FilePath: "testdata/single-track.json", StatusCode: http.StatusOK},
				id:         "51584179",
			},
			expected{
				artist: "New Order",
				title:  "Age of Consent (2015 Remaster)",
				isrc:   "GBAAP1500379",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Client{
				httpClient: tt.args.httpClient,
			}

			track, err := c.GetSingleTrack(context.Background(), tt.args.id)
			if err != nil {
				t.Errorf("Client.GetSingleTrack() error = %v", err)
				return
			}

			if track.Artists[0].Name != tt.expected.artist {
				t.Errorf("Client.GetSingleTrack() track artist = %v, want %v", track.Artists[0].Name, tt.expected.artist)
			}

			if track.Title != tt.expected.title {
				t.Errorf("Client.GetSingleTrack() track title = %v, want %v", track.Title, tt.expected.title)
			}

			if track.ISRC != tt.expected.isrc {
				t.Errorf("Client.GetSingleTrack() track isrc = %v, want %v", track.ISRC, tt.expected.isrc)
			}
		})
	}
}
