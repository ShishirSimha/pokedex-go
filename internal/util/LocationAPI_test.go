package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestFetchLocationAreas(t *testing.T) {
	expectedResp := LocationAreaResponse{
		Count: 2,
		Next:  nil,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "canalave-city-area", URL: "https://pokeapi.co/api/v2/location-area/1/"},
			{Name: "eterna-city-area", URL: "https://pokeapi.co/api/v2/location-area/2/"},
		},
	}

	var requestCount int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	// 1. First fetch (should hit mock server, populate cache)
	resp1, err := FetchLocationAreas(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp1.Count != expectedResp.Count {
		t.Errorf("expected count %d, got %d", expectedResp.Count, resp1.Count)
	}

	if atomic.LoadInt32(&requestCount) != 1 {
		t.Errorf("expected server to be called 1 time, got %d", requestCount)
	}

	// 2. Second fetch (should hit the cache, not mock server)
	resp2, err := FetchLocationAreas(server.URL)
	if err != nil {
		t.Fatalf("expected no error on second fetch, got %v", err)
	}

	if resp2.Count != expectedResp.Count {
		t.Errorf("expected count %d from cached response, got %d", expectedResp.Count, resp2.Count)
	}

	if atomic.LoadInt32(&requestCount) != 1 {
		t.Errorf("expected server call count to remain 1 (cache hit), but got %d", requestCount)
	}

	// 3. Test status code error
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	_, err = FetchLocationAreas(errorServer.URL)
	if err == nil {
		t.Error("expected error for non-200 status code, got nil")
	}
}
