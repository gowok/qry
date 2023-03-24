package qry

import (
	"bytes"
)

type DeleteQuery struct {
	command string
	table   string
	where   string
}

func Delete(table string) DeleteQuery {
	q := DeleteQuery{
		command: "DELETE FROM",
		table:   table,
	}

	return q
}

func (q DeleteQuery) Where(cond string) DeleteQuery {
	q.where = cond
	return q
}

func (q DeleteQuery) SQL() string {
	result := bytes.Buffer{}
	result.WriteString(q.command)
	result.WriteString(" " + q.table)

	if q.where != "" {
		result.WriteString(" WHERE " + q.where)
	}

	return result.String()
}
