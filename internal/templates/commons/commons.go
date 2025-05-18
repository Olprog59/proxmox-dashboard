package template_commons

import "fmt"

func GetOnline[N int | int8 | int64](online N, data *float64) string {
	if online > 0 {
		return fmt.Sprintf("relative overflow-hidden bg-transparent", data)
	} else {
		return "bg-gray-200 text-gray-600"
	}
}
