// Usage:   findlinks3 [url]
// Example: findlinks3 http://golang.org
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token (push value into channel)
	list, err := links.Extract(url)
	<-tokens // release the token (read value of channel)

	if err != nil {
		log.Print(err)
	}

	return list
}
