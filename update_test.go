package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

var testUpdate = map[string][]string{
	"minimum": {
		"UPDATE table SET a = 1, b = 2",
		qry.Update("table").Set("a", 1).Set("b", 2).SQL(),
	},
	"where": {
		"UPDATE table WHERE id = 1",
		qry.Update("table").Where("id = 1").SQL(),
	},
	"sufix": {
		"UPDATE table RETURNING *",
		qry.Update("table").Suffix("RETURNING *").SQL(),
	},
}

func TestUpdateToSQL(t *testing.T) {
	for name, test := range testUpdate {
		t.Run(name, func(t *testing.T) {
			must := must.New(t)
			must.Equal(test[0], test[1])
		})
	}
}
