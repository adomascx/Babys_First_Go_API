package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r)
		log.Println("GET request received")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Serving traffic on port ", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
