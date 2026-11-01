package qry

import (
	"strconv"
	"strings"
)

// SelectQuery accumulates a SQL SELECT statement. Methods return an updated
// copy, so a query can be built up incrementally.
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
	group    []string

	prefix string
	suffix string
}

// pushString returns dst with value appended without sharing its backing array,
// so queries copied by value never mutate each other.
func pushString(dst []string, value string) []string {
	out := make([]string, len(dst), len(dst)+1)
	copy(out, dst)
	return append(out, value)
}

// Select starts a SELECT query. When no columns are given, "*" is used.
func Select(columns ...string) SelectQuery {
	q := SelectQuery{
		command: "SELECT",
	}

	if len(columns) == 0 {
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
	if q.where == "" {
		q.where = cond
	} else {
		q.where += " AND " + cond
	}
	return q
}

func (q SelectQuery) OrWhere(cond string) SelectQuery {
	if q.where == "" {
		q.where = cond
	} else {
		q.where += " OR " + cond
	}
	return q
}

func (q SelectQuery) AndWhere(cond string) SelectQuery {
	return q.Where(cond)
}

func (q SelectQuery) GroupBy(column string) SelectQuery {
	q.group = pushString(q.group, column)
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
	q.joins = pushString(q.joins, "JOIN "+table+" ON "+cond)
	return q
}

func (q SelectQuery) LeftJoin(table string, cond string) SelectQuery {
	q.joins = pushString(q.joins, "LEFT JOIN "+table+" ON "+cond)
	return q
}

func (q SelectQuery) RightJoin(table string, cond string) SelectQuery {
	q.joins = pushString(q.joins, "RIGHT JOIN "+table+" ON "+cond)
	return q
}

func (q SelectQuery) InnerJoin(table string, cond string) SelectQuery {
	q.joins = pushString(q.joins, "INNER JOIN "+table+" ON "+cond)
	return q
}

func (q SelectQuery) OrderBy(column string, dir ...string) SelectQuery {
	if len(dir) <= 0 {
		dir = []string{"ASC"}
	}

	q.orders = pushString(q.orders, column+" "+dir[0])
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

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity, which makes it
// the allocation-free way to render a query.
func (q SelectQuery) AppendTo(dst []byte) []byte {
	if q.prefix != "" {
		dst = append(dst, q.prefix...)
		dst = append(dst, ' ')
	}

	dst = append(dst, q.command...)
	if q.distinct {
		dst = append(dst, " DISTINCT"...)
	}

	dst = append(dst, ' ')
	dst = append(dst, q.columns...)

	if q.table != "" {
		dst = append(dst, " FROM "...)
		dst = append(dst, q.table...)
	}

	for _, v := range q.joins {
		dst = append(dst, ' ')
		dst = append(dst, v...)
	}

	if q.where != "" {
		dst = append(dst, " WHERE "...)
		dst = append(dst, q.where...)
	}

	if len(q.group) > 0 {
		dst = append(dst, " GROUP BY "...)
		dst = appendJoin(dst, ", ", q.group)
	}

	if len(q.orders) > 0 {
		dst = append(dst, " ORDER BY "...)
		dst = appendJoin(dst, ", ", q.orders)
	}

	if q.limit != 0 {
		dst = append(dst, " LIMIT "...)
		dst = strconv.AppendInt(dst, int64(q.limit), 10)
	}

	if q.offset != 0 {
		dst = append(dst, " OFFSET "...)
		dst = strconv.AppendInt(dst, int64(q.offset), 10)
	}

	if q.suffix != "" {
		dst = append(dst, ' ')
		dst = append(dst, q.suffix...)
	}

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q SelectQuery) appendSize() int {
	n := len(q.command) + len(q.columns) + 1

	if q.prefix != "" {
		n += len(q.prefix) + 1
	}

	if q.distinct {
		n += len(" DISTINCT")
	}

	if q.table != "" {
		n += len(" FROM ") + len(q.table)
	}

	for _, v := range q.joins {
		n += len(v) + 1
	}

	if q.where != "" {
		n += len(" WHERE ") + len(q.where)
	}

	if len(q.group) > 0 {
		n += len(" GROUP BY ") + joinSize(", ", q.group)
	}

	if len(q.orders) > 0 {
		n += len(" ORDER BY ") + joinSize(", ", q.orders)
	}

	if q.limit != 0 {
		n += len(" LIMIT ") + intSize(q.limit)
	}

	if q.offset != 0 {
		n += len(" OFFSET ") + intSize(q.offset)
	}

	if q.suffix != "" {
		n += len(q.suffix) + 1
	}

	return n
}

func (q SelectQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
