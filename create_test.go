package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestCreateToSQL(t *testing.T) {

	t.Run("database", func(t *testing.T) {
		expected := "CREATE DATABASE company"
		result := Create().Database("company").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("table", func(t *testing.T) {
		expected := "CREATE TABLE users"
		result := Create().Table("users").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
