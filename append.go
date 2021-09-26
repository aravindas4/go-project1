package main

import (
	"fmt"
)

func simple_append() {
	a := []int{1, 2, 3}    // It is a slice
	b := append(a[:1], 10) //appends to existing underlying array
	fmt.Printf("a = %v, b = %v ", a, b)

	c := make([]int, len(a))
	copy(c, a) // copying the slice
	// c = append(c[:1], 4)
	c[0] = 2 // copy by value
	// c = append(c, 2)
	fmt.Printf("\na = %v, c = %v ", a, c)
}
