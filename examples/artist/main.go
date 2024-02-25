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
	log.Println("Single artist")
	log.Println(" ")

	artist, err := client.GetSingleArtist(ctx, "5907")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s - %s", artist.Name, artist.URL)

	log.Println("-------------------------------------------------")
	log.Println("Get albums for an artist")
	log.Println(" ")

	albums, err := client.GetAlbumsByArtist(ctx, "5907", gotidal.PaginationParams{Limit: 5})
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albums {
		log.Printf("%s - %s", album.Title, album.Artists[0].Name)
	}

	log.Println("-------------------------------------------------")
	log.Println("Get multiple artists")
	log.Println(" ")

	artists, err := client.GetMultipleArtists(ctx, []string{"5907", "3502119", "31874"})
	if err != nil {
		log.Fatal(err)
	}

	for _, artist := range artists {
		log.Printf("%s - %s", artist.Name, artist.URL)
	}
}
