package database

import (
	"fmt"
	"log"
	"nls-go-messaging/internal/models"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	_ = godotenv.Load()
}

type Document struct {
	gorm.Model
	TicketID  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(100)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varchar(100)"`
	Topic     string `gorm:"type:varchar(100)"`
	Watermark string `gorm:"type:varchar(100)"`
}

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")

	// Connexion temporaire à la base 'postgres'
	connStr := os.Getenv(("CONNEXION_STRING"))

	// postgresDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=nls_db port=%s sslmode=disable", host, user, password, port)
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erreur connexion PostgreSQL (postgres): ", err)
	}
	defer sqlDB.Close()

	// Vérifie si la base existe, sinon la crée
	var exists bool
	checkQuery := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbname)
	err = sqlDB.QueryRow(checkQuery).Scan(&exists)
	if err == sql.ErrNoRows {
		_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			log.Fatal("Erreur création DB: ", err)
		}
	} else if err != nil && err != sql.ErrNoRows {
		log.Fatal("Erreur vérification DB: ", err)
	}

	// Connexion GORM normale
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur connexion PostgreSQL: ", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.ThreadModel{}, Document{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
	return db, err
}

type PG_DB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *PG_DB {
	return &PG_DB{DB: db}
}

/*
func InitDB( /* dialect, host, port, dbname, pass string  ) (*gorm.DB, error) {
	//var DB *gorm.DB
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")

	/*
	   	 dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	   		os.Getenv("DB_HOST"),
	   		os.Getenv("DB_USER"),
	   		os.Getenv("DB_PASS"),
	   		os.Getenv("DB_NAME"),
	   		os.Getenv("DB_PORT"),

	   )

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
	// dsn := "host=localhost user=testuser password=testpass dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatal("Erreur connexion PostgreSQL: ", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.ThreadModel{}, Document{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
	//db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dialect, host, port, dbname, pass))
	return db, err
	//defer db.Close()
}
*/
