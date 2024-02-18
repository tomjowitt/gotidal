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

	log.Println(client.Token)

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

	log.Println("-------------------------------------------------")
	log.Println("Albums")
	log.Println("-------------------------------------------------")

	for _, album := range results.Albums {
		log.Printf("%s - %s", album.Title, album.Artists[0].Name)
		log.Printf("%d - %s", album.Duration, album.ReleaseDate)
	}

	log.Println("-------------------------------------------------")
	log.Println("Artists")
	log.Println("-------------------------------------------------")

	for _, artist := range results.Artists {
		log.Printf("%s - %s", artist.Name, artist.URL)
	}

	log.Println("-------------------------------------------------")
	log.Println("Tracks")
	log.Println("-------------------------------------------------")

	for _, track := range results.Tracks {
		log.Printf("%s - %s", track.Title, track.Album.Title)
	}

	log.Println("-------------------------------------------------")
	log.Println("Videos")
	log.Println("-------------------------------------------------")

	for _, video := range results.Videos {
		log.Printf("%s - %s", video.Title, video.Artists[0].Name)
	}
}
