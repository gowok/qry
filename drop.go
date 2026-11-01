package qry

// DropQuery accumulates a SQL DROP statement. Methods return an updated copy,
// so a query can be built up incrementally.
type DropQuery struct {
	command    string
	name       string
	createType DDLType
}

// Drop starts a DROP statement.
func Drop() DropQuery {
	q := DropQuery{
		command: "DROP",
	}

	return q
}

func (q DropQuery) Database(name string) DropQuery {
	q.createType = DDLTypeDatabase
	q.name = name
	return q
}

func (q DropQuery) Table(name string) DropQuery {
	q.createType = DDLTypeTable
	q.name = name
	return q
}

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity.
func (q DropQuery) AppendTo(dst []byte) []byte {
	dst = append(dst, q.command...)
	switch q.createType {
	case DDLTypeDatabase:
		dst = append(dst, " DATABASE"...)
	case DDLTypeTable:
		dst = append(dst, " TABLE"...)
	}

	dst = append(dst, ' ')
	dst = append(dst, q.name...)

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q DropQuery) appendSize() int {
	n := len(q.command) + len(q.name) + 1

	switch q.createType {
	case DDLTypeDatabase:
		n += len(" DATABASE")
	case DDLTypeTable:
		n += len(" TABLE")
	}

	return n
}

func (q DropQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
