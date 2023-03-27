package qry

import (
	"bytes"
)

type DropQuery struct {
	command    string
	name       string
	createType DDLType
	renameTo   string
	addColumn  string
}

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

func (q DropQuery) SQL() string {
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
