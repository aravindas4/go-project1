package main

import (
	"fmt"
	"math"
)

func main() {
	var age = 23
	fmt.Println("My age is", age)
	age = 29
	fmt.Println("My age is", age)
	age = 54
	fmt.Println("My age is", age)

	var width, height = 100, 50
	fmt.Println("width is", width, "and height is", height)

	var (
		name    = "Aravinda"
		age1    = 27
		height1 int
	)
	fmt.Println("my name is", name, ", age is", age1, "and height is", height1)

	count := 10
	fmt.Println("Count =", count)

	name, age = "Aravinda1", 29
	fmt.Println("my name is", name, "age is", age)

	a, b := 3, 4
	fmt.Println(a + b)
	b, c := 4, 5
	fmt.Println(c + b)
	b, c = 3, 5
	fmt.Println(c + b)

	aq1, bq1 := 145.2, 534.4
	cq1 := math.Min(aq1, bq1)
	fmt.Println(cq1)
}
