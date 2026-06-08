package model_test

import (
	"reflect"
	"testing"

	. "github.com/adomascx/Skelbiu_API/internal/model"
)

// Query

func TestQueryString(t *testing.T) {
	want := "category_id=80&cities=465%2C43&condition=2&cost_max=420&cost_min=69&keywords=test&type=1&user_type=1"
	have := QueryParams{
		Keyword:       "test",
		Cities:        "465,43",
		CategoryID:    80,
		ListingOrigin: 1,
		MinCost:       69,
		MaxCost:       420,
		IsSeller:      1,
		ConditionType: 2,
	}.String()

	if want != have {
		t.Errorf("query.String(fullQuery) - Did not return valid query:\nHave = %v\nWant =  %v", have, want)
	}

}

func BenchmarkQueryStringFull(b *testing.B) {
	query := QueryParams{
		Keyword:       "test",
		Cities:        "465,43",
		CategoryID:    80,
		ListingOrigin: 0,
		OrderBy:       0,
		MinCost:       69,
		MaxCost:       420,
		IsSeller:      1,
		ConditionType: 2,
	}

	for b.Loop() {
		_ = query.String()
	}
}

func BenchmarkQueryStringEmpty(b *testing.B) {
	query := QueryParams{}

	for b.Loop() {
		_ = query.String()
	}
}

// Listings

var testListings = Listings{
	Listing{Link: "aaa", Price: 45.6, Title: "aaaTitle"},
	Listing{Link: "bbb", Price: 12.3, Title: "bbbTitle"},
	Listing{Link: "ccc", Price: 45.6, Title: "cccTitle"},
	Listing{Link: "ddd", Price: 12.3, Title: "dddTitle"},
	Listing{Link: "aaa", Price: 45.6, Title: "aaaTitle"},
	Listing{Link: "eee", Price: 12.3, Title: "eeeTitle"},
	Listing{Link: "fff", Price: 45.6, Title: "fffTitle"},
	Listing{Link: "bbb", Price: 12.3, Title: "bbbTitle"},
}

func TestFilter(t *testing.T) {
	want := Listings{
		Listing{Link: "aaa", Price: 45.6, Title: "aaaTitle"},
		Listing{Link: "bbb", Price: 12.3, Title: "bbbTitle"},
		Listing{Link: "ccc", Price: 45.6, Title: "cccTitle"},
		Listing{Link: "ddd", Price: 12.3, Title: "dddTitle"},
		Listing{Link: "eee", Price: 12.3, Title: "eeeTitle"},
		Listing{Link: "fff", Price: 45.6, Title: "fffTitle"},
	}

	have := Listings{
		Listing{Link: "aaa", Price: 45.6, Title: "aaaTitle"},
		Listing{Link: "bbb", Price: 12.3, Title: "bbbTitle"},
		Listing{Link: "ccc", Price: 45.6, Title: "cccTitle"},
		Listing{Link: "ddd", Price: 12.3, Title: "dddTitle"},
		Listing{Link: "aaa", Price: 45.6, Title: "aaaTitle"},
		Listing{Link: "eee", Price: 12.3, Title: "eeeTitle"},
		Listing{Link: "fff", Price: 45.6, Title: "fffTitle"},
		Listing{Link: "bbb", Price: 12.3, Title: "bbbTitle"},
	}

	have.Filter()

	if !reflect.DeepEqual(have, want) {
		t.Errorf("Filter() - did not filter resuts as expected:\nHave = %v\nWant =  %v", have, want)
	}
}

func BenchmarkStructMapDedupe(b *testing.B) {

	for b.Loop() {

		len := len(testListings)
		filtered := make([]Listing, len)
		seen := make(map[Listing]bool, len)

		for _, listing := range testListings {
			if !seen[listing] {
				seen[listing] = true
				filtered = append(filtered, listing)
			}
		}
	}
}

func BenchmarkLinkMapDedupe(b *testing.B) {

	for b.Loop() {

		len := len(testListings)
		filtered := make([]Listing, len)
		seen := make(map[string]bool, len)

		for _, listing := range testListings {
			if !seen[listing.Link] {
				seen[listing.Link] = true
				filtered = append(filtered, listing)
			}
		}
	}
}
