package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	// adapte ce nom selon ton module
)

func init() {
	&InitDB()
}

func TestRegisterAndLogin(t *testing.T) {
	// Enregistre un nouvel utilisateur
	body := map[string]string{"username": "testuser", "password": "testpass"}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	messaging_backend.RegisterHandler(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("échec register: status %d", w.Code)
	}

	// Connecte l'utilisateur
	req = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	messaging_backend.LoginHandler(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("échec login: status %d", w.Code)
	}
}
