package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestInsertToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "INSERT INTO table(a, b, c) VALUES(1, ?, ?) RETURNING a"
		result := Insert("table").Column("a", "b", "c").Values("1", "?", "?").Suffix("RETURNING a").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})
}
