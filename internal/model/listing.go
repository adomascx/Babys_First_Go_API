package model

type Listing struct {
	ID          int    `json:"id"`          // Unique ID of this listing
	Title       string `json:"title"`       // listing's title text
	Description string `json:"description"` // truncated text description of the listing
	PriceCents  int    `json:"price_cents"` // store price as int, where 1 = 0,01€
	Date        string `json:"date"`        // date since last update to listing, in plaintext
}
