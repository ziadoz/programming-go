// Usage: go run main.go <<< '<html><title>Hello, World</title></html>'
//        go run main.go <<< '<html><title>Hello, World</title><p>Paragraph</p><title>Panic</title></html>'
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	log.SetFlags(0)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("could not parse html: %v\n", err)
	}

	fmt.Println(soleTitle(doc))
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic, carry on panicking
		}
	}()

	// Bail out of recursion if we find more than one non-empty title.
	forEachNode(doc, func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}

			title = node.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
