// Exercise 4.9
// Word frequency determines the number of words in the given input.
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	words := map[string]int{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := CleanWord(scanner.Text())
		words[word]++
	}

	for word, count := range words {
		fmt.Printf("%q\t%d\n", word, count)
	}
}

// Remove any spaces or punctuation from the words.
func CleanWord(word string) (clean string) {
	for _, char := range word {
		if !unicode.IsSpace(char) && !unicode.IsPunct(char) {
			clean += string(char)
		}
	}

	return clean
}
