// Xmlselect prints the text of selected elements of an XML document.
// Usage: go run main.go https://www.w3.org/TR/2006/REC-xml11-20060816/
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("xmlnodetree")

	if len(os.Args) == 0 {
		log.Fatalf("Please specify a URL")
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalf("Could not get XML: %s", err)
	}

	defer resp.Body.Close()

	tree, err := ParseXML(resp.Body)
	if err != nil {
		log.Fatalf("Could not parse XML into node tree: %s", err)
	}

	fmt.Printf("Node Tree: \n%+v\n\n", tree)
}

func ParseXML(src io.Reader) (Node, error) {
	decoder := xml.NewDecoder(src)

	var root Node        // The root node that represents the final tree.
	var stack []*Element // A running stack to keep track of Elements

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			// Start building the element.
			element := &Element{
				Type: tok.Name,
				Attr: tok.Attr,
			}

			if len(stack) == 0 {
				// If the stack is empty this element is the root node.
				root = element
			} else {
				// Otherwise we need to append the element as a child.
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, element)
			}

			// Push the element onto the running stack.
			stack = append(stack, element)

		case xml.CharData:
			// Push the chardata onto the last element of the stack.
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(tok))
			}

		case xml.EndElement:
			// Pop the last element off of the stack because we're done with it.
			stack = stack[:len(stack)-1]
		}
	}

	return root, nil
}
