package qry

import (
	"fmt"
	"strings"
)

// UpdateQuery accumulates a SQL UPDATE statement. Unlike the other builders it
// is mutated in place.
type UpdateQuery struct {
	command string
	table   string
	sets    []string
	where   string
	suffix  string
}

// Update starts an UPDATE query for table.
func Update(table string) *UpdateQuery {
	q := UpdateQuery{
		command: "UPDATE",
		table:   table,
	}

	return &q
}

func (q *UpdateQuery) Set(col string, val any) {
	q.sets = append(q.sets, fmt.Sprintf("%s = %v", col, val))
}

func (q *UpdateQuery) Where(cond string) {
	q.where = cond
}

func (q *UpdateQuery) Suffix(suffix string) {
	q.suffix = suffix
}

func (q UpdateQuery) SQL() string {
	var result strings.Builder
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
