package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adomascx/Skelbiu_API/internal/handler"
	"github.com/joho/godotenv"
)

const DEFAULT_PORT = "8080"

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Print("Couldn't load .env file, using default values")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.GetHealth)
	mux.HandleFunc("GET /listings", handler.GetListings)

	// HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	log.Println("Listening on port", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
