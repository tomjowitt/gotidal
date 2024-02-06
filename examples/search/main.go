package main

import (
	"log"
	"os"

	"github.com/tomjowitt/gotidal"
)

func main() {
	clientId := os.Getenv("TIDAL_CLIENT_ID")
	clientSecret := os.Getenv("TIDAL_CLIENT_SECRET")

	client, err := gotidal.NewClient(clientId, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	client.Search("searchQuery")
}
