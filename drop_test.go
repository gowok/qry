package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestDropToSQL(t *testing.T) {

	t.Run("database", func(t *testing.T) {
		expected := "DROP DATABASE company"
		result := Drop().Database("company").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("table", func(t *testing.T) {
		expected := "DROP TABLE users"
		result := Drop().Table("users").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
