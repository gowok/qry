package qry_test

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/qry"
)

func TestSelectToSQL(t *testing.T) {
	cases := map[string]string{
		"SELECT 1":                                           qry.Select("1").SQL(),
		"SELECT DISTINCT id FROM table":                      qry.Select("id").Distinct().From("table").SQL(),
		"SELECT * FROM table WHERE a = ?":                    qry.Select().From("table").AndWhere("a = ?").SQL(),
		"SELECT * FROM table WHERE b = ?":                    qry.Select().From("table").OrWhere("b = ?").SQL(),
		"SELECT * FROM table WHERE a = ? AND b = ? OR c = ?": qry.Select().From("table").Where("a = ?").AndWhere("b = ?").OrWhere("c = ?").SQL(),
		"SELECT a, b FROM table LIMIT 1 OFFSET 1":            qry.Select("a", "b").From("table").Limit(1).Offset(1).SQL(),
		"SELECT * FROM table1 t1 JOIN table2 t2 ON t1.table2_id=t2.id LEFT JOIN table3 t3 ON t1.table3_id=t3.id RIGHT JOIN table4 t4 ON t1.table4_id=t4.id INNER JOIN table5 t5 ON t1.table5_id=t5.id": qry.Select().From("table1 t1").
			Join("table2 t2", "t1.table2_id=t2.id").
			LeftJoin("table3 t3", "t1.table3_id=t3.id").
			RightJoin("table4 t4", "t1.table4_id=t4.id").
			InnerJoin("table5 t5", "t1.table5_id=t5.id").
			SQL(),
		"SELECT * FROM table ORDER BY a ASC, b DESC, c ASC": qry.Select().From("table").OrderBy("a").Desc("b").Asc("c").SQL(),
		"WITH a AS 1 SELECT a":                              qry.Select("a").Prefix("WITH a AS 1").SQL(),
		"SELECT 1 FOR UPDATE":                               qry.Select("1").Suffix("FOR UPDATE").SQL(),
		"SELECT * FROM teams GROUP BY score, groups":        qry.Select().From("teams").GroupBy("score").GroupBy("groups").SQL(),
	}

	for expected, result := range cases {
		t.Run(expected, func(t *testing.T) {
			must := must.New(t)
			must.Equal(expected, result)
		})
	}

}

func BenchmarkSelectSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qry.Select("1").SQL()
	}
}

func BenchmarkSelectComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qry.Select().From("table1 t1").
			Join("table2 t2", "t1.table2_id=t2.id").
			LeftJoin("table3 t3", "t1.table3_id=t3.id").
			RightJoin("table4 t4", "t1.table4_id=t4.id").
			InnerJoin("table5 t5", "t1.table5_id=t5.id").
			SQL()
	}
}
