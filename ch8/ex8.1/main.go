// Clockwall reads multiple clocks to build a clock wall.
// Start clocks:
// TZ=US/Eastern    ./clock2 -port 8010 &
// TZ=Asia/Tokyo    ./clock2 -port 8020 &
// TZ=Europe/London ./clock2 -port 8030 &
//
// Usage: ./clockwall US=localhost:8010 Japan=localhost:8020 UK=localhost:8030
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Clock struct {
	Name string
	Conn net.Conn
}

// Dial a connection to the clock.
func (c *Clock) Dial(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	c.Conn = conn
}

// Watch the clock and write its output to the dst io.Writer.
func (c *Clock) Watch(dst io.Writer) {
	s := bufio.NewScanner(c.Conn)

	for s.Scan() {
		fmt.Fprintf(dst, "%s: %s\n", c.Name, s.Text())
	}

	fmt.Fprintln(dst, "Done!")

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: name=host [name=host...]")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2 {
			fmt.Printf("skipped invalid clock: %s\n", arg)
			continue
		}

		clock := &Clock{Name: fields[0]}
		clock.Dial(fields[1])

		defer clock.Conn.Close()
		go clock.Watch(os.Stdout)
	}

	time.Sleep(time.Minute)
}
