package main

import "fmt"

func main() {
	s := []string{"foo", "bar", "bar", "baz", "qux", "qux", "foo"}
	fmt.Println(s)
	fmt.Println(removeAdjacent(s))
}

func removeAdjacent(strings []string) []string {
	index := 0
	last := ""

	for _, s := range strings {
		if s != last {
			strings[index] = s
			index++
			last = s
		}
	}

	return strings[:index]
}
