package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreaResponse represents the paginated response from /location-area
type LocationAreaResponse struct {
	Count    int `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// FetchLocationAreas calls the PokeAPI and returns the response
func FetchLocationAreas(url string) (*LocationAreaResponse, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &locationResp, nil
}
