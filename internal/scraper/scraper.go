package scraper

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adomascx/Skelbiu_API/internal/model"
	"github.com/gocolly/colly"
)

const SKELBIU_URL = "https://www.skelbiu.lt/skelbimai/"

// If pages is set to 0, defaults to 5
func ScrapeListings(query model.QueryParams, pages int) ([]model.Listing, error) {
	if query == (model.QueryParams{}) {
		return nil, errors.New("No query given")
	}
	if pages < 0 {
		return nil, errors.New("Pages must be >= 0")
	}

	// Default pages value and error handling
	if pages == 0 {
		pages = 5
	}

	// Listings init
	listings := make([]model.Listing, 24*pages)

	// Set concurrency based on env file.
	// May result in blocked requests.
	useConcurrency := strings.EqualFold(os.Getenv("USE_CONCURRENCY"), "true")
	collector := colly.NewCollector(colly.Async(useConcurrency))

	collector.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       1 * time.Second,
	})

	var mutex sync.Mutex

	for page := 1; page <= pages; page++ {

		collector.OnHTML(".standard-list-item", func(h *colly.HTMLElement) {
			listing := model.Listing{
				Link:        h.Attr("href"),
				Title:       h.ChildText(".title"),
				Description: h.ChildText(".first-dataline"),
				Date:        h.ChildText(".second-dataline"),
			}

			Price, _ := strconv.ParseFloat(h.ChildText(".price"), 64)
			listing.Price = Price

			mutex.Lock()
			listings = append(listings, listing)
			mutex.Unlock()
		})

		url := SKELBIU_URL + strconv.Itoa(page) + "?" + query.String()
		collector.Visit(url)

		collector.Wait()
	}

	return listings, nil
}
