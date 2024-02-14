package gotidal

type Artist struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Picture []Image `json:"picture"`
	Main    bool    `json:"main"`
}
