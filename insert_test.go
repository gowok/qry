package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestInsertToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "INSERT INTO table(a, b, c) VALUES(1, ?, ?) RETURNING a"
		result := qry.Insert("table").Column("a", "b", "c").Values("1", "?", "?").Suffix("RETURNING a").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})
}
