package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	users []User
	mutex sync.Mutex
)

// Return Types
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

// Handlers
func GetHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("health request received")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET user request received")
	w.WriteHeader(http.StatusOK) // maybe a different status should be used?
	w.Write([]byte{})            // temporarily
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("POST user request received")

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	user.ID = len(users) + 1
	users = append(users, user)
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", GetHealth)

	mux.HandleFunc("GET /user", GetUser)
	mux.HandleFunc("POST /user", CreateUser)

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
