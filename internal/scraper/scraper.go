package scraper

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adomascx/Skelbiu_API/internal/model"
	"github.com/gocolly/colly/v2"
	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

const (
	skelbiuURL      = "https://www.skelbiu.lt/skelbimai/"
	skelbiuHost     = "www.skelbiu.lt"
	chromeUA        = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
	listingsPerPage = 24
	defaultPages    = 5
)

// If pages is set to 0, defaults to 5
func ScrapeListings(query model.QueryParams, pages int) (model.Listings, error) {
	if query == (model.QueryParams{}) {
		return nil, errors.New("no query given")
	}
	if pages < 0 {
		return nil, errors.New("pages must be >= 0")
	}

	// Default pages value and error handling
	if pages == 0 {
		pages = defaultPages
	}

	// Listings init
	listings := make([]model.Listing, 0, listingsPerPage*pages)

	// Set concurrency based on env file. May result in blocked requests.
	// Defaults to 'false'
	useConcurrency := strings.EqualFold(os.Getenv("USE_CONCURRENCY"), "true")

	// Set up colly collector
	collector := colly.NewCollector(
		colly.Async(useConcurrency),
		colly.UserAgent(chromeUA),
		colly.AllowedDomains(skelbiuHost),
	)

	delay, err := time.ParseDuration(os.Getenv("RATE_LIMIT") + "ms")
	if err != nil {
		return nil, fmt.Errorf("could not parse duration for RATE_LIMIT: %w\n", err)
	}

	err = collector.Limit(&colly.LimitRule{
		DomainGlob:  skelbiuHost,
		Parallelism: 2,
		Delay:       delay * time.Millisecond,
		RandomDelay: 500 * time.Millisecond,
	})
	if err != nil {
		return nil, err
	}

	// Create dialer raw TCP connection
	dialer := &net.Dialer{
		Timeout:   15 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	// Connect and wrap with uTLS
	utlsTransport := &http2.Transport{
		DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
			// Open TCP connection
			rawConn, err := dialer.DialContext(ctx, network, addr)
			if err != nil {
				rawConn.Close()
				return nil, fmt.Errorf("could not do dialer.DialContext(): %w\n", err)
			}

			// Configure uTLS
			utlsConfig := &utls.Config{
				ServerName: skelbiuHost,
				NextProtos: []string{"h2", "http/1.1"},
				MinVersion: utls.VersionTLS12,
			}

			// Create uTLS connection
			uconn := utls.UClient(rawConn, utlsConfig, utls.HelloChrome_Auto)

			// Perform TLS handshake
			err = uconn.HandshakeContext(ctx)
			if err != nil {
				return nil, fmt.Errorf("could not do uconn.HandshakeContext(%v): %w\n", ctx, err)
			}

			return uconn, nil
		},
	}

	// Attach uTLS connection to Colly
	utlsClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: utlsTransport,
	}

	collector.SetClient(utlsClient)

	var mutex sync.Mutex

	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "lt-LT,lt;q=0.9,en-US;q=0.8,en;q=0.7")

		log.Println("Currently scraping:", r.URL.Path)
	})

	collector.OnHTML(".standard-list-item > .extended-info", func(h *colly.HTMLElement) {
		// filter sold items
		if h.DOM.Find(".item.sold").Length() > 0 {
			return
		}

		listing := model.Listing{
			Link:        h.Request.AbsoluteURL(h.Attr("href")),
			Title:       h.ChildText(".title"),
			Description: h.ChildText(".first-dataline"),
			Date:        h.ChildText(".second-dataline"),
		}

		price, _ := parsePrice(h.ChildText(".price"))
		listing.Price = price

		mutex.Lock()
		listings = append(listings, listing)
		mutex.Unlock()
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Printf("colly encountered a problem when scraping: %v\n%+v", err, r)
	})

	for page := 1; page <= pages; page++ {

		url := skelbiuURL + strconv.Itoa(page) + "?" + query.String()
		err := collector.Visit(url)
		if err != nil {
			log.Printf("could not do Visit(%v): %v\n", url, err)
			continue
		}
	}

	collector.Wait()

	return listings, nil
}

func parsePrice(priceStr string) (float64, error) {
	var cleanedStr strings.Builder

	for _, rune := range priceStr {
		// change decimal point to Go float standard
		if rune == ',' {
			rune = '.'
		} else if rune < '0' || rune > '9' { // only accept digits 0-9
			continue
		}

		cleanedStr.WriteRune(rune)
	}

	// if price not provided, return default value of 0
	if cleanedStr.String() == "" {
		return 0, nil
	}

	output, err := strconv.ParseFloat(cleanedStr.String(), 64)

	return output, err
}
