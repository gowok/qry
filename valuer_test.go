package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestSQLString(t *testing.T) {
	testCases := map[string][]string{
		"string": {
			"'alex'",
			qry.SQLString("alex"),
		},
		"int": {
			"'10'",
			qry.SQLString(10),
		},
		"float": {
			"'12.45'",
			qry.SQLString(12.45),
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			must := must.New(t)
			must.Equal(v[0], v[1])
		})
	}

}
