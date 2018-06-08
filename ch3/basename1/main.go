package main

import "fmt"

func main() {
	fmt.Println(basename("a/b/c.go"))   // c
	fmt.Println(basename("c.d.go"))     // c.d
	fmt.Println(basename("abc"))        // abc
	fmt.Println(basename("a/b/c.d.go")) // c.d
}

// Basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:] // Substring everything after '/' to the end of the string.
			break
		}
	}

	// Preserve everything before the last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i] // Substring everything before the '.' at this position.
			break
		}
	}

	return s
}
