package scraper

import (
	"fmt"
	"testing"

	"github.com/adomascx/Skelbiu_API/internal/model"
)

// Temp non-functional test
func TestGetSkelbiuHTML(t *testing.T) {
	result, err := ScrapeListings("https://en.wikipedia.org/wiki/List_of_lists_of_lists", model.QueryParams{}, 1)
	if err != nil {
		t.Errorf("could not do ScrapeListings(): %v\n", err)
	}

	fmt.Println(result)
}

func TestTrimTrailingSlash(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "en.wikipedia.org/",
			want:  "en.wikipedia.org",
		},
		{
			input: "skelbiu.lt",
			want:  "skelbiu.lt",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			have := trimTrailingSlash(tC.input)
			if have != tC.want {
				t.Errorf("got %s, want %s", have, tC.want)
			}
		})
	}
}
