// Dup4 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of names files.
// Usage: go run main.go
//        go run main.go fake.txt
//        go run main.go foo.txt
//        go run main.go foo.txt bar.txt
package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]

    // Handle stdin input and then exit afterwards.
    if (len(files) == 0) {
        countLines(os.Stdin, counts)
        fmt.Println("Duplicate lines captured from standard input: ")
        printLines(counts)
        os.Exit(0);
    }

    // Handle list of files passed in as arguments.
    occurrences := make(map[string][]string)

    for _, filename := range files {
        file, err := os.Open(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
            continue
        }

        fileCounts := make(map[string]int)
        countLines(file, fileCounts)
        file.Close()

        for line, _ := range fileCounts {
            occurrences[line] = append(occurrences[line], filename)
        }

        for line, n := range fileCounts {
            counts[line] += n
        }
    }

    fmt.Println("Duplicate lines captured from files: ")
    printLines(counts)

    fmt.Println("\nDuplicate lines appeared in the following files: ")
    printOccurrences(occurrences)
}

func countLines(file *os.File, counts map[string]int) {
    input := bufio.NewScanner(file)

    for input.Scan() {
        counts[input.Text()]++
    }
}

func printLines(counts map[string]int) {
    for line, n := range counts {
        if (n > 1) {
            fmt.Printf("%d - %s\n", n, line)
        }
    }
}

func printOccurrences(occurrences map[string][]string) {
    for line, files := range occurrences {
       fmt.Printf("%s - %s \n", line, strings.Join(files, ", "))
    }
}
