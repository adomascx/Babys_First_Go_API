package scraper

import (
	"fmt"
	"testing"

	"github.com/adomascx/Skelbiu_API/internal/model"
)

// Temp non-functional test
func TestGetSkelbiuHTML(t *testing.T) {
	result, err := ScrapeListings(model.QueryParams{}, 1)
	if err != nil {
		t.Errorf("could not do ScrapeListings(): %v\n", err)
	}

	fmt.Println(result)
}
