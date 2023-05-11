package qry

import (
	"bytes"
	"fmt"
	"strings"
)

type SelectQuery struct {
	command string
	columns string
	table   string
	where   string
	limit   int
	offset  int

	distinct bool
	joins    []string
	orders   []string

	prefix string
	suffix string
}

func Select(columns ...string) SelectQuery {
	q := SelectQuery{
		command: "SELECT",
		joins:   make([]string, 0),
		orders:  make([]string, 0),
	}

	if len(columns) <= 0 {
		q.columns = "*"
	} else {
		q.columns = strings.Join(columns, ", ")
	}

	return q
}

func (q SelectQuery) From(table string) SelectQuery {
	q.table = table
	return q
}

func (q SelectQuery) Where(cond string) SelectQuery {
	q.where = cond
	return q
}

func (q SelectQuery) Distinct() SelectQuery {
	q.distinct = true
	return q
}

func (q SelectQuery) Limit(lim int) SelectQuery {
	q.limit = lim
	return q
}

func (q SelectQuery) Offset(offset int) SelectQuery {
	q.offset = offset
	return q
}

func (q SelectQuery) Join(table string, cond string) SelectQuery {
	q.joins = append(q.joins, fmt.Sprintf("%s ON %s", table, cond))
	return q
}

func (q SelectQuery) OrderBy(column string, dir ...string) SelectQuery {
	if len(dir) <= 0 {
		dir = []string{"ASC"}
	}

	q.orders = append(q.orders, fmt.Sprintf("%s %s", column, dir[0]))
	return q
}

func (q SelectQuery) Asc(column string) SelectQuery {
	return q.OrderBy(column, "ASC")
}

func (q SelectQuery) Desc(column string) SelectQuery {
	return q.OrderBy(column, "DESC")
}

func (q SelectQuery) Prefix(prefix string) SelectQuery {
	q.prefix = prefix
	return q
}

func (q SelectQuery) Suffix(suffix string) SelectQuery {
	q.suffix = suffix
	return q
}

func (q SelectQuery) SQL() string {
	result := bytes.Buffer{}

	if q.prefix != "" {
		result.WriteString(q.prefix + " ")
	}

	result.WriteString(q.command)
	if q.distinct {
		result.WriteString(" DISTINCT")
	}

	result.WriteString(" " + q.columns)

	if q.table != "" {
		result.WriteString(" FROM " + q.table)
	}

	for _, v := range q.joins {
		result.WriteString(" JOIN " + v)
	}

	if q.where != "" {
		result.WriteString(" WHERE " + q.where)
	}

	if len(q.orders) > 0 {
		result.WriteString(" ORDER BY ")
		result.WriteString(strings.Join(q.orders, ", "))
	}

	if q.limit != 0 {
		result.WriteString(fmt.Sprintf(" LIMIT %d", q.limit))
	}

	if q.offset != 0 {
		result.WriteString(fmt.Sprintf(" OFFSET %d", q.offset))
	}

	if q.suffix != "" {
		result.WriteString(" " + q.suffix)
	}

	return result.String()
}
