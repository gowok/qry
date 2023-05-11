package qry

import (
	"testing"

	"github.com/golang-must/must"
)

func TestSelectToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "SELECT 1"
		result := Select("1").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("prepared", func(t *testing.T) {
		expected := "SELECT * FROM table WHERE a = 1 AND b = ?"
		result := Select().From("table").Where("a = 1 AND b = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("limit offset", func(t *testing.T) {
		expected := "SELECT a, b FROM table LIMIT 1 OFFSET 1"
		result := Select("a", "b").From("table").Limit(1).Offset(1).SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("join", func(t *testing.T) {
		expected := "SELECT * FROM table1 t1 JOIN table2 t2 ON t1.table2_id=t2.id JOIN table3 t3 ON t1.table3_id=t3.id"
		result := Select().From("table1 t1").
			Join("table2 t2", "t1.table2_id=t2.id").
			Join("table3 t3", "t1.table3_id=t3.id").
			SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("order", func(t *testing.T) {
		expected := "SELECT * FROM table ORDER BY a ASC, b DESC, c ASC"
		result := Select().From("table").OrderBy("a").Desc("b").Asc("c").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
