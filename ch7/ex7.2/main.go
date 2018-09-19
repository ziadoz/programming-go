// Wrap a writer to count the number of bytes written.
package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	writer io.Writer
	count  int
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	bytes, err := b.writer.Write(p)
	b.count += bytes
	return bytes, err
}

func CountingWriter(writer io.Writer) (io.Writer, *int) {
	bc := &ByteCounter{writer: writer, count: 0}
	return bc, &bc.count
}

func main() {
	osw, count := CountingWriter(os.Stdin)
	fmt.Println(*count) // 0
	fmt.Fprintln(osw, "Hello, World")
	fmt.Println(*count) // 13
	fmt.Fprintln(osw, "Always pack a towel.")
	fmt.Println(*count) // 34
}
