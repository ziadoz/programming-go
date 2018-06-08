package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("+1234567890.123")) // +1,234,567,890.123
	fmt.Println(comma("1234567890.123"))  // 1,234,567,890.123
	fmt.Println(comma("-1234567890.123")) // -1,234,567,890.123
	fmt.Print("\n")

	fmt.Println(comma("+12345.123")) // +12,345.123
	fmt.Println(comma("12345.123"))  // 12,345.123
	fmt.Println(comma("-12345.123")) // -12,345.123
	fmt.Print("\n")

	fmt.Println(comma("+1234.123")) // +1,234.123
	fmt.Println(comma("1234.123"))  // 1,234.123
	fmt.Println(comma("-1234.123")) // -1,234.123
	fmt.Print("\n")

	fmt.Println(comma("+123.123")) // +123.123
	fmt.Println(comma("123.123"))  // 123.123
	fmt.Println(comma("-123.123")) // -123.123
}

func comma(s string) string {
	var buffer bytes.Buffer

	sign, s := sign(s)
	decimal, s := decimal(s)
	start := start(s)

	// Write any sign to the buffer.
	buffer.WriteString(sign)

	// Substring that many characters off the string and into the buffer.
	buffer.WriteString(s[:start])

	// Work through the remaining part, which is now exactly divisible by three.
	// Substring out each chunk and write a comma and then that into the buffer.
	for i := start; i < len(s); i += 3 {
		buffer.WriteRune(',')
		buffer.WriteString(s[i : i+3])
	}

	// Now we're done, write any decimal to the buffer.
	buffer.WriteString(decimal)

	return buffer.String()
}

// Figure out if the first character is a sign.
// If so, substring it off and store it, and swap out the original string minus the sign.
func sign(s string) (sign string, result string) { // Named return parameters.
	result = s
	if s != "" && s[0] == '+' || s[0] == '-' {
		sign = s[:1]
		result = s[1:]
	}

	return sign, result
}

// Figure out if there are any decimals.
// If so, substring them off, and swap out the original string minus that.
func decimal(s string) (decimal string, result string) { // Named return parameters.
	result = s
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		decimal = s[dot:]
		result = s[:dot-1]
	}

	return decimal, result
}

// Work out if the number is odd/even by doing modulus to get the remainder.
// This tells us how many numbers we need to substring before the first comma.
func start(s string) int {
	start := len(s) % 3
	if start == 0 {
		start = 3
	}

	return start
}
