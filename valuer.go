package qry

import "fmt"

// SQLString renders val as a quoted SQL literal.
func SQLString[T comparable](val T) string {
	return fmt.Sprintf("'%v'", val)
}
