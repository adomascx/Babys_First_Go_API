package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/adomascx/Skelbiu_API/internal/model"
	"github.com/adomascx/Skelbiu_API/internal/scraper"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("health request received")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetListings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("could not do r.ParseForm(%v): %v\n", r.Form, err)
		http.Error(w, "failed to parse form", http.StatusInternalServerError)
		return
	}

	log.Printf("GET listing request received")

	// parse API request
	query, pages, err := parseQuery(r.Form)
	if err != nil {
		log.Printf("could not do parseQuery(%v): %v\n", r.Form, err)
		http.Error(w, "invalid query", http.StatusBadRequest)
		return
	}
	log.Printf("parsed query: %v, %v", query, pages)

	// scrape listings with the given params
	listings, err := scraper.ScrapeListings(query, pages)
	if err != nil {
		log.Printf("could not do ScrapeListings(%v, %v): %v\n", query, pages, err)
		http.Error(w, "failed to scrape", http.StatusInternalServerError)
		return
	}

	log.Printf("Scraper found %v active listings.", len(listings))

	// If no listings are found, stop processing and return empty JSON
	if len(listings) == 0 {
		log.Printf("No listings found.")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}

	// Filter duplicates
	unfilteredLen := len(listings)
	listings.Filter()

	delta := unfilteredLen - len(listings)
	if delta != 0 {
		log.Printf("Filtered %v items, now have %v", delta, len(listings))
	}

	log.Printf("[testing] first found listing:\n%+v", listings[0])

	// Write response as encoded JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listings)

}

func parseQuery(rawQuery url.Values) (model.QueryParams, int, error) {
	query, pages := model.QueryParams{}, 0

	// parse pages
	if rawQuery.Has("Pages") {
		var err error
		pages, err = strconv.Atoi(rawQuery.Get("Pages"))
		if err != nil {
			return model.QueryParams{}, 0, fmt.Errorf("could not do type conversion for Pages: %w\n", err)
		}
	}

	// parse QueryParams using runtime reflection
	// not the fastest, but ensures changes to the underlying struct are handled properly
	queryStruct := reflect.ValueOf(&query).Elem()

	for key, valueSlice := range rawQuery {
		// Pages are parsed separately, don't attempt to parse them to struct
		if key == "Pages" {
			continue
		}

		// take only the first value of the slice, since struct contains no containers
		value := valueSlice[0]

		field := queryStruct.FieldByName(key)

		// check if field exists in struct
		if !field.IsValid() {
			return model.QueryParams{}, 0, fmt.Errorf("the field \"%v\" does not exist", key)
		}

		// check if field can be set
		if !field.CanSet() {
			return model.QueryParams{}, 0, fmt.Errorf("the field \"%v\" cannot be set (not exported/addressable)", field)
		}

		// set the value of the current field based on the field's underlying type
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)

		case reflect.Int:
			valueInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return model.QueryParams{}, 0, fmt.Errorf("could not do type conversion for value %v, field %v", value, field)
			}
			field.SetInt(valueInt)

		case reflect.Float64:
			valueFloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return model.QueryParams{}, 0, fmt.Errorf("could not do type conversion for value %v, field %v", value, field)
			}
			field.SetFloat(valueFloat)
		}

	}

	if query.CategoryID == 0 {
		return model.QueryParams{}, 0, errors.New("query must contain category ID")
	}

	return query, pages, nil
}
