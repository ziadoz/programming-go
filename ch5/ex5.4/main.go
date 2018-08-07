// Usage: curl https://www.bbc.co.uk | go run main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, link := range ExtractLinks(nil, doc) {
		fmt.Println(link)
	}
}

func ExtractLinks(links []string, node *html.Node) []string {
	if IsLinkNode(node) {
		if link := strings.TrimSpace(ExtractLink(node)); link != "" {
			links = append(links, link)
		}
	}

	if node.FirstChild != nil {
		links = ExtractLinks(links, node.FirstChild)
	}

	if node.NextSibling != nil {
		links = ExtractLinks(links, node.NextSibling)
	}

	return links
}

func IsLinkNode(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}

	for _, elem := range []string{"a", "link", "img", "script"} {
		if node.Data == elem {
			return true
		}
	}

	return false
}

func ExtractLink(node *html.Node) string {
	attrs := map[string]string{
		"a":      "href",
		"link":   "href",
		"img":    "src",
		"script": "src",
	}

	for _, attr := range node.Attr {
		if attr.Key == attrs[node.Data] {
			return attr.Val
		}
	}

	return ""
}
