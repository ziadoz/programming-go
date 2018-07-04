package main

import (
	"fmt"
	"sync"
)

func process(num int, wg *sync.WaitGroup) {
	fmt.Printf("Goroutine #%d\n", num)
	wg.Done() // Decrement counter.
}

func main() {
	num := 25
	fmt.Printf("Spinning up %d goroutines\n", num)

	var wg sync.WaitGroup
	wg.Add(num) // Increment counter.

	for i := 1; i <= num; i++ {
		// wg.Add(1) // Increment counter. Could do this instead of Wait() call above.
		go process(i, &wg)
	}

	fmt.Println("Waiting for goroutines to finish")
	wg.Wait() // Blocks current gorountine until counter is zero.
	fmt.Println("Done!")
}
