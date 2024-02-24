# gotidal

<p align="center" width="100%">
    <img width="33%" src="assets/gotidal.png">
</p>

An unofficial Go library for interacting with the TIDAL API.

[![Go Reference](https://pkg.go.dev/badge/badge/.svg)](https://pkg.go.dev/github.com/tomjowitt/gotidal)
![GitHub License](https://img.shields.io/github/license/tomjowitt/gotidal)
![GitHub Tag](https://img.shields.io/github/v/tag/tomjowitt/gotidal)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/tomjowitt/gotidal/test.yml?label=tests)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/tomjowitt/gotidal/lint.yml?label=lint)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tomjowitt/gotidal)

## Guidelines

As the public TIDAL API is still in beta and subject to change, we cannot guarantee this package will always work
and therefore should not be used in production systems. We endeavor to patch and release fixes as soon as possible.

Hopefully, in the future, we will be able to match parity and provide backward compatibility with TIDAL.

## Official Documentation

Developer documentation and API keys:
<https://developer.tidal.com/>

Discussion and feature requests:
<https://github.com/orgs/tidal-music/discussions>

## Roadmap

Please see the issues tab for work left to be completed.

## Installation

```bash
go get -u github.com/tomjowitt/gotidal
```

## Usage

There are working examples for each TIDAL endpoint in the `/examples` folder.

```go
package main

import (
 ...

 "github.com/tomjowitt/gotidal"
)

const maxSearchResults = 5

func main() {
    ctx := context.Background()

    clientID := os.Getenv("TIDAL_CLIENT_ID")
    clientSecret := os.Getenv("TIDAL_CLIENT_SECRET")

    client, err := gotidal.NewClient(clientID, clientSecret, "AU")
    if err != nil {
        log.Fatal(err)
    }

    params := gotidal.SearchParams{
        Query:       "Peso Pluma",
        Limit:       maxSearchResults,
        Popularity:  gotidal.SearchPopularityCountry,
    }

    results, err := client.Search(ctx, params)
    if err != nil {
        log.Fatal(err)
    }

    for _, album := range results.Albums {
        log.Printf("%s - %s", album.Title, album.Artists[0].Name)
    }
```

## Credits

Logo created with Gopher Konstructor <https://github.com/quasilyte/gopherkon> based on original artwork
by Renee French.
