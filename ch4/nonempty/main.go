// Nonempty is an example of an inplace slice algorithm.
package main

import "fmt"

func main() {
	// Slice manipulated in place.
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one", "three"]
	fmt.Printf("%q\n", data)           // ["one", "three", "three"]

	// Overwrite slice variable.
	data2 := []string{"one", "", "three"}
	data2 = nonempty(data2)
	fmt.Printf("%q\n", data2) // ["one", "three"]

	// Using append() nonempty variant.
	data3 := []string{"one", "", "three"}
	data3 = nonempty2(data3)
	fmt.Printf("%q\n", data3) // ["one", "three"]
}

// Nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// Nonempty written using append().
func nonempty2(strings []string) []string {
	out := strings[:0] // Zero length version of the original.
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
