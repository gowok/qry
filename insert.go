package qry

import (
	"bytes"
	"fmt"
	"strings"
)

type InsertQuery struct {
	command string
	table   string
	columns string
	values  string
	suffix  string
}

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
	q.columns += fmt.Sprintf("%s", col)
	q.values += fmt.Sprintf("%s", val)
}

func (q InsertQuery) SQL() string {
	result := bytes.Buffer{}
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
