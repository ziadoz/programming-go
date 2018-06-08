package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890")) // 1,234,567,890
	fmt.Println(comma("12345"))      // 12,345
	fmt.Println(comma("1234"))       // 1,234
	fmt.Println(comma("123"))        // 123
}

func comma(s string) string {
	var buffer bytes.Buffer

	// Work out if the number is odd/even by doing modulus to get the remainder.
	// This tells us how many numbers we need to substring before the first comma.
	start := len(s) % 3
	if start == 0 {
		start = 3
	}

	// Substring that many characters off the string and into the buffer.
	buffer.WriteString(s[:start])

	// Work through the remaining part, which is now exactly divisible by three.
	// Substring out each chunk and write a comma and then that into the buffer.
	for i := start; i < len(s); i += 3 {
		buffer.WriteRune(',')
		buffer.WriteString(s[i : i+3])
	}

	return buffer.String()
}
