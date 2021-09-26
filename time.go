package main

import (
	"fmt"
	"time"
)

func timesout() {
	timeout := 3
	// var timeout time.Duration = 3
	// const timeout = 3
	fmt.Println("Before....")
	fmt.Println(time.Duration(timeout) * time.Microsecond)
	fmt.Println("After....")
}
