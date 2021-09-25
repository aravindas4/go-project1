package main

import "fmt"

func main() {
	var dict map[string]int
	fmt.Println(dict["errors"])

	// Dynamic type
	new := make(map[interface{}]interface{})
	fmt.Println(new)
}
