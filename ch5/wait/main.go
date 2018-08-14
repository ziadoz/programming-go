// Usage: wait https://httpstat.us/200?sleep=5000
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Customise log formatting.
	log.SetPrefix("wait: ") // Add a prefix.
	log.SetFlags(0)         // Disable default flags.

	if len(os.Args) != 2 {
		log.Fatalf("usage: wait [url]\n")
	}

	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		log.Fatalf("site is down: %v\n", err)
	}

	fmt.Printf("visited %s\n", url)
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}

		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
