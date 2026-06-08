package handler

import (
	"net/url"
	"testing"

	"github.com/adomascx/Skelbiu_API/internal/model"
)

type output struct {
	query model.QueryParams
	page  int
	err   string
}

func TestParseQuery(t *testing.T) {
	testCases := [...]struct {
		description string
		inputQuery  url.Values
		want        output
	}{
		{
			description: "empty input",
			inputQuery:  url.Values{},
			want:        output{model.QueryParams{}, 0, ""},
		},
		{
			description: "2 simple params",
			inputQuery: url.Values{
				"IsSeller": []string{"2"},
				"Pages":    []string{"3"},
			},
			want: output{model.QueryParams{IsSeller: model.Buyer}, 3, ""},
		},
		{
			description: "all types",
			inputQuery: url.Values{
				"Cities":   []string{"465,43"},
				"MinCost":  []string{"12.34"},
				"IsSeller": []string{"2"},
			},
			want: output{model.QueryParams{Cities: "465,43", MinCost: 12.34, IsSeller: model.Buyer}, 0, ""},
		},
		{
			description: "incorrect query struct key",
			inputQuery: url.Values{
				"Cityz": []string{"465,43"},
			},
			want: output{model.QueryParams{}, 0, "The field \"Cityz\" does not exist"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			query, page, err := parseQuery(tC.inputQuery)

			errString := ""
			if err != nil {
				errString = err.Error()
			}

			have := output{
				query: query,
				page:  page,
				err:   errString,
			}

			if tC.want.err != "" {
				if have.err != tC.want.err {
					t.Errorf("parseQuery(%v) - Did not return correct error:\nHave = %v\nWant =  %v", tC.inputQuery, have.err, tC.want.err)
				}
				return
			}

			if have.err != "" {
				t.Errorf("parseQuery(%v) - Returned unexpected error:\n%v", tC.inputQuery, have.err)
			}

			if have != tC.want {
				t.Errorf("parseQuery(%v) - Did not return correct result:\nHave = %v\nWant =  %v", tC.inputQuery, have, tC.want)
			}
		})
	}
}
