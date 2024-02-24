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
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

const (
	contentType = "application/vnd.tidal.v1+json"
	environment = "https://openapi.tidal.com"
	oauthURI    = "https://auth.tidal.com/v1/oauth2/token"
)

var ErrUnexpectedResponseCode = errors.New("returned an unexpected status code")

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient  HTTPClient
	ContentType string
	Environment string
	Token       string
	CountryCode string
}

// NewClient returns an API client based on a users credentials and location.
func NewClient(clientID string, clientSecret string, countryCode string) (*Client, error) {
	ctx := context.Background()

	httpClient := &http.Client{}

	token, err := getAccessToken(ctx, httpClient, clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient:  httpClient,
		ContentType: contentType,
		Environment: environment,
		Token:       token,
		CountryCode: countryCode,
	}, nil
}

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getAccessToken(ctx context.Context, httpClient HTTPClient, clientID string, clientSecret string) (string, error) {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))

	requestBody := []byte(`grant_type=client_credentials`)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, oauthURI, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create OAuth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", concat("Basic ", basicAuth))

	responseBody, err := processRequest(httpClient, req)
	if err != nil {
		return "", fmt.Errorf("failed to process the request: %w", err)
	}

	var authResponse authResponse

	err = json.Unmarshal(responseBody, &authResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal the OAuth response body: %w", err)
	}

	return authResponse.AccessToken, nil
}

func processRequest(httpClient HTTPClient, req *http.Request) ([]byte, error) {
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to process request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusMultiStatus {
		return nil, fmt.Errorf("%w: %d", ErrUnexpectedResponseCode, response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the OAuth response body: %w", err)
	}

	return responseBody, nil
}

// nolint:unparam
func (c *Client) request(ctx context.Context, method string, path string, params any) ([]byte, error) {
	uri := fmt.Sprintf("%s%s?%s", c.Environment, path, toURLParams(params, c.CountryCode))

	req, err := http.NewRequestWithContext(ctx, method, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", uri, err)
	}

	req.Header.Set("Content-Type", c.ContentType)
	req.Header.Set("Authorization", concat("Bearer ", c.Token))
	req.Header.Set("accept", c.ContentType)

	return processRequest(c.httpClient, req)
}

func toURLParams(input interface{}, countryCode string) string {
	var params []string
	params = append(params, fmt.Sprintf("%s=%s", "countryCode", countryCode))

	if input == nil {
		return strings.Join(params, "&")
	}

	v := reflect.ValueOf(input)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.IsValid() {
			var paramValue string

			switch value.Kind() {
			case reflect.String:
				paramValue = value.String()
			case reflect.Int:
				paramValue = strconv.FormatInt(value.Int(), 10)
			default:
				continue
			}

			if paramValue != "" && paramValue != "0" {
				paramName := url.QueryEscape(lowercaseFirstLetter(field.Name))
				paramValue = url.QueryEscape(paramValue)

				params = append(params, fmt.Sprintf("%s=%s", paramName, paramValue))
			}
		}
	}

	return strings.Join(params, "&")
}

// lowercaseFirstLetter converts the first letter of a string the lowercase to match the camel-casing of the TIDAL
// API URL parameters.
func lowercaseFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}

	firstChar := []rune(str)[0]
	lowerFirstChar := unicode.ToLower(firstChar)

	return string(lowerFirstChar) + str[1:]
}

// concat joins strings together in a more efficient way than fmt.Sprintf.
func concat(strs ...string) string {
	return strings.Join(strs, "")
}
