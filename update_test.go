package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

var testUpdate = map[string][]string{
	"minimum": {
		"UPDATE table SET a = 1, b = 2",
		func() string {
			q := qry.Update("table")
			q.Set("a", 1)
			q.Set("b", 2)
			return q.SQL()
		}(),
	},
	"where": {
		"UPDATE table WHERE id = 1",
		func() string {
			q := qry.Update("table")
			q.Where("id = 1")
			return q.SQL()
		}(),
	},
	"sufix": {
		"UPDATE table RETURNING *",
		func() string {
			q := qry.Update("table")
			q.Suffix("RETURNING *")
			return q.SQL()
		}(),
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
