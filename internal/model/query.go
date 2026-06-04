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
	Keyword       string        `url:"keywords,omitempty"`    // presumably search bar keywords
	Cities        string        `url:"cities,omitempty"`      // comma-separated city IDs, e.g. Vilnius and Kaunas = "465,43"
	CategoryID    int           `url:"category_id,omitempty"` // specific category ID
	ListingOrigin ListingOrigin `url:"user_type,omitempty"`   // 0 = either, 1 = private, 2 = organization
	OrderBy       OrderBy       `url:"orderBy,omitempty"`     // 0 = atnaujinti, 1 = naujausi, 2 = tinkamiausi, 3 = pigiausi, 8 = brangiausi
	MinCostCents  int           `url:"cost_min,omitempty"`    // minimum price as int, where 1 = 0,01€
	MaxCostCents  int           `url:"cost_max,omitempty"`    // maximum price as int, where 1 = 0,01€
	IsSeller      IsSeller      `url:"type,omitempty"`        // 0 = either, 1 = seller, 2 = buyer
	ConditionType ConditionType `url:"condition,omitempty"`   // 0 = either, 1 = new, 2 = used
}

func (q QueryParams) Encode() string {
	values, _ := query.Values(q)

	return values.Encode()
}

func (q QueryParams) Decode(string) {

}
