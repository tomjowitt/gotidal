package main

import (
	"log"
	"os"

	"github.com/tomjowitt/gotidal"
)

func main() {
	clientID := os.Getenv("TIDAL_CLIENT_ID")
	clientSecret := os.Getenv("TIDAL_CLIENT_SECRET")

	client, err := gotidal.NewClient(clientID, clientSecret, "AU")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------------------------------------------------")
	log.Println("OAuth Token")
	log.Println("-------------------------------------------------")

	log.Println(client.Token)
}
