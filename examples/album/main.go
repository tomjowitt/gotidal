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
	log.Println("Single Album")
	log.Println(" ")

	album, err := client.GetSingleAlbum(ctx, "51584178")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s - %s", album.Title, album.Artists[0].Name)

	log.Println("-------------------------------------------------")
	log.Println("Albums By Barcode ID")
	log.Println(" ")

	results, err := client.GetAlbumByBarcodeID(ctx, "197189111396")
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range results {
		log.Printf("%s - %s", album.Title, album.Artists[0].Name)
	}

	log.Println("-------------------------------------------------")
	log.Println("Multiple Albums")
	log.Println(" ")

	multiAlbums, err := client.GetMultipleAlbums(ctx, []string{"301846384", "2579864"})
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range multiAlbums {
		log.Printf("%s - %s", album.Title, album.Artists[0].Name)
	}

	log.Println("-------------------------------------------------")
	log.Println("Album Tracks")
	log.Println(" ")

	items, err := client.GetAlbumTracks(ctx, "37267701")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		log.Printf("#%d (Vol #%d) - %s - %s", item.TrackNumber, item.VolumeNumber, item.Title, item.Album.Title)
	}

	log.Println("-------------------------------------------------")
	log.Println("Similar Albums to 'New York Dolls - s/t'")
	log.Println(" ")

	similarIDs, err := client.GetSimilarAlbums(ctx, "3992356", gotidal.PaginationParams{Limit: 10})
	if err != nil {
		log.Fatal(err)
	}

	similarAlbums, err := client.GetMultipleAlbums(ctx, similarIDs)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range similarAlbums {
		log.Printf("%s - %s", item.Title, item.Artists[0].Name)
	}
}
