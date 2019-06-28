// Usage:   findlinks3 [url]
// Example: findlinks3 http://golang.org
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func main() {
	worklist := make(chan []string)

	// Start with the command line arguments.
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
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

func crawl(url string) []string {
	fmt.Println(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
