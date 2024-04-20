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

	track, err := client.GetSingleTrack(ctx, "51584179")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------------------------------------------------")
	log.Println("Single Track")
	log.Println("-------------------------------------------------")

	log.Printf("%s - %s", track.Title, track.Artists[0].Name)

	log.Println("-------------------------------------------------")
	log.Println("Tracks By ISRC")
	log.Println("-------------------------------------------------")

	tracks, err := client.GetTracksByISRC(ctx, "GBBLY1600675", gotidal.PaginationParams{Limit: 5})
	if err != nil {
		log.Fatal(err)
	}

	for _, track := range tracks {
		log.Printf("%s - %s - %s", track.Title, track.Artists[0].Name, track.Album.Title)
	}

	log.Println("-------------------------------------------------")
	log.Println("Multiple Tracks")
	log.Println("-------------------------------------------------")

	multipleTracks, err := client.GetMultipleTracks(ctx, []string{"251380837", "251380838"})
	if err != nil {
		log.Fatal(err)
	}

	for _, track := range multipleTracks {
		log.Printf("%s - %s - %s", track.Title, track.Artists[0].Name, track.Album.Title)
	}
}
