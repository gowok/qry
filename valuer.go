package qry

import "fmt"

func SQLString[T comparable](val T) string {
	return fmt.Sprintf("'%v'", val)
}
