package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestInsertToSQL(t *testing.T) {

	testCases := map[string]string{
		"INSERT INTO table(a, b, c) VALUES(1, ?, ?) RETURNING a": qry.Insert("table").Column("a", "b", "c").Values("1", "?", "?").Suffix("RETURNING a").SQL(),
		"INSERT INTO table(a, b, c) VALUES(1, 1, ?) RETURNING a": qry.Insert("table").Column("a").Values("1").Set("b", "1").Set("c", "?").Suffix("RETURNING a").SQL(),
	}

	for expected, result := range testCases {
		t.Run(expected, func(t *testing.T) {
			must := must.New(t)
			must.Equal(expected, result)
		})
	}
}
