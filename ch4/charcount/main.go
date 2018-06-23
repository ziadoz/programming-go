// Charcount computes the counts of Unicode characters.
// Usage: go run main.go < chars.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Count of unicode characters.
	var utflen [utf8.UTFMax + 1]int // Count of lengths of UTF-8 encodings (an array with a length of 5).
	invalid := 0                    // Count of invalid UTF-8 characters.

	in := bufio.NewReader(os.Stdin)
	for {
		// Returns rune, number of bytes and error.
		r, n, err := in.ReadRune()

		// If we've reached the end of the file we're done.
		if err == io.EOF {
			break
		}

		// If there was error we need to exit the program.
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		// If the rune we parsed is the replacement character �, then we record it as invalid,
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++ // Record the rune count.
		utflen[n]++ // Record the rune byte count.
	}

	// Output the runes and counts.
	fmt.Printf("rune\tcounts\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	// Output the rune byte lengths and counts.
	fmt.Printf("\nlen\tcounts\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	// Output any invalid characters,
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
