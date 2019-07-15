package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // An outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// Creates a channel of all the connected clients.
// Forwards messages from the global messages channel to all connected clients.
// Records clients that have entered the chat.
// Removes clients from the channel if they've left.
func broadcaster() {
	clients := make(map[client]bool) // All connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming messages to all clients outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// Creates a new outgoing message channel for the client.
// Starts a goroutine for the client that receives outgoing messages and sends them to the client's connection.
// Records entering the chat.
// Scans text from the client and sends it to the global messages channel.
// Records leaving the chat.
// Closes the client's connection.
func handleConn(conn net.Conn) {
	ch := make(chan string) // Outgoing client messages.
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	//NOTE: Ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// Receives outgoing messages and sends them to the client's connection.
// Ends when the channel is closed when the client leaves the chat.
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: Ignoring network errors.
	}
}
