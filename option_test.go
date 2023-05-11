package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestOption(t *testing.T) {

	t.Run("WithDriver", func(t *testing.T) {
		result := qry.NewOption()
		expected := "mysql"

		qry.WithDriver(expected)(result)
		must := must.New(t)
		must.Equal(expected, result.Driver)
	})

}
