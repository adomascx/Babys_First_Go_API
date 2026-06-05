package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/adomascx/Skelbiu_API/internal/model"
)

var (
	listings []model.Listing
	mutex    sync.Mutex
)

// Handlers
func GetHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("health request received")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetListings(w http.ResponseWriter, r *http.Request) {
	log.Println("GET listing request received")

	// scraper.ScrapeListings()

	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
}
