// Count the number of distinct character types in a document.
// Usage: go run main.go < chars.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	chars := map[string]map[rune]int{
		"letter":      make(map[rune]int),
		"digit":       make(map[rune]int),
		"punctuation": make(map[rune]int),
		"space":       make(map[rune]int),
		"symbol":      make(map[rune]int),
	}
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if unicode.ReplacementChar == r && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r):
			chars["letter"][r]++
		case unicode.IsDigit(r):
			chars["digit"][r]++
		case unicode.IsPunct(r):
			chars["punctuation"][r]++
		case unicode.IsSpace(r):
			chars["space"][r]++
		case unicode.IsSymbol(r):
			chars["symbol"][r]++
		}
	}

	for category, runes := range chars {
		if len(runes) == 0 {
			continue
		}

		fmt.Printf("Identified %d characters of the category %s:\n", len(runes), strings.Title(category))
		for char, count := range runes {
			fmt.Printf("%q\t%d\n", char, count)
		}
	}
}
