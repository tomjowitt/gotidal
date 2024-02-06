package gotidal

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	ContentType = "application/vnd.tidal.v1+json"
	Environment = "https://openapi.tidal.com"
	OauthURI    = "https://auth.tidal.com/v1/oauth2/token"
)

var ErrUnexpectedResponseCode = errors.New("returned an unexpected status code")

type Client struct {
	ContentType string
	Environment string
	Token       string
}

func NewClient(clientID string, clientSecret string) (*Client, error) {
	ctx := context.Background()
	token, err := getAccessToken(ctx, clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	return &Client{
		ContentType: ContentType,
		Environment: Environment,
		Token:       token,
	}, nil
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getAccessToken(ctx context.Context, clientID string, clientSecret string) (string, error) {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))

	client := http.Client{}
	requestBody := []byte(`grant_type=client_credentials`)
	req, err := http.NewRequestWithContext(ctx, "POST", OauthURI, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create OAuth request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to process OAuth request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: %d", ErrUnexpectedResponseCode, response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the OAuth response body: %w", err)
	}

	var authResponse AuthResponse
	err = json.Unmarshal(responseBody, &authResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal the OAuth response body: %w", err)
	}

	return authResponse.AccessToken, nil
}
