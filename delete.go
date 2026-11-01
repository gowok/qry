package qry

import "strings"

// DeleteQuery accumulates a SQL DELETE statement. Methods return an updated
// copy, so a query can be built up incrementally.
type DeleteQuery struct {
	command string
	table   string
	where   string
}

// Delete starts a DELETE FROM query for table.
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
	var result strings.Builder
	result.WriteString(q.command)
	result.WriteString(" " + q.table)

	if q.where != "" {
		result.WriteString(" WHERE " + q.where)
	}

	return result.String()
}
