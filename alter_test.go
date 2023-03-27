package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestAlterToSQL(t *testing.T) {

	t.Run("database", func(t *testing.T) {
		expected := "ALTER DATABASE company RENAME TO mycomp"
		result := Alter().Database("company").RenameTo("mycomp").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("table", func(t *testing.T) {
		expected := "ALTER TABLE user RENAME TO users"
		result := Alter().Table("user").RenameTo("users").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("table add column", func(t *testing.T) {
		expected := "ALTER TABLE users ADD COLUMN address TEXT"
		result := Alter().Table("users").AddColumn("address TEXT").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
