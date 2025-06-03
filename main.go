package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"nls-go-messaging/internal/handlers"
)

func main() {
	InitDB() // ← Ajout ici

	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/ws/{threadId}", handlers.HandleWebSocket)

	log.Println("Serveur sur :4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
