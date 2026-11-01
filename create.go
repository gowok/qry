package qry

// CreateQuery accumulates a SQL CREATE statement. Methods return an updated
// copy, so a query can be built up incrementally.
type CreateQuery struct {
	command     string
	name        string
	createType  DDLType
	columns     []string
	ifNotExists bool
}

// Create starts a CREATE statement.
func Create() CreateQuery {
	q := CreateQuery{
		command: "CREATE",
	}

	return q
}

func (q CreateQuery) Database(name string) CreateQuery {
	q.createType = DDLTypeDatabase
	q.name = name
	return q
}

func (q CreateQuery) Table(name string) CreateQuery {
	q.createType = DDLTypeTable
	q.name = name
	return q
}

func (q CreateQuery) Columns(columns ...string) CreateQuery {
	q.columns = columns
	return q
}

func (q CreateQuery) IfNotExists() CreateQuery {
	q.ifNotExists = true
	return q
}

// AppendTo appends the statement to dst and returns the extended slice. It
// does not allocate as long as dst has enough spare capacity.
func (q CreateQuery) AppendTo(dst []byte) []byte {
	dst = append(dst, q.command...)
	switch q.createType {
	case DDLTypeDatabase:
		dst = append(dst, " DATABASE"...)
	case DDLTypeTable:
		dst = append(dst, " TABLE"...)
	}

	if q.ifNotExists {
		dst = append(dst, " IF NOT EXISTS"...)
	}

	dst = append(dst, ' ')
	dst = append(dst, q.name...)

	if len(q.columns) > 0 {
		dst = append(dst, " ("...)
		dst = appendJoin(dst, ", ", q.columns)
		dst = append(dst, ')')
	}

	return dst
}

// appendSize reports the exact number of bytes AppendTo writes.
func (q CreateQuery) appendSize() int {
	n := len(q.command) + len(q.name) + 1

	switch q.createType {
	case DDLTypeDatabase:
		n += len(" DATABASE")
	case DDLTypeTable:
		n += len(" TABLE")
	}

	if q.ifNotExists {
		n += len(" IF NOT EXISTS")
	}

	if len(q.columns) > 0 {
		n += len(" (") + joinSize(", ", q.columns) + 1
	}

	return n
}

func (q CreateQuery) SQL() string {
	return string(q.AppendTo(make([]byte, 0, q.appendSize())))
}
