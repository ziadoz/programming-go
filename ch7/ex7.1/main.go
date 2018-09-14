package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

// Satisfies the io.Writer interface contract.
func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return len(p), nil
}

type LineCounter int

// Satisfies the io.Writer interface contract.
func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	return len(p), nil
}

func main() {
	// WordCounter
	var w WordCounter
	w.Write([]byte("This is a sentence containing lots of words."))
	fmt.Println(w) // 8

	w = 0 // Reset the counter.
	var sentence = "This is another sentence of words."
	fmt.Fprintf(&w, "Sentence: %s", sentence)
	fmt.Println(w) // 7

	// LineCounter
	var l LineCounter
	l.Write([]byte("Line. \n Line. \n Line"))
	fmt.Println(l) // 3

	l = 0 // Reset the counter.
	var lines = "Line \n Line \n Line \t Line \r Line \n Line."
	fmt.Fprintf(&l, "Lines: \n %s", lines)
	fmt.Println(w) // 7
}
