package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 1. Counter (separate goroutine)
	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// 2. Squarer (separate goroutine)
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// 3. Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
