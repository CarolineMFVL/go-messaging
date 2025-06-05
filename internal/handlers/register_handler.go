package handlers

import (
	"encoding/json"
	"net/http"
	"nls-go-messaging/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	var DB gorm.DB
	json.NewDecoder(r.Body).Decode(&creds)

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username et mot de passe requis", http.StatusBadRequest)
		return
	}

	// Vérifier si l'utilisateur existe déjà
	var existing models.User
	result := DB.Where("username = ?", creds.Username).First(&existing)
	if result.Error == nil {
		http.Error(w, "Nom d'utilisateur déjà utilisé", http.StatusConflict)
		return
	}

	// Hacher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur de hachage", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: creds.Username,
		Password: string(hashedPassword),
	}

	DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur créé"})
}
