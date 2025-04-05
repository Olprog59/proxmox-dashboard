package templates

func getOnline[N int | int8 | int64](online N) string {
	if online > 0 {
		return "bg-green-100 text-green-600"
	} else {
		return "bg-gray-200 text-gray-600"
	}
}
