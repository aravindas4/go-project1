package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := true
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	//
	var a1 int = 2224444444222
	b1 := 95
	fmt.Println("value of a is", a1, "and b is", b1)
	fmt.Printf("type of a is %T, size of a is %d\n", a1, unsafe.Sizeof(a1))
	fmt.Printf("type of b is %T, size of b is %d\n", b1, unsafe.Sizeof(b1))

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("prodduct:", cmul)

	i := 2
	j := 4.5

	fmt.Println(i + int(j))
	fmt.Println(float64(i) + j)
}