package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestSelectToSQL(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := "SELECT 1"
		result := qry.Select("1").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("distinct", func(t *testing.T) {
		expected := "SELECT DISTINCT id FROM table"
		result := qry.Select("id").Distinct().From("table").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("prepared", func(t *testing.T) {
		expected := "SELECT * FROM table WHERE a = ?"
		result := qry.Select().From("table").Where("a = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("where", func(t *testing.T) {
		expected := "SELECT * FROM table WHERE a = ? AND b = ?"
		result := qry.Select().From("table").OrWhere("a = ?").AndWhere("b = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("where 2", func(t *testing.T) {
		expected := "SELECT * FROM table WHERE a = ? OR b = ?"
		result := qry.Select().From("table").AndWhere("a = ?").OrWhere("b = ?").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("limit offset", func(t *testing.T) {
		expected := "SELECT a, b FROM table LIMIT 1 OFFSET 1"
		result := qry.Select("a", "b").From("table").Limit(1).Offset(1).SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("join", func(t *testing.T) {
		expected := "SELECT * FROM table1 t1 JOIN table2 t2 ON t1.table2_id=t2.id LEFT JOIN table3 t3 ON t1.table3_id=t3.id RIGHT JOIN table4 t4 ON t1.table4_id=t4.id INNER JOIN table5 t5 ON t1.table5_id=t5.id"
		result := qry.Select().From("table1 t1").
			Join("table2 t2", "t1.table2_id=t2.id").
			LeftJoin("table3 t3", "t1.table3_id=t3.id").
			RightJoin("table4 t4", "t1.table4_id=t4.id").
			InnerJoin("table5 t5", "t1.table5_id=t5.id").
			SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("order", func(t *testing.T) {
		expected := "SELECT * FROM table ORDER BY a ASC, b DESC, c ASC"
		result := qry.Select().From("table").OrderBy("a").Desc("b").Asc("c").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("prefix", func(t *testing.T) {
		expected := "WITH a AS 1 SELECT a"
		result := qry.Select("a").Prefix("WITH a AS 1").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

	t.Run("suffix", func(t *testing.T) {
		expected := "SELECT 1 FOR UPDATE"
		result := qry.Select("1").Suffix("FOR UPDATE").SQL()

		must := must.New(t)
		must.Equal(expected, result)
	})

}
