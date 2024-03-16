package gotidal

// Video represents an individula video.
type Video struct {
	videoResource `json:"resource"`
}

type videoResource struct {
	ID           string           `json:"id"`
	Title        string           `json:"title"`
	Version      string           `json:"version"`
	Images       []Image          `json:"image"`
	Album        AlbumResource    `json:"album"`
	ReleaseDate  string           `json:"releaseDate"`
	Artists      []artistResource `json:"artists"`
	Duration     int              `json:"duration"`
	TrackNumber  int              `json:"trackNumber"`
	VolumeNumber int              `json:"volumeNumber"`
	ISRC         string           `json:"isrc"`
	Copyright    string           `json:"copyright"`
	Properties   VideoProperties  `json:"properties"`
	TidalURL     string           `json:"tidalUrl"`
	ProviderInfo ProviderInfo     `json:"providerInfo"`
}

type VideoProperties struct {
	Content   []string `json:"content"`
	VideoType string   `json:"video-type"`
}
