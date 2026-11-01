// Package qry builds SQL statements — SELECT, INSERT, UPDATE, DELETE, CREATE,
// DROP, and ALTER — as plain strings.
//
// The builders only generate SQL text; they never execute statements or bind
// parameters.
package qry

// DDLType identifies the schema object a DDL statement targets.
type DDLType int

const (
	// DDLTypeDatabase targets a database.
	DDLTypeDatabase DDLType = iota + 1
	// DDLTypeTable targets a table.
	DDLTypeTable
)
