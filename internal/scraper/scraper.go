package scraper

import (
	"errors"
	"strconv"

	"github.com/adomascx/Skelbiu_API/internal/model"
	"github.com/gocolly/colly"
)

func trimTrailingSlash(path string) string {
	lastIdx := len(path) - 1
	if path[lastIdx] == '/' {
		return path[:lastIdx]
	}
	return path
}

func ScrapeListings(baseURL string, query model.QueryParams, pages int) ([]model.Listing, error) {
	// Handle incorrect user input
	if len(baseURL) == 0 {
		return nil, errors.New("Base URL is empty")
	}
	if query == (model.QueryParams{}) {
		return nil, errors.New("No query given")
	}
	if pages <= 0 {
		return nil, errors.New("Pages is less than or equal to 0")
	}

	// Listings init
	listings := make([]model.Listing, 24*pages)

	collector := colly.NewCollector()

	collector.OnHTML(".standard-list-item", func(h *colly.HTMLElement) {
		listing := model.Listing{
			Link:        h.Attr("href"),
			Title:       h.ChildText(".title"),
			Description: h.ChildText(".first-dataline"),
			Date:        h.ChildText(".second-dataline"),
		}

		Price, _ := strconv.ParseFloat(h.ChildText(".price"), 64)
		listing.Price = Price

		listings = append(listings, listing)
	})

	url := trimTrailingSlash(baseURL) + "/" + strconv.Itoa(pages) + "?" + query.Encode()
	collector.Visit(url)

	return listings, nil
}
