package main

import (
	"context"
	"fmt"
)

func fibs(ctx context.Context, n int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch) // No deadlock error
		a, b := 1, 1
		for i := 0; i < n; i++ {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-ctx.Done():
				fmt.Println("cancelled")
			}
		}
	}()

	return ch
}

func rang() {
	ctx, cancel := context.WithCancel(context.Background())
	channel := fibs(ctx, 5)
	for i := 0; i < 5; i++ {
		val := <-channel
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	cancel()
}
