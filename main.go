// @title API Messaging
// @version 1.0
// @description API de messagerie sécurisée par JWT
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" suivi d'un espace et de votre token JWT. Exemple : "Bearer {token}"
package main

import (
	"log"
	"net/http"
	"nls-go-messaging/internal/handlers"
	"nls-go-messaging/internal/handlers/database"
	"nls-go-messaging/internal/middlewares"
	"os"

	corsH "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "nls-go-messaging/docs" // Importez les docs générées par swag
	"nls-go-messaging/internal/utils"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" suivi d'un espace et de votre token JWT. Exemple : "Bearer {token}"
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

	// Configuration CORS sécurisée (adapte les origines autorisées)
	corsHandler := corsH.CORS(
		corsH.AllowedOrigins([]string{"http://localhost:3000"}), // à adapter selon ton front
		corsH.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
		corsH.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		corsH.AllowCredentials(),
	)

	// Routes API
	r.Handle("/ws/{threadId}", middlewares.JWTMiddleware(http.HandlerFunc(handlers.HandleWebSocket)))

	// (Optionnel) Route racine
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue sur l'API Messaging"))
	})

	log.Println("Serveur sur :4000")
	log.Fatal(http.ListenAndServe(":4000", corsHandler(r)))
}
