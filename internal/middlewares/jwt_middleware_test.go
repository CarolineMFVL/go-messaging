package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func generateTestJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "testuser",
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func TestJWTMiddleware_ValidToken(t *testing.T) {
	token := generateTestJWT()
	req := httptest.NewRequest("GET", "/ws/123", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rr := httptest.NewRecorder()

	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	handler := JWTMiddleware(next)
	handler.ServeHTTP(rr, req)

	if !called {
		t.Error("This handler was not called with a valid token")
	}
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestJWTMiddleware_MissingToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/ws/123", nil)
	rr := httptest.NewRecorder()

	handler := JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Missing or invalid Authorization header") {
		t.Error("Expected error message not found")
	}
}
