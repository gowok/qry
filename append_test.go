package qry

import (
	"testing"

	"github.com/golang-must/must"
)

// query is implemented by every builder, so all of them can be checked the same
// way.
type query interface {
	AppendTo([]byte) []byte
	SQL() string
	appendSize() int
}

func queries() map[string]query {
	return map[string]query{
		"select": Select("a", "b").From("t").
			Join("j", "t.id = j.t_id").
			Where("a = ?").OrWhere("b = ?").
			GroupBy("a").OrderBy("a", "DESC").
			Limit(1).Offset(2).
			Prefix("WITH x AS 1").
			Suffix("FOR UPDATE"),
		"select bare":     Select(),
		"select distinct": Select("a").Distinct().From("t"),
		"select negative": Select("a").From("t").
			Limit(-1).Offset(-100),
		"insert": func() query {
			q := Insert("t")
			q.Column("a", "b")
			q.Values("1", "?")
			q.Suffix("RETURNING id")
			return q
		}(),
		"update": func() query {
			q := Update("t")
			q.Set("a", 1)
			q.Where("id = ?")
			q.Suffix("RETURNING *")
			return q
		}(),
		"delete": Delete("t").Where("a = 1"),
		"create": Create().Table("t").
			Columns("id INT PRIMARY KEY", "email TEXT NOT NULL").
			IfNotExists(),
		"create without columns": Create().Database("db").IfNotExists(),
		"drop":                   Drop().Database("db"),
		"drop table":             Drop().Table("t"),
		"alter":                  Alter().Table("t").AddColumn("address TEXT"),
		"alter rename":           Alter().Database("db").RenameTo("db2"),
	}
}

func TestAppendToMatchesSQL(t *testing.T) {
	for name, q := range queries() {
		t.Run(name, func(t *testing.T) {
			must := must.New(t)
			must.Equal(q.SQL(), string(q.AppendTo(nil)))
		})
	}
}

// TestAppendSizeIsExact pins the size hints to the rendered length: an
// undersized hint would silently cost an extra allocation in SQL.
func TestAppendSizeIsExact(t *testing.T) {
	for name, q := range queries() {
		t.Run(name, func(t *testing.T) {
			must := must.New(t)
			must.Equal(q.appendSize(), len(q.AppendTo(nil)))
		})
	}
}

func TestAppendToKeepsDestination(t *testing.T) {
	dst := []byte("<")

	must := must.New(t)
	must.Equal("<SELECT a FROM t", string(Select("a").From("t").AppendTo(dst)))
}

func TestAppendToReusesBuffer(t *testing.T) {
	buf := make([]byte, 0, 128)
	q := Select("a").From("t")

	first := string(q.AppendTo(buf))
	second := string(q.AppendTo(buf[:0]))

	must := must.New(t)
	must.Equal(first, second)
	must.Equal(128, cap(buf))
}
