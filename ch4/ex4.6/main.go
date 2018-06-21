package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	bs := []byte("Hello, \t\n World! 🤘 \t\t Let's\n\nRock!")
	bs = trimAdjacentSpaces(bs)
	fmt.Println(string(bs))
}

// Trims down any adjacent spaces
func trimAdjacentSpaces(bs []byte) []byte {
	out := bs[:0]            // A new output slice based upon the original.
	runes := bytes.Runes(bs) // Turning the byte slice into runes guarantees things will work with any characters (e.g. emojis).
	index := 0               // Keep a running index so we know exactly how many characters we've inserted.

	for _, char := range runes {
		if unicode.IsSpace(char) {
			if index > 0 && unicode.IsSpace(runes[index-1]) {
				// We're in multiple spaces here, so we don't need to append anything.
				continue
			} else {
				out = append(out, ' ') // Append a single space to the output instead of multiple.
				index++
				continue
			}
		}

		// Here we turn the rune into a string, and then the string into a byte slice so we can append it to the output.
		for _, b := range []byte(string(char)) {
			out = append(out, b)
			index++
		}
	}

	return out[:index]
}
