// Exercise 8.3 - Close the read and write halves of a TCP connection manually.
// Used type assertions to determine concrete type and then call the methods CloseRead() and CloseWrite().
// Could have used a type switch too and fallen back to Close().
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: Ignoring errors
		conn.(*net.TCPConn).CloseRead()
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin) // Ctrl+D will send EOF and this will stop blocking
	conn.(*net.TCPConn).CloseWrite()
	<-done // wait for the background goroutine to finsh
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
