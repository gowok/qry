package qry

import (
	"bytes"
	"strings"
)

type InsertQuery struct {
	command string
	table   string
	columns string
	values  string
	suffix  string
}

func Insert(table string) InsertQuery {
	q := InsertQuery{
		command: "INSERT INTO",
		table:   table,
	}

	return q
}

func (q InsertQuery) Column(cols ...string) InsertQuery {
	q.columns = strings.Join(cols, ", ")
	return q
}

func (q InsertQuery) Values(vals ...string) InsertQuery {
	q.values = strings.Join(vals, ", ")
	return q
}

func (q InsertQuery) Suffix(suffix string) InsertQuery {
	q.suffix = suffix
	return q
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
