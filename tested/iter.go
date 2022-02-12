package tested

func Repeat(value string, times int) (result string) {
	for index := 0; index < times; index++ {
		result += value
	}

	return
}
