package qry

import (
	"bytes"
	"fmt"
	"strings"
)

type UpdateQuery struct {
	command string
	table   string
	sets    []string
	values  string
	where   string
	suffix  string
}

func Update(table string) UpdateQuery {
	q := UpdateQuery{
		command: "UPDATE",
		table:   table,
	}

	return q
}

func (q UpdateQuery) Set(col string, val any) UpdateQuery {
	q.sets = append(q.sets, fmt.Sprintf("%s = %v", col, val))
	return q
}

func (q UpdateQuery) Where(cond string) UpdateQuery {
	q.where = cond
	return q
}

func (q UpdateQuery) Suffix(suffix string) UpdateQuery {
	q.suffix = suffix
	return q
}

func (q UpdateQuery) SQL() string {
	result := bytes.Buffer{}
	result.WriteString(q.command)
	result.WriteString(" " + q.table)

	if len(q.sets) > 0 {
		result.WriteString(" SET ")
		result.WriteString(strings.Join(q.sets, ", "))
	}

	if q.where != "" {
		result.WriteString(" WHERE " + q.where)
	}

	if q.suffix != "" {
		result.WriteString(" " + q.suffix)
	}
	return result.String()
}
