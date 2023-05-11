package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestDropToSQL(t *testing.T) {

	t.Run("database", func(t *testing.T) {
		expected := "DROP DATABASE company"
		result := qry.Drop().Database("company").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("table", func(t *testing.T) {
		expected := "DROP TABLE users"
		result := qry.Drop().Table("users").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
