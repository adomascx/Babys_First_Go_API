package scraper

import (
	"os"
	"testing"

	"github.com/adomascx/Skelbiu_API/internal/model"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestEnvLoad(t *testing.T) {
	env := os.Getenv("USE_CONCURRENCY")
	if env != "true" && env != "false" {
		t.Errorf("env loading - variable not present or empty:\nHave = %v\nWant =  %v", os.Getenv("USE_CONCURRENCY"), "true or false")
	}
}

// Non-standard test, only checks for error
// The only logical way to implement actual tests would be to mimic Skelbiu.lt's HTTP endpoint, which is out of scope for this project
func TestScrapeListings(t *testing.T) {
	_, err := ScrapeListings(model.QueryParams{CategoryID: 80}, 1)
	if err != nil {
		t.Errorf("could not do ScrapeListings(): %v\n", err)
	}
}

func TestParsePrice(t *testing.T) {
	testCases := []struct {
		want  float64
		input string
	}{
		{
			want:  1234.5,
			input: "1 234,5€",
		},
		{
			want:  123456789.0,
			input: "kaina =  1 2 3 4 5 6 7 8 9 , 0 eurai/€",
		},
		{
			want:  0,
			input: "kaina sutartine",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			have, err := parsePrice(tC.input)
			if err != nil {
				t.Errorf("parsePrice() - returned error:\nHave = %v\nWant =  %v\nError =  %v", have, tC.want, err)
			}
			if tC.want != have {
				t.Errorf("parsePrice() - did not parse price as expected:\nHave = %v\nWant =  %v", have, tC.want)
			}
		})
	}
}
