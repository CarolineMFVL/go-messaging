package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	RespondWithError(rr, http.StatusBadRequest, "Error test")

	if rr.Code != http.StatusBadRequest {
		t.Errorf("status code wanted %d, but got %d", http.StatusBadRequest, rr.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("JSON decoding error: %v", err)
	}
	if resp["error"] != "Error test" {
		t.Errorf("expected 'Error test', got '%s'", resp["error"])
	}
}
