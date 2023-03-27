package qry

type DDLType int

const (
	DDLTypeDatabase DDLType = iota
	DDLTypeTable
)
