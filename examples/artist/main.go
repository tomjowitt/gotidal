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

	client, err := gotidal.NewClient(clientID, clientSecret, "AU")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------------------------------------------------")
	log.Println("Single Artist")
	log.Println(" ")

	artist, err := client.GetSingleArtist(ctx, "5907")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s - %s", artist.Name, artist.URL)
}
