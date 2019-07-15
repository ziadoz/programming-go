package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

type client struct {
	name     string      // The name of the connected client
	outgoing chan string // An outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // All connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming messages to all clients outgoing message channels.
			for cli := range clients {
				cli.outgoing <- msg
			}
		case cli := <-entering:
			online := []string{}
			for client, _ := range clients {
				online = append(online, client.name)
			}

			if len(online) > 0 {
				cli.outgoing <- "currently online: " + strings.Join(online, ", ")
			}

			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outgoing)
		}
	}
}

func handleConn(conn net.Conn) {
	client := client{
		name:     conn.RemoteAddr().String(),
		outgoing: make(chan string),
	}

	go clientWriter(conn, client.outgoing)

	client.outgoing <- "You are " + client.name
	messages <- client.name + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- client.name + ": " + input.Text()
	}
	//NOTE: Ignoring potential errors from input.Err()

	leaving <- client
	messages <- client.name + " has left"
	conn.Close()
}

// Receives outgoing messages and sends them to the client's connection.
// Ends when the channel is closed when the client leaves the chat.
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: Ignoring network errors.
	}
}
