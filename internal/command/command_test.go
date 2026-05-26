package command

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ShishirSimha/pokedex-go/internal/util"
)

func TestCommandHelp(t *testing.T) {
	InitCommands()
	cfg := &Config{}
	err := commandHelp(cfg)
	if err != nil {
		t.Errorf("expected no error for help command, got %v", err)
	}
}

func TestCommandExit(t *testing.T) {
	InitCommands()
	cfg := &Config{}
	err := commandExit(cfg)
	if err == nil || err.Error() != "exit" {
		t.Errorf("expected error 'exit', got %v", err)
	}
}

func TestCommandMapAndMapBack(t *testing.T) {
	InitCommands()

	// Mock response from API
	nextURL := "http://mock-next-url"
	prevURL := "http://mock-prev-url"
	mockResponse := util.LocationAreaResponse{
		Count:    40,
		Next:     &nextURL,
		Previous: &prevURL,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "area-1", URL: "url-1"},
			{Name: "area-2", URL: "url-2"},
		},
	}

	// Create test server to mock the API calls
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	cfg := &Config{}

	// Test Map command by setting cfg.Next to server.URL to hit mock server
	cfg.Next = &server.URL

	err := commandMap(cfg)
	if err != nil {
		t.Fatalf("commandMap returned error: %v", err)
	}

	// Verify config is updated
	if cfg.Next == nil || *cfg.Next != nextURL {
		t.Errorf("expected Next to be %s, got %v", nextURL, cfg.Next)
	}
	if cfg.Previous == nil || *cfg.Previous != prevURL {
		t.Errorf("expected Previous to be %s, got %v", prevURL, cfg.Previous)
	}

	// Test MapBack command by setting cfg.Previous to server.URL
	cfg.Previous = &server.URL

	err = commandMapBack(cfg)
	if err != nil {
		t.Fatalf("commandMapBack returned error: %v", err)
	}

	// Verify config is updated again
	if cfg.Next == nil || *cfg.Next != nextURL {
		t.Errorf("expected Next to be %s, got %v", nextURL, cfg.Next)
	}
	if cfg.Previous == nil || *cfg.Previous != prevURL {
		t.Errorf("expected Previous to be %s, got %v", prevURL, cfg.Previous)
	}

	// Test MapBack first page validation
	firstPageCfg := &Config{Previous: nil}
	err = commandMapBack(firstPageCfg)
	if err != nil {
		t.Errorf("expected no error for first page validation, got %v", err)
	}
}
