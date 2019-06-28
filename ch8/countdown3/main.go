package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte from stdin
		abort <- struct{}{}
	}()

	// This carries on ticking when the main goroutine is down.
	// This is a goroutine leak.
	tick := time.After(1 * time.Second)
	fmt.Println("Commencing countdown. Press return to abort.")
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launch!")
}
