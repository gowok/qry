package qry

type DDLType int

const (
	DDLTypeDatabase DDLType = iota + 1
	DDLTypeTable
)
