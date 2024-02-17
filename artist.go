package gotidal

// Artist represents an individual artist.
type Artist struct {
	Resource artistResource `json:"resource"`
}

// ID returns the ID of the artist.
func (a Artist) ID() string {
	return a.Resource.ID
}

// Name returns the name of the artist.
func (a Artist) Name() string {
	return a.Resource.Name
}

// URL returns the TIDAL URL of the artist.
func (a Artist) URL() string {
	return a.Resource.TidalURL
}

// Pictures returns a list of pictures associated with the artist.
func (a Artist) Pictures() []Image {
	return a.Resource.Picture
}

type artistResource struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Picture  []Image `json:"picture"`
	TidalURL string  `json:"tidalURL"`
	Main     bool    `json:"main"`
}
