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
		err := title(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func title(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	// Check content type is HTML (e.g. text/html; charset=utf-8)
	contentType := response.Header.Get("Content-Type")
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html;") {
		response.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, contentType)
	}

	doc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			fmt.Println(node.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)
	return nil
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
