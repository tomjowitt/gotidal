package main

import (
	"context"
	"log"
	"os"

	"github.com/tomjowitt/gotidal"
)

const maxSearchResults = 5

func main() {
	ctx := context.Background()

	clientID := os.Getenv("TIDAL_CLIENT_ID")
	clientSecret := os.Getenv("TIDAL_CLIENT_SECRET")

	client, err := gotidal.NewClient(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	params := gotidal.SearchParams{
		Query:       "Peso Pluma",
		CountryCode: "MX",
		Limit:       maxSearchResults,
		Popularity:  gotidal.SearchPopularityCountry,
	}

	results, err := client.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range results.Albums {
		log.Printf("%s - %s", album.Title(), album.ArtistsToString())
		log.Printf("%d - %s", album.Duration(), album.ReleaseDate())
	}

	for _, artist := range results.Artists {
		log.Printf("%s - %s", artist.Name(), artist.URL())
	}

	for _, track := range results.Tracks {
		log.Printf("%s - %s", track.Title(), track.Album().Title())
	}
}
