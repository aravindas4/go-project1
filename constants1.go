package main

import (
	"fmt"
)

func main() {
	const a = 50
	fmt.Println("a =", a)

	const (
		name = "Aravinda"
		age = 27
		country = "Bhaarata"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	const typedhello string = "hello, world"
	type myString string
	var customName myString = "Sam"
	fmt.Println(customName)

	const trueConst = true
	type myBool bool 
	var defaultBool = trueConst 
	var customBool myBool = trueConst
	// trueConst = customBool
	fmt.Println(trueConst && defaultBool)
	fmt.Println(customBool)

	const b = 5
	var intVar int = b
	var int32Var int32 = b
	var float64Var float64 = b
	var complex64Var complex64 = b 
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}