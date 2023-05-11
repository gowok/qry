package qry

import (
	"bytes"
)

type CreateQuery struct {
	command     string
	name        string
	createType  DDLType
	columns     []string
	ifNotExists bool
}

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

func (q CreateQuery) SQL() string {
	result := bytes.Buffer{}
	result.WriteString(q.command)
	if q.createType == DDLTypeDatabase {
		result.WriteString(" DATABASE")
	} else if q.createType == DDLTypeTable {
		result.WriteString(" TABLE")
	}

	if q.ifNotExists {
		result.WriteString(" IF NOT EXISTS")
	}

	result.WriteString(" " + q.name)

	if len(q.columns) > 0 {
		result.WriteString(" (")
		for i, v := range q.columns {
			if i > 0 {
				result.WriteString(", ")
			}
			result.WriteString(v)
		}
		result.WriteString(")")
	}

	return result.String()
}
