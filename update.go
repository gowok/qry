package qry

import "fmt"

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

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity.
func (q UpdateQuery) AppendTo(dst []byte) []byte {
	dst = append(dst, q.command...)
	dst = append(dst, ' ')
	dst = append(dst, q.table...)

	if len(q.sets) > 0 {
		dst = append(dst, " SET "...)
		dst = appendJoin(dst, ", ", q.sets)
	}

	if q.where != "" {
		dst = append(dst, " WHERE "...)
		dst = append(dst, q.where...)
	}

	if q.suffix != "" {
		dst = append(dst, ' ')
		dst = append(dst, q.suffix...)
	}

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q UpdateQuery) appendSize() int {
	n := len(q.command) + len(q.table) + 1

	if len(q.sets) > 0 {
		n += len(" SET ") + joinSize(", ", q.sets)
	}

	if q.where != "" {
		n += len(" WHERE ") + len(q.where)
	}

	if q.suffix != "" {
		n += len(q.suffix) + 1
	}

	return n
}

func (q UpdateQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
