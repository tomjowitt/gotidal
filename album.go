package gotidal

type Album struct {
	Resource Resource `json:"resource"`
	ID       string   `json:"id"`
	Status   int      `json:"status"`
	Message  string   `json:"message"`
}

type Resource struct {
	ID              string        `json:"id"`
	BarcodeID       string        `json:"barcodeID"`
	Title           string        `json:"title"`
	Artists         []Artist      `json:"artists"`
	Duration        int           `json:"duration"`
	ReleaseDate     string        `json:"releaseDate"`
	ImageCover      []Image       `json:"imageCover"`
	VideoCover      []Image       `json:"videoCover"`
	NumberOfVolumes int           `json:"numberOfVolumes"`
	NumberOfTracks  int           `json:"numberOfTracks"`
	NumberOfVideos  int           `json:"numberOfVideos"`
	Type            string        `json:"type"`
	Copyright       string        `json:"copyright"`
	MediaMetadata   MediaMetaData `json:"mediaMetadata"`
	Properties      Properties    `json:"properties"`
	TidalURL        string        `json:"tidalUrl"`
}

type MediaMetaData struct {
	Tags []string `json:"tags"`
}

type Properties struct {
	Content []string `json:"content"`
}
