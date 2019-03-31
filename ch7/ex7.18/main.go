// Xmlselect prints the text of selected elements of an XML document.
// Usage: go run main.go https://www.w3.org/TR/2006/REC-xml11-20060816/
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Errorf("Could not load XML: %s", err)
	}

	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	element := ""
	depth := 0

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			fmt.Println(strings.Repeat(" > ", depth), tok.Name.Local)
			element = tok.Name.Local
			depth++
		case xml.EndElement:
			depth--
			if tok.Name.Local != element {
				fmt.Println(strings.Repeat(" > ", depth), tok.Name.Local)
			}
		}
	}
}
