package qry

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

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity.
func (q DeleteQuery) AppendTo(dst []byte) []byte {
	dst = append(dst, q.command...)
	dst = append(dst, ' ')
	dst = append(dst, q.table...)

	if q.where != "" {
		dst = append(dst, " WHERE "...)
		dst = append(dst, q.where...)
	}

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q DeleteQuery) appendSize() int {
	n := len(q.command) + len(q.table) + 1

	if q.where != "" {
		n += len(" WHERE ") + len(q.where)
	}

	return n
}

func (q DeleteQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
