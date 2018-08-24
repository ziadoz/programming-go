package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x is zero
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
