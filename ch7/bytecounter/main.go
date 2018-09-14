package main

import "fmt"

type ByteCounter int

// Satisfies the io.Writer interface contract.
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Convert integer to ByteCounter.
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5 = len("hello").

	c = 0 // Reset the counter.
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12 = len("hello, Dolly")
}
