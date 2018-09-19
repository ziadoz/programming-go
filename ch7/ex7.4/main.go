// Implement something like strings.NewReader() that satisfies io.Reader.
package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type StringReader struct {
	text string
}

func (sr *StringReader) Read(p []byte) (bytes int, err error) {
	bytes = copy(p, sr.text)  // Copy bytes to the slice.
	sr.text = sr.text[bytes:] // Remove the number of copied bytes from the struct text.

	// If the length of the struct text is zero, we need to return an EOF error.
	if len(sr.text) == 0 {
		err = io.EOF
	}

	return
}

func NewStringReader(s string) io.Reader {
	return &StringReader{text: s}
}

func main() {
	// Parse HTML from string via StringReader.
	markup := `
	<html>
		<title>Hello, World!</title>
		<p>Gophers, baby, gophers!</p>
	</html>
	`

	doc, err := html.Parse(NewStringReader(markup))
	if err != nil {
		log.Fatalln(err)
	}

	forEachNode(doc, nil, func(node *html.Node) {
		if data := strings.TrimSpace(node.Data); node.Type == html.TextNode && data != "" {
			fmt.Println(node.Data)
		}
	})
}

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, post)
	}

	if post != nil {
		post(node)
	}
}
