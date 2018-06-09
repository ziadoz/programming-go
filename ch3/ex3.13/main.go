package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("%v Bytes = 1KB\n", KB)
	fmt.Printf("%v Bytes = 1MB\n", MB)
	fmt.Printf("%v Bytes = 1GB\n", GB)
	fmt.Printf("%v Bytes = 1TB\n", TB)
	fmt.Printf("%v Bytes = 1PB\n", PB)
	fmt.Printf("%v Bytes = 1EB\n", EB)
	// fmt.Printf("%v Bytes = 1ZB\n", ZB) // Overflows integer.
	// fmt.Printf("%v Bytes = 1YB\n", YB) // Overflows integer.
}
