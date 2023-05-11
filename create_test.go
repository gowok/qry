package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestCreateToSQL(t *testing.T) {

	t.Run("database", func(t *testing.T) {
		expected := "CREATE DATABASE IF NOT EXISTS company"
		result := Create().Database("company").IfNotExists().SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("table", func(t *testing.T) {
		expected := "CREATE TABLE users (id INT PRIMARY KEY, email TEXT NOT NULL)"
		result := Create().
			Table("users").
			Columns(
				"id INT PRIMARY KEY",
				"email TEXT NOT NULL",
			).
			SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
