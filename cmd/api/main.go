package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adomascx/Skelbiu_API/internal/handler"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.GetHealth)
	mux.HandleFunc("GET /listings", handler.GetListings)

	// HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Serving traffic on port ", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
