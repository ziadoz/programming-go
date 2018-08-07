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
		fmt.Fprintf(os.Stderr, "findtext: %s", err)
		os.Exit(1)
	}

	for _, text := range TextualContent(nil, doc) {
		fmt.Println(text)
	}
}

func TextualContent(text []string, node *html.Node) []string {
	if IsReadableNode(node) {
		if content := strings.TrimSpace(ExtractText(node)); content != "" {
			println(node.Data, content)
			text = append(text, content)
		}
	}

	if node.FirstChild != nil {
		text = TextualContent(text, node.FirstChild)
	}

	if node.NextSibling != nil {
		text = TextualContent(text, node.NextSibling)
	}

	return text
}

func IsReadableNode(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}

	elems := []string{"script", "noscript", "style", "link"}

	for _, elem := range elems {
		if node.Data == elem {
			return false
		}
	}

	return true
}

func ExtractText(node *html.Node) string {
	text := []string{}

	if fc := node.FirstChild; fc != nil && fc.Type == html.TextNode {
		text = append(text, fc.Data)
	}

	if ns := node.NextSibling; ns != nil && ns.Type == html.TextNode {
		text = append(text, ns.Data)
	}

	return strings.Join(text, " ")
}
