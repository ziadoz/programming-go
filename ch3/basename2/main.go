package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go"))   // c
	fmt.Println(basename("c.d.go"))     // c.d
	fmt.Println(basename("abc"))        // abc
	fmt.Println(basename("a/b/c.d.go")) // c.d
}

// Basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	if slash := strings.LastIndex(s, "/"); slash >= 0 {
		s = s[slash+1:]
	}

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
