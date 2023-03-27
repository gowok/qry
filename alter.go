package qry

import (
	"bytes"
)

type AlterQuery struct {
	command    string
	name       string
	createType DDLType
	renameTo   string
	addColumn  string
}

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

func (q AlterQuery) SQL() string {
	result := bytes.Buffer{}
	result.WriteString(q.command)
	if q.createType == DDLTypeDatabase {
		result.WriteString(" DATABASE")
	} else if q.createType == DDLTypeTable {
		result.WriteString(" TABLE")
	}

	result.WriteString(" " + q.name)

	if q.renameTo != "" {
		result.WriteString(" RENAME TO " + q.renameTo)
	}

	if q.addColumn != "" {
		result.WriteString(" ADD COLUMN " + q.addColumn)
	}

	return result.String()
}
