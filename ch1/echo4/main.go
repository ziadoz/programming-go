// Echo4 prints its command line arguments.
// 1.1 - Print out the name of the program.
// 1.2 - Print the index value of each argument.
// 1.3 - Time building a string vs using the string package to join.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Running Program:", os.Args[0])
	fmt.Println("Debug:", os.Args[1:])
	fmt.Println(len(os.Args)-1, "arguments or flags passed: ")

	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("String Concatenation: %.2fs elapsed \n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("String Joining: %.2fs elapsed \n", time.Since(start).Seconds())

	for index, arg := range os.Args[1:] {
		fmt.Println("Index/Argument:", index, arg)
	}
}
