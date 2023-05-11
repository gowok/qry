package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestDeleteToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "DELETE FROM table WHERE a = 1 AND b = ?"
		result := qry.Delete("table").Where("a = 1 AND b = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})
}
