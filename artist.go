package gotidal

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
