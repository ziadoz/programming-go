package main

import "fmt"

func main() {
	fmt.Println(comma("1234567890")) // 1,234,567,890
}

// Comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}

	// Recurisvely call comma() with everything but the last three characters.
	// Then append a substring containing just the last three characters to the end of it.
	// Uncomment line below to show how this breaks the string down recursively.
	// fmt.Println(s)
	return comma(s[:n-3]) + "," + s[n-3:]
}
