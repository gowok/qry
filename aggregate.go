package qry

// Avg returns the SQL AVG aggregate of column.
func Avg(column string) string {
	return "AVG(" + column + ")"
}

// Count returns the SQL COUNT aggregate of column.
func Count(column string) string {
	return "COUNT(" + column + ")"
}

// Max returns the SQL MAX aggregate of column.
func Max(column string) string {
	return "MAX(" + column + ")"
}

// Min returns the SQL MIN aggregate of column.
func Min(column string) string {
	return "MIN(" + column + ")"
}

// Sum returns the SQL SUM aggregate of column.
func Sum(column string) string {
	return "SUM(" + column + ")"
}
