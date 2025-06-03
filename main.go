package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"nls-go-messaging/internal/handlers"
)

func main() {
	InitDB() // ← Ajout ici

	if os.Getenv("SEED_DB") == "1" {
		SeedDB()
		log.Println("Base de données seedée")
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/ws/{threadId}", handlers.HandleWebSocket)

	log.Println("Serveur sur :4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

InitDB() {
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

	err = DB.AutoMigrate(&models.User{}, &models.ThreadModel{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
}

