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

	results, err := client.GetAlbumByBarcodeID(ctx, "197189111396")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------------------------------------------------")
	log.Println("Albums By Barcode ID")
	log.Println("-------------------------------------------------")

	for _, album := range results {
		log.Printf("%s - %s", album.Title, album.Artists[0].Name)
	}
}
