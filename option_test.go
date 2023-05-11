package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestOption(t *testing.T) {

	t.Run("WithDriver", func(t *testing.T) {
		result := &option{}
		expected := "mysql"

		WithDriver(expected)(result)
		must := must.New(t)
		must.Equal(expected, result.driver)
	})

}
