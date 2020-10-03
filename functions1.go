package main

import (
	"fmt"
)

func calculateBill(price int, num int) (int, int) {
	return price * num, price - num
}

func func1(value1 int, value2 int) (sum int, product int) {
	sum = value1 + value2
	product = value2 * value1
	return
}

func main() {
	fmt.Println(calculateBill(30, 6))

	total, diff := calculateBill(14, 5) 
	fmt.Println(total, diff)

	fmt.Println(func1(3,4))

	sum, _ := func1(6, 5)
	fmt.Println(sum)
}