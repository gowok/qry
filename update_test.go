package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestUpdateToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "UPDATE table SET a = 1, b = 2 WHERE c = ?"
		result := Update("table").Set("a", 1).Set("b", 2).Where("c = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})
}
