package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // Don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second)
}

// Starts a timer, then returns a function which can be deferred to time the function.
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
