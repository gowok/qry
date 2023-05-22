package qry

import "fmt"

func Avg(column string) string {
	return fmt.Sprintf("AVG(%s)", column)
}

func Count(column string) string {
	return fmt.Sprintf("COUNT(%s)", column)
}

func Max(column string) string {
	return fmt.Sprintf("MAX(%s)", column)
}

func Min(column string) string {
	return fmt.Sprintf("MIN(%s)", column)
}

func Sum(column string) string {
	return fmt.Sprintf("SUM(%s)", column)
}
