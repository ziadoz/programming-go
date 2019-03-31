// Xmlselect prints the text of selected elements of an XML document.
// Usage: curl https://www.w3.org/TR/2006/REC-xml11-20060816/ | go run main.go name=a id=sec-bibliography
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type eleminfo struct {
	name  string
	attrs map[string]string
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*eleminfo

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			//fmt.Println(tok.Name.Local)

			ei := &eleminfo{
				name:  tok.Name.Local,
				attrs: make(map[string]string),
			}

			for _, val := range tok.Attr {
				if val.Value != "" {
					ei.attrs[val.Name.Local] = val.Value
				}
			}

			stack = append(stack, ei)

			if _, ok := ei.attrs["id"]; ok {
				fmt.Printf("%#v\n", ei)
			}
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				// fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				fmt.Println(string(tok))
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		}
	}
}

// containsAll reports whether x contains of y, in order
func containsAll(stack []*eleminfo, exp []string) bool {
	// fmt.Println("starting match ")

	for len(exp) <= len(stack) {
		if len(exp) == 0 {
			return true
		}

		// fmt.Printf("matching %s against %#v\n", exp[0], stack[0])

		if isMatch("name", exp[0]) && getMatchVal("name", exp[0]) == stack[0].name {
			// fmt.Println("name matched ", stack[0].name)
			exp = exp[1:]
		} else if isMatch("id", exp[0]) && getMatchVal("id", exp[0]) == stack[0].attrs["id"] {
			// fmt.Println("id matched ", stack[0].attrs["id"])
			exp = exp[1:]
		} else if isMatch("class", exp[0]) && getMatchVal("class", exp[0]) == stack[0].attrs["class"] {
			// fmt.Println("class matched ", stack[0].attrs["class"])
			exp = exp[1:]
		}

		stack = stack[1:]
	}

	return false
}

// isMatch takes a selector (name, class or id) and then checks if the string begins with "selector="
func isMatch(selector, exp string) bool {
	prefix := selector + "="
	return (len(exp) > len(prefix) && exp[:len(prefix)] == prefix)
}

// getMatchVal returns the value from a string in the form "selector=value"
func getMatchVal(selector, exp string) string {
	prefix := selector + "="
	return exp[len(prefix):]
}
