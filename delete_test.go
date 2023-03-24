package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestDeleteToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "DELETE FROM table WHERE a = 1 AND b = ?"
		result := Delete("table").Where("a = 1 AND b = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})
}
