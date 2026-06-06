package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ShishirSimha/pokedex-go/internal/pokecache"
)

// LocationAreaResponse represents the paginated response from /location-area
type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// LocationAreaDetail represents the detailed response from /location-area/{name}
type LocationAreaDetail struct {
	Name             string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// Create the cache
var cache = pokecache.NewCache(5 * time.Minute)

// FetchLocationAreas calls the PokeAPI and returns the response
func FetchLocationAreas(url string) (*LocationAreaResponse, error) {

	var locationResp LocationAreaResponse

	if data, found := cache.Get(url); found {
		//Its a cache hit
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
		}
		return &locationResp, nil
	}

	//Make the actual HTTP request

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

	cache.Add(url, body)

	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &locationResp, nil
}

// FetchLocationAreaDetail fetches detailed info for a named location area, with caching.
func FetchLocationAreaDetail(areaName string) (*LocationAreaDetail, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	var detail LocationAreaDetail

	if data, found := cache.Get(url); found {
		err := json.Unmarshal(data, &detail)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling cached JSON: %w", err)
		}
		return &detail, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d — is '%s' a valid location area name?", resp.StatusCode, areaName)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	cache.Add(url, body)

	err = json.Unmarshal(body, &detail)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &detail, nil
}
