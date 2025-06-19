package main

import (
	"log"
	"net/http"
	"nls-go-messaging/internal/handlers"
	"nls-go-messaging/internal/handlers/database"
	"os"

	_ "nls-go-messaging/docs" // Importez les docs générées par swag
	"nls-go-messaging/internal/utils"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	utils.InitLogger()
	database.InitDB()

	if os.Getenv("SEED_DB") == "1" {
		//SeedDB()
		log.Println("Base de données seedée")
		return
	}

	// Création du routeur principal
	r := mux.NewRouter()

	// Documentation Swagger accessible sur /swagger/index.html
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Routes API
	r.HandleFunc("/ws/{threadId}", handlers.HandleWebSocket)

	// (Optionnel) Route racine
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue sur l'API Messaging"))
	})

	log.Println("Serveur sur :4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
