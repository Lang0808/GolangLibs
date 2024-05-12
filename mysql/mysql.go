package mysql

import "fmt"

func FormatArrayMysql[K int | int32 | int64 | string | float32 | float64](arr []K) string {
	ans := "("
	first := true
	for _, element := range arr {
		if !first {
			ans = ans + ", "
		}
		ans = ans + fmt.Sprint(element)
		first = false
	}

	ans = ans + ")"
	return ans
}

func IsEmptyRowError(err error) bool {
	return err.Error() == "sql: no rows in result set"
}
