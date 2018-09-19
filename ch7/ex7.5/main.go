// Implement a limit reader that wraps another reader but is exhausted after n bytes.
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"gopl.io/ch7/ex7.5/limitreader"
)

func main() {
	reader := limitreader.LimitReader(strings.NewReader("Hello, Gophers!"), 5)
	buffer := &bytes.Buffer{}

	bytes, _ := buffer.ReadFrom(reader)
	fmt.Println(bytes)  // 5
	fmt.Println(buffer) // Hello

	text, _ := ioutil.ReadAll(reader)
	fmt.Println(string(text)) // Nothing, reader is exhausted.
}
