package main

import "fmt"

func nothing() {
	// Will not compile as it does not have a type
	// nill := nil
	var nill *string = nil
	fmt.Println(nill)
	var nill2 string // Empty string
	fmt.Printf("%x", nill2)
	fmt.Printf("%+q", nill2)
}
