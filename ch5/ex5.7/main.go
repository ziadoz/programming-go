// Usage: outline2 http://golang.org
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	doc, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the function pre(x) and post(x) for each node
// x in the tree root at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(node *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

var depth int

func startElement(node *html.Node) {
	switch node.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", node.Data)

		for _, attr := range node.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}

		if node.FirstChild != nil {
			fmt.Printf(">\n")
		} else {
			fmt.Printf(" />\n")
		}

		depth++
	case html.TextNode:
		if text := strings.TrimSpace(node.Data); text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", text)
		}
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", strings.TrimSpace(node.Data))
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		if node.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
		}
	}
}
