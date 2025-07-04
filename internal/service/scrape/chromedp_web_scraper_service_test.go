package scrape

import (
	"testing"
	"time"
)

func TestChromeDPWebScraperService_Close(t *testing.T) {
	service := newChromeDPWebScraperService()
	chromeDPService, ok := service.(*ChromeDPWebScraperService)
	if !ok {
		t.Fatal("Expected ChromeDPWebScraperService")
	}

	// Test that cancel function exists
	if chromeDPService.cancel == nil {
		t.Error("Cancel function should not be nil")
	}

	// Test cleanup
	err := service.Close()
	if err != nil {
		t.Errorf("Close() returned error: %v", err)
	}

	// Multiple calls to Close should be safe
	err = service.Close()
	if err != nil {
		t.Errorf("Second Close() call returned error: %v", err)
	}
}

func TestChromeDPWebScraperService_GetHTML_Timeout(t *testing.T) {
	service := newChromeDPWebScraperService()
	defer service.Close()

	// Test with a URL that should respond quickly
	// Using a local test or mock would be better, but this tests the timeout mechanism
	start := time.Now()
	_, err := service.GetHTML("https://httpbin.org/delay/1")
	duration := time.Since(start)

	// Should complete within reasonable time (much less than 30s timeout)
	if duration > 10*time.Second {
		t.Errorf("Request took too long: %v", duration)
	}

	// Error handling test (optional, depends on network)
	if err != nil {
		t.Logf("Request failed (may be network-related): %v", err)
	}
}