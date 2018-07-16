package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	elems := TagFrequency()
	fmt.Println("Number of elements in page: ")
	for elem, total := range elems {
		fmt.Printf(" - %s: %d\n", elem, total)
	}
}

func TagFrequency() map[string]int {
	elems := map[string]int{}
	tokenizer := html.NewTokenizer(os.Stdin)
	for {
		token := tokenizer.Next()
		switch token {
		case html.StartTagToken:
			tagname, _ := tokenizer.TagName()
			elems[string(tagname)]++
		case html.ErrorToken:
			return elems
		}
	}
}
