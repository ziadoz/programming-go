// Echo2 prints its command line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    // 1. s: = ""
    // 2. var s string
    // 3. var s = ""
    // 4. var s string = ""
    // Prefer forms 1 and 2 of variable declaration where possible.
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
