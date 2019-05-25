package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 1. Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// 2. Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// 3. Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
