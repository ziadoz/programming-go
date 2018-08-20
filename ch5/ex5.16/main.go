package main

import "fmt"

func main() {
	fmt.Println(join("Hello", ", ", "World", "!"))
}

func join(parts ...string) string {
	str := ""
	for _, part := range parts {
		str += part
	}
	return str
}
