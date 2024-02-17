package gotidal

import (
	"strings"
	"time"
)

// Album represents an individual release.
type Album struct {
	Resource albumResource `json:"resource"`
	ID       string        `json:"id"`
	Status   int           `json:"status"`
	Message  string        `json:"message"`
}

// BarcodeId returns the bardcode ID of the album.
func (a Album) BarcodeID() string {
	return a.Resource.BarcodeID
}

// Title returns the title of the album.
func (a Album) Title() string {
	return a.Resource.Title
}

// Artists returns a list of artists that appear on the album.
func (a Album) Artists() []Artist {
	var artists []Artist
	for _, artist := range a.Resource.Artists {
		artists = append(artists, Artist{Resource: artist})
	}

	return artists
}

// ArtistsToString returns a string of artists that appear on the album.
func (a Album) ArtistsToString() string {
	var artistNames []string
	for _, artist := range a.Artists() {
		artistNames = append(artistNames, artist.Name())
	}

	return strings.Join(artistNames, " / ")
}

// Duration returns the duration of the album in seconds.
func (a Album) Duration() int {
	return a.Resource.Duration
}

// ReleaseDate returns the release date of the album as a time object.
func (a Album) ReleaseDate() *time.Time {
	layout := "2006-01-02"

	releaseDate, err := time.Parse(layout, a.Resource.ReleaseDate)
	if err != nil {
		return nil
	}

	return &releaseDate
}

// CoverImages returns a list of album images.
func (a Album) CoverImages() []Image {
	return a.Resource.ImageCover
}

// VideoImages returns a list of video images.
func (a Album) VideoImages() []Image {
	return a.Resource.VideoCover
}

// NumberOfVolumes returns the number of volumes in the album.
func (a Album) NumberOfVolumes() int {
	return a.Resource.NumberOfVolumes
}

// NumberOfTracks returns the number of tracks on the album.
func (a Album) NumberOfTracks() int {
	return a.Resource.NumberOfTracks
}

// NumberOfVideos returns the number of videos on the album.
func (a Album) NumberOfVideos() int {
	return a.Resource.NumberOfVideos
}

// Type returns the type of the album.
func (a Album) Type() string {
	return a.Resource.Type
}

// Copyright returns the compyright string for the album.
func (a Album) Copyright() string {
	return a.Resource.Copyright
}

// Metadata returns the metadata for the album.
func (a Album) Metadata() MediaMetaData {
	return a.Resource.MediaMetaData
}

// Properties returns a list of the properties for the album.
func (a Album) Properties() []string {
	return a.Resource.Properties.Content
}

// URL returns the TIDAL URL for the album.
func (a Album) URL() string {
	return a.Resource.TidalURL
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

type AlbumProviderInfo struct {
	ID   string `json:"providerId"`
	Name string `json:"providerName"`
}

// Track represents an individual track on an album.
type Track struct {
	Resource trackResource `json:"resource"`
}

// ID returns the ID of the track.
func (t Track) ID() string {
	return t.Resource.ID
}

// Title returns the title of the track.
func (t Track) Title() string {
	return t.Resource.Title
}

// Version returns the version of the track.
func (t Track) Version() string {
	return t.Resource.Version
}

// Artists returns a list of Artists that appear on the track.
func (t Track) Artists() []Artist {
	return t.Resource.Artists
}

// Album returns the Album that the track appears on.
func (t Track) Album() Album {
	return Album{Resource: t.Resource.Album}
}

type trackResource struct {
	ID      string        `json:"id"`
	Title   string        `json:"title"`
	Version string        `json:"version"`
	Artists []Artist      `json:"artists"`
	Album   albumResource `json:"album"`
}
