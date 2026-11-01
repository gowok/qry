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

func TestSelectCopiesAreIndependent(t *testing.T) {
	base := qry.Select().From("table").
		Join("j1", "c1").
		Join("j2", "c2").
		Join("j3", "c3").
		GroupBy("g1").
		GroupBy("g2").
		GroupBy("g3").
		OrderBy("o1").
		OrderBy("o2").
		OrderBy("o3").
		Where("a = ?")

	first := base.AndWhere("b = ?").Join("j4", "c4").GroupBy("g4").OrderBy("o4")
	second := base.OrWhere("x = ?").Join("j5", "c5").GroupBy("g5").OrderBy("o5")

	cases := map[string]string{
		"base":   "SELECT * FROM table JOIN j1 ON c1 JOIN j2 ON c2 JOIN j3 ON c3 WHERE a = ? GROUP BY g1, g2, g3 ORDER BY o1 ASC, o2 ASC, o3 ASC",
		"first":  "SELECT * FROM table JOIN j1 ON c1 JOIN j2 ON c2 JOIN j3 ON c3 JOIN j4 ON c4 WHERE a = ? AND b = ? GROUP BY g1, g2, g3, g4 ORDER BY o1 ASC, o2 ASC, o3 ASC, o4 ASC",
		"second": "SELECT * FROM table JOIN j1 ON c1 JOIN j2 ON c2 JOIN j3 ON c3 JOIN j5 ON c5 WHERE a = ? OR x = ? GROUP BY g1, g2, g3, g5 ORDER BY o1 ASC, o2 ASC, o3 ASC, o5 ASC",
	}

	results := map[string]string{
		"base":   base.SQL(),
		"first":  first.SQL(),
		"second": second.SQL(),
	}

	for name, expected := range cases {
		t.Run(name, func(t *testing.T) {
			must := must.New(t)
			must.Equal(expected, results[name])
		})
	}
}

func BenchmarkSelectSimple(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qry.Select("1").SQL()
	}
}

func BenchmarkSelectComplex(b *testing.B) {
	b.ReportAllocs()
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
