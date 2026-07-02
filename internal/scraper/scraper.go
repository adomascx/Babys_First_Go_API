package scraper

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adomascx/Skelbiu_API/internal/model"
	"github.com/gocolly/colly/v2"
	utls "github.com/refraction-networking/utls"
)

const (
	skelbiuURL      = "https://www.skelbiu.lt/skelbimai/"
	skelbiuHost     = "www.skelbiu.lt"
	chromeUA        = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
	listingsPerPage = 24
	defaultPages    = 5
)

var (
	cookieJar, _ = cookiejar.New(nil)
	// TODO: Add exported browser cookies to request
	// OTAdditionalConsentString: 1~
	// OptanonAlertBoxClosed: 2026-06-13T16:38:32.191Z
	// eupubconsent-v2: CQlvIEgQlvIEgAcABBLTCjFgAAAAAELAAChQAAAUBQDMFComhLB0kChQWAIERAgjiACAABgAAkBQAQJgQQByBAEPsIgAAAQAAgAMAABBAACAAASABCIAAACAQAiACEQABgAACAQAUCAACIixAAgABANBQCAggEgwAQIQojBAgAAAAAAACAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAIABAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAEAQhAOAAWADUAKoAbwBgAFzAMUAbQBHoCrAFigLRAXIAuwgAIAAeAGgAZYA5ACOAKTAWIBAQtAEABqAMAA
	// __cf_bm: GgnXM1q_14jnYaSP8v_quGPH
	// sessionRemID: cpb564v2gm977jul1mkllac4bu
	// PHPSESSID: cpb564v2gm977jul1mkllac4bu
	// OptanonConsent: isGpcEnabled
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

	// set the rate limit from env
	rateLimit := os.Getenv("RATE_LIMIT")

	if rateLimit == "" {
		rateLimit = "1500"
	}

	delay, err := time.ParseDuration(rateLimit + "ms")
	if err != nil {
		return nil, fmt.Errorf("could not parse duration for RATE_LIMIT: %w\n", err)
	}

	if pages > 1 {
		err = collector.Limit(&colly.LimitRule{
			DomainGlob:  skelbiuHost,
			Parallelism: 2,
			Delay:       delay,
			RandomDelay: 500 * time.Millisecond,
		})
		if err != nil {
			return nil, err
		}
	}

	// create HTTP connection via uTLS
	collector.SetClient(newUtlsClient())

	var (
		mutex     sync.Mutex
		scrapeErr error
	)

	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "lt-LT,lt;q=0.9,en-US;q=0.8,en;q=0.7")

		log.Println("Currently scraping:", r.URL.Path)
	})

	collector.OnHTML(".standard-list-item", func(h *colly.HTMLElement) {
		// filter sold items
		if h.DOM.Find(".item.sold").Length() > 0 {
			return
		}

		listing := model.Listing{
			Link:        h.Request.AbsoluteURL(h.Attr("href")),
			Title:       h.ChildTexts(".title")[0],
			Description: h.ChildText(".first-dataline"),
			Date:        h.ChildText(".second-dataline"),
		}

		priceStr := h.ChildTexts(".price")
		price := 0.0

		// filter listings with no price tag (usually needed when lister type is buyer)
		if len(priceStr) != 0 {
			price, err = parsePrice(priceStr[0])
			if err != nil {
				fmt.Printf("could not do parsePrice(%v): %v\n", priceStr[0], err)
				return
			}
		}

		listing.Price = price

		mutex.Lock()
		listings = append(listings, listing)
		mutex.Unlock()
	})

	collector.OnError(func(r *colly.Response, err error) {

		mutex.Lock()
		scrapeErr = fmt.Errorf("scraping failed: status = %v, url = %v, err = %w\n", r.StatusCode, r.Request.URL.String(), err)
		mutex.Unlock()

		// log extra info to help fix 403 errors
		log.Printf("colly encountered a problem when scraping: status = %d, url = %s, err = %v",
			r.StatusCode,
			r.Request.URL.String(),
			err,
		)

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

	if scrapeErr != nil {
		return listings, scrapeErr
	}

	return listings, nil
}

func newUtlsClient() *http.Client {
	// Create TCP connection dialer
	dialer := &net.Dialer{
		Timeout:   15 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	// Connect and wrap with uTLS
	utlsTransport := &http.Transport{
		DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// Open TCP connection
			rawConn, err := dialer.DialContext(ctx, network, addr)
			if err != nil {
				return nil, fmt.Errorf("could not dial TCP: %w\n", err)
			}

			// Config uTLS
			utlsConfig := &utls.Config{
				ServerName: skelbiuHost,
				NextProtos: []string{"http/1.1"},
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
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: utlsTransport,
		Jar:       cookieJar,
	}
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
