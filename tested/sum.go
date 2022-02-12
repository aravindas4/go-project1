package tested

func Sum(iter []int) (sum int) {

	for _, value := range iter {
		sum += value
	}

	return
}

func SumAll(iters ...[]int) (sum []int) {
	for _, iter := range iters {
		sum = append(sum, Sum(iter))
	}
	return
}

func SumAllTails(iters ...[]int) (sum []int) {
	for _, iter := range iters {
		if len(iter) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(iter[1:]))
		}
	}
	return
}
