package main

import (
	"log"
	"net/http"
	"nls-go-messaging/internal/handlers"
	"nls-go-messaging/internal/handlers/database"
	"os"

	"github.com/gorilla/mux"
	//"nls-go-messaging/internal/handlers/database/orm"
	//"log"
	"nls-go-messaging/internal/utils"
)

func main() {
	//log.SetFlags(log.LstdFlags | log.Lshortfile) // Affiche date, heure et fichier source
	utils.InitLogger()
	database.InitDB() // ← Ajout ici

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
