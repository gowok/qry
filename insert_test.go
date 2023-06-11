package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestInsertToSQL(t *testing.T) {

	testCases := map[string]string{
		"INSERT INTO table(a, b, c) VALUES(1, ?, ?) RETURNING a": func() string {
			q := qry.Insert("table")
			q.Column("a", "b", "c")
			q.Values("1", "?", "?")
			q.Suffix("RETURNING a")
			return q.SQL()
		}(),
		"INSERT INTO table(a, b, c) VALUES(1, 1, ?) RETURNING a": func() string {
			q := qry.Insert("table")
			q.Column("a")
			q.Values("1")
			q.Set("b", "1")
			q.Set("c", "?")
			q.Suffix("RETURNING a")
			return q.SQL()
		}(),
	}

	for expected, result := range testCases {
		t.Run(expected, func(t *testing.T) {
			must := must.New(t)
			must.Equal(expected, result)
		})
	}
}
