package model

type Listing struct {
	ID          int    `json:"id"`          // Unique ID of this listing
	Title       string `json:"title"`       // listing's title text
	Description string `json:"description"` // truncated text description of the listing
	PriceCents  int    `json:"price_cents"` // store price as int, where 1 = 0,01€
	Date        string `json:"date"`        // date since last update to listing, in plaintext
}

type ListingOrigin int

const (
	ListerEither ListingOrigin = iota
	ListerPrivate
	ListerOrganization
)

type OrderBy int

const (
	OrderUpdated OrderBy = iota
	OrderNewest
	OrderBest
	OrderCheapest
	OrderMostExpensive OrderBy = 8
)

type IsSeller int

const (
	SellerOrBuyer IsSeller = iota
	Seller
	Buyer
)

type ConditionType int

const (
	ConditionEither ConditionType = iota
	ConditionNew
	ConditionUsed
)

type QueryParams struct {
	Keyword       string        // presumably search bar keywords; URL = "keywords"
	Cities        []int         // city IDs (internally, comma-separated internal city values, e.g. Vilnius and Kaunas = "465,43"); URL = "cities"
	CategoryID    int           // specific category ID; URL = "category_id"
	ListingOrigin ListingOrigin // 0 = either, 1 = private, 2 = organization; URL = "user_type"
	OrderBy       OrderBy       // 0 = atnaujinti, 1 = naujausi, 2 = tinkamiausi, 3 = pigiausi, 8 = brangiausi; URL = "orderBy"
	MinCostCents  int           // minimum price as int, where 1 = 0,01€; URL = "cost_min"
	MaxCostCents  int           // maximum price as int, where 1 = 0,01€; URL = "cost_max"
	IsSeller      IsSeller      // 0 = either, 1 = seller, 2 = buyer; URL = "type"
	ConditionType ConditionType // 0 = either, 1 = new, 2 = used; URL = "condition"
}
