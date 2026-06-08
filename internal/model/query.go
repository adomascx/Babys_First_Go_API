package model

import "github.com/google/go-querystring/query"

type ListingOrigin int

const (
	ListerEither ListingOrigin = iota
	ListerPrivate
	ListerOrganization
)

type OrderBy int

const (
	Updated OrderBy = iota
	Newest
	Best
	Cheapest
	MostExpensive OrderBy = 8
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
	Keyword       string        `url:"keywords,omitempty"`    // presumably search bar keywords
	Cities        string        `url:"cities,omitempty"`      // comma-separated city IDs, e.g. Vilnius and Kaunas = "465,43"
	CategoryID    int           `url:"category_id,omitempty"` // specific category ID
	ListingOrigin ListingOrigin `url:"user_type,omitempty"`   // ListerEither = 0, ListerPrivate = 1, ListerOrganization = 2
	OrderBy       OrderBy       `url:"orderBy,omitempty"`     // Updated = 0, Newest = 1, Best = 2, Cheapest = 3, MostExpensive = 8
	MinCost       float64       `url:"cost_min,omitempty"`    // minimum price in €
	MaxCost       float64       `url:"cost_max,omitempty"`    // maximum price in €
	IsSeller      IsSeller      `url:"type,omitempty"`        // SellerOrBuyer = 0, Seller = 1, Buyer = 2
	ConditionType ConditionType `url:"condition,omitempty"`   // ConditionEither = 0, ConditionNew = 1, ConditionUsed = 2
}

func (q QueryParams) String() string {
	values, _ := query.Values(q)

	return values.Encode()
}
