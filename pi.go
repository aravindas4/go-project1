package main

import "fmt"

func pi() {
	var π = 22 / 7.0
	// This do not compile
	// var a, b = 22, 7.0
	// pi := a / b
	fmt.Println(π)
	const a, b = 22, 7.0
	pi := a / b
	fmt.Println(pi)
	const myUint8 = 1000000000000000000000000000000000000000000000
	// Invalid
	// const myUint int64 = 1000000000000000000000000000000000000000000000
	var answer = 3 * 0.333
	_ = answer
}
