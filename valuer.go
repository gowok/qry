package qry

import "fmt"

// SQLString renders val as a quoted SQL literal.
func SQLString[T comparable](val T) string {
	// Strings are the common case; quoting one directly avoids fmt entirely.
	if s, ok := any(val).(string); ok {
		return "'" + s + "'"
	}

	return fmt.Sprintf("'%v'", val)
}
