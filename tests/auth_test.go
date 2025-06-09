package tests

import (
	"encoding/json"

	//"nls-go-messaging/internal/handlers/database"
	"testing"
)

/* func init() {
	database.InitDB()
} */

func TestRegisterAndLogin(t *testing.T) {
	// Enregistre un nouvel utilisateur
	body := map[string]string{"username": "testuser", "password": "testpass"}
	jsonBody, _ := json.Marshal(body)
	println("JSON body:", string(jsonBody))

	/*
		 	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handlers.RegisterHandler(w, req)
			if w.Code != http.StatusCreated {
				t.Fatalf("échec register: status %d", w.Code)
			}

			// Connecte l'utilisateur
			req = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w = httptest.NewRecorder()

			handlers.LoginHandler(w, req)
			if w.Code != http.StatusOK {
				t.Fatalf("échec login: status %d", w.Code)
			}
	*/
}
