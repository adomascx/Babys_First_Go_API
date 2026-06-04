package model

type Listing struct {
	Link        string  `json:"link"`        // link to the listing page
	Title       string  `json:"title"`       // listing's title text
	Description string  `json:"description"` // truncated text description of the listing
	Date        string  `json:"date"`        // date since last update to listing, in plaintext
	Price       float64 `json:"price"`       // price in €
}
