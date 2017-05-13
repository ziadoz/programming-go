// Dup1 prints the text of each line that appears more than once in the
// standard input, preceded by its count.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main () {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    // Press Ctrl+D to end the input scanning by sending EOF.
    for input.Scan() {
        counts[input.Text()]++
    }

    // Note: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
