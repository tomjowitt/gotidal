package main

import (
	"context"
	"log"
	"os"

	"github.com/tomjowitt/gotidal"
)

func main() {
	ctx := context.Background()

	clientID := os.Getenv("TIDAL_CLIENT_ID")
	clientSecret := os.Getenv("TIDAL_CLIENT_SECRET")

	client, err := gotidal.NewClient(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	params := gotidal.SearchParams{
		Query:       "Black Flag",
		CountryCode: "AU",
	}

	results, err := client.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range results.Albums {
		log.Printf("%s - %s", album.Resource.Title, album.Resource.Artists[0].Name)
	}
}
