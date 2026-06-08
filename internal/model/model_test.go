package model_test

import (
	"testing"

	. "github.com/adomascx/Skelbiu_API/internal/model"
)

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
