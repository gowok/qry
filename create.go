package qry

import (
	"bytes"
)

type CreateQuery struct {
	command    string
	name       string
	createType DDLType
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

func (q CreateQuery) SQL() string {
	result := bytes.Buffer{}
	result.WriteString(q.command)
	if q.createType == DDLTypeDatabase {
		result.WriteString(" DATABASE")
	} else if q.createType == DDLTypeTable {
		result.WriteString(" TABLE")
	}

	result.WriteString(" " + q.name)

	return result.String()
}
