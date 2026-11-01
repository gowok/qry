package qry

// AlterQuery accumulates a SQL ALTER statement. Methods return an updated copy,
// so a query can be built up incrementally.
type AlterQuery struct {
	command    string
	name       string
	createType DDLType
	renameTo   string
	addColumn  string
}

// Alter starts an ALTER statement.
func Alter() AlterQuery {
	q := AlterQuery{
		command: "ALTER",
	}

	return q
}

func (q AlterQuery) Database(name string) AlterQuery {
	q.createType = DDLTypeDatabase
	q.name = name
	return q
}

func (q AlterQuery) RenameTo(name string) AlterQuery {
	q.renameTo = name
	return q
}

func (q AlterQuery) AddColumn(column string) AlterQuery {
	q.addColumn = column
	return q
}

func (q AlterQuery) Table(name string) AlterQuery {
	q.createType = DDLTypeTable
	q.name = name
	return q
}

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity.
func (q AlterQuery) AppendTo(dst []byte) []byte {
	dst = append(dst, q.command...)
	switch q.createType {
	case DDLTypeDatabase:
		dst = append(dst, " DATABASE"...)
	case DDLTypeTable:
		dst = append(dst, " TABLE"...)
	}

	dst = append(dst, ' ')
	dst = append(dst, q.name...)

	if q.renameTo != "" {
		dst = append(dst, " RENAME TO "...)
		dst = append(dst, q.renameTo...)
	}

	if q.addColumn != "" {
		dst = append(dst, " ADD COLUMN "...)
		dst = append(dst, q.addColumn...)
	}

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q AlterQuery) appendSize() int {
	n := len(q.command) + len(q.name) + 1

	switch q.createType {
	case DDLTypeDatabase:
		n += len(" DATABASE")
	case DDLTypeTable:
		n += len(" TABLE")
	}

	if q.renameTo != "" {
		n += len(" RENAME TO ") + len(q.renameTo)
	}

	if q.addColumn != "" {
		n += len(" ADD COLUMN ") + len(q.addColumn)
	}

	return n
}

func (q AlterQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
