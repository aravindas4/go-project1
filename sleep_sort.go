package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep_sort() {
	var wg sync.WaitGroup
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			time.Sleep(time.Duration(number) * time.Millisecond)
			fmt.Printf("%d ", number)
		}(n)
	}
	wg.Wait()
	fmt.Println()
}
