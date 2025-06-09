package database

import (
	"fmt"
	"log"
	"nls-go-messaging/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	TicketID  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(100)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varchar(100)"`
	Topic     string `gorm:"type:varchar(100)"`
	Watermark string `gorm:"type:varchar(100)"`
}

func InitDB( /* dialect, host, port, dbname, pass string */ ) (*gorm.DB, error) {
	var DB *gorm.DB
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur connexion PostgreSQL: ", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.ThreadModel{}, &database.Document{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
	//db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dialect, host, port, dbname, pass))
	return DB, err
	//defer db.Close()
}
