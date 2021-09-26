package main

import (
	"fmt"
	"unicode/utf8"
)

func city() {
	city := "Kraków"
	fmt.Println(len(city))
	fmt.Println(utf8.RuneCountInString(city))
}
