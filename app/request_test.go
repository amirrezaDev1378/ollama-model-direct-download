package app

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func TestMakeRequestToRegistry(t *testing.T) {
	testURL := "https://" + DefaultRegistry
	parsedURL, err := url.Parse(testURL)
	if err != nil {
		t.Fatalf("failed to parse URL: %v", err)
	}
	resp, err := makeRequest(context.Background(), http.MethodGet, parsedURL, nil, nil, nil)
	if err != nil {
		t.Fatalf("makeRequest() error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200 OK, got %d", resp.StatusCode)
	}

}
