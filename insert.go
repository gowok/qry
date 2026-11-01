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

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity.
func (q InsertQuery) AppendTo(dst []byte) []byte {
	dst = append(dst, q.command...)
	dst = append(dst, ' ')
	dst = append(dst, q.table...)
	dst = append(dst, '(')
	dst = append(dst, q.columns...)
	dst = append(dst, ')')
	dst = append(dst, " VALUES"...)
	dst = append(dst, '(')
	dst = append(dst, q.values...)
	dst = append(dst, ')')

	if q.suffix != "" {
		dst = append(dst, ' ')
		dst = append(dst, q.suffix...)
	}

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q InsertQuery) appendSize() int {
	// Two paren pairs: one around the columns, one around the values.
	const parens = 4

	n := len(q.command) + len(" ") + len(q.table) + parens + len(q.columns) +
		len(" VALUES") + len(q.values)

	if q.suffix != "" {
		n += len(" ") + len(q.suffix)
	}

	return n
}

func (q InsertQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
