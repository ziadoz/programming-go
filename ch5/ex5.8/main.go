// Usage:   elembyid [url] [id]
// Example: elembyid http://golang.org gopher
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	log.SetPrefix("elembyid: ")
	log.SetFlags(0)

	if len(os.Args) != 3 {
		log.Fatalf("usage: elembyid [url] [id]")
	}

	url, id := os.Args[1], os.Args[2]

	doc, err := fetch(url)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	node := elementByID(doc, id)
	if node != nil {
		fmt.Printf("matched element %s to #%s\n", node.Data, id)
	} else {
		fmt.Printf("could not find element matching #%s\n", id)
	}
}

func fetch(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func elementByID(doc *html.Node, id string) *html.Node {
	pre := func(node *html.Node) bool {
		if node.Type != html.ElementNode {
			return true
		}

		for _, attr := range node.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}

		return true
	}

	return forEachNode(doc, pre, nil)
}

// forEachNode calls the function pre(x) and post(x) for each node
// x in the tree root at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(node *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil && !pre(node) {
		return node
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result := forEachNode(c, pre, post); result != nil {
			return result
		}
	}

	if post != nil && !post(node) {
		return node
	}

	return nil
}
