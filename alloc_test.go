package qry_test

import (
	"testing"

	"github.com/gowok/qry"
)

// Benchmarks report allocations per operation so the renderers can be measured.
//
//	...BuildAndRender — build the query and render it, the typical use case
//	...Render         — render with SQL(), which returns a new string
//	...AppendTo       — render into a caller-owned buffer
//
// Run them with:
//
//	go test -run=none -bench=. -benchmem

// renderBufferSize is enough for every statement built in this file, so the
// AppendTo benchmarks measure rendering alone and never growth.
const renderBufferSize = 512

// renderSink keeps rendered output reachable, so the compiler cannot delete the
// work the benchmarks are meant to measure.
var renderSink string

func BenchmarkSelectBuildAndRender(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		renderSink = qry.Select().From("table1 t1").
			Join("table2 t2", "t1.table2_id=t2.id").
			Where("t1.deleted_at IS NULL").
			GroupBy("t1.id").
			OrderBy("t1.id", "DESC").
			Limit(10).
			Offset(20).
			SQL()
	}
}

func BenchmarkSelectRender(b *testing.B) {
	q := qry.Select().From("table1 t1").
		Join("table2 t2", "t1.table2_id=t2.id").
		Where("t1.deleted_at IS NULL").
		GroupBy("t1.id").
		OrderBy("t1.id", "DESC").
		Limit(10).
		Offset(20)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkSelectAppendTo(b *testing.B) {
	q := qry.Select().From("table1 t1").
		Join("table2 t2", "t1.table2_id=t2.id").
		Where("t1.deleted_at IS NULL").
		GroupBy("t1.id").
		OrderBy("t1.id", "DESC").
		Limit(10).
		Offset(20)

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkInsertBuildAndRender(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q := qry.Insert("table")
		q.Column("a", "b", "c")
		q.Values("1", "2", "3")
		q.Suffix("RETURNING a")
		renderSink = q.SQL()
	}
}

func BenchmarkInsertRender(b *testing.B) {
	q := qry.Insert("table")
	q.Column("a", "b", "c")
	q.Values("1", "2", "3")
	q.Suffix("RETURNING a")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkInsertAppendTo(b *testing.B) {
	q := qry.Insert("table")
	q.Column("a", "b", "c")
	q.Values("1", "2", "3")
	q.Suffix("RETURNING a")

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkUpdateBuildAndRender(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q := qry.Update("table")
		q.Set("a", 1)
		q.Set("b", 2)
		q.Where("id = 1")
		q.Suffix("RETURNING *")
		renderSink = q.SQL()
	}
}

func BenchmarkUpdateRender(b *testing.B) {
	q := qry.Update("table")
	q.Set("a", 1)
	q.Set("b", 2)
	q.Where("id = 1")
	q.Suffix("RETURNING *")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkUpdateAppendTo(b *testing.B) {
	q := qry.Update("table")
	q.Set("a", 1)
	q.Set("b", 2)
	q.Where("id = 1")
	q.Suffix("RETURNING *")

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkDeleteRender(b *testing.B) {
	q := qry.Delete("table").Where("a = 1 AND b = ?")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkDeleteAppendTo(b *testing.B) {
	q := qry.Delete("table").Where("a = 1 AND b = ?")

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkCreateRender(b *testing.B) {
	q := qry.Create().
		Table("users").
		Columns("id INT PRIMARY KEY", "email TEXT NOT NULL").
		IfNotExists()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkCreateAppendTo(b *testing.B) {
	q := qry.Create().
		Table("users").
		Columns("id INT PRIMARY KEY", "email TEXT NOT NULL").
		IfNotExists()

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkDropRender(b *testing.B) {
	q := qry.Drop().Table("users")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkDropAppendTo(b *testing.B) {
	q := qry.Drop().Table("users")

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkAlterRender(b *testing.B) {
	q := qry.Alter().Table("users").AddColumn("address TEXT")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = q.SQL()
	}
}

func BenchmarkAlterAppendTo(b *testing.B) {
	q := qry.Alter().Table("users").AddColumn("address TEXT")

	buf := make([]byte, 0, renderBufferSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = q.AppendTo(buf[:0])
	}
}

func BenchmarkSQLString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = qry.SQLString("alex")
	}
}

func BenchmarkSQLStringNonString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderSink = qry.SQLString(12.45)
	}
}
