package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"nls-go-messaging/internal/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur connexion DB: ", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.ThreadModel{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
}
