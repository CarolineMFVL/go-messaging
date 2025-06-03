package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedDB() {
	users := []User{
		{Username: "alice", Password: "password123"},
		{Username: "bob", Password: "secure456"},
	}

	for _, u := range users {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Erreur hash %s: %v", u.Username, err)
			continue
		}
		u.Password = string(hashed)
		result := DB.Create(&u)
		if result.Error != nil {
			log.Printf("Erreur insertion %s: %v", u.Username, result.Error)
		} else {
			log.Printf("Utilisateur seedé : %s", u.Username)
		}
	}
}
