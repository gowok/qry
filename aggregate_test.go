package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestAggregateSimple(t *testing.T) {
	cases := map[string]string{
		"AVG(id)":   qry.Avg("id"),
		"COUNT(id)": qry.Count("id"),
		"MAX(id)":   qry.Max("id"),
		"MIN(id)":   qry.Min("id"),
		"SUM(id)":   qry.Sum("id"),
	}

	for expected, result := range cases {
		t.Run(expected, func(t *testing.T) {
			must := must.New(t)
			must.Equal(expected, result)
		})
	}
}

