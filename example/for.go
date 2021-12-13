package example

func makeRange(min, max int) []int {
	slice := make([]int, max-min+1)
	for index := range slice {
		slice[index] = min + 1
	}
	return slice
}

func for_caller(times int) string {
	value := ""

	for index, _ := range makeRange(0, times-1) {
		value += "go"
		_ = index
	}
	return value
}
