package qry

import "strings"

// InsertQuery accumulates a SQL INSERT statement. Unlike the other builders it
// is mutated in place.
type InsertQuery struct {
	command string
	table   string
	columns string
	values  string
	suffix  string
}

// Insert starts an INSERT INTO query for table.
func Insert(table string) *InsertQuery {
	q := InsertQuery{
		command: "INSERT INTO",
		table:   table,
	}

	return &q
}

func (q *InsertQuery) Column(cols ...string) {
	q.columns = strings.Join(cols, ", ")
}

func (q *InsertQuery) Values(vals ...string) {
	q.values = strings.Join(vals, ", ")
}

func (q *InsertQuery) Suffix(suffix string) {
	q.suffix = suffix
}

func (q *InsertQuery) Set(col string, val string) {
	if q.columns != "" {
		q.columns += ", "
	}
	if q.values != "" {
		q.values += ", "
	}
	q.columns += col
	q.values += val
}

func (q InsertQuery) SQL() string {
	var result strings.Builder
	result.WriteString(q.command)
	result.WriteString(" " + q.table)
	result.WriteString("(" + q.columns + ")")
	result.WriteString(" VALUES")
	result.WriteString("(" + q.values + ")")

	if q.suffix != "" {
		result.WriteString(" " + q.suffix)
	}
	return result.String()
}
