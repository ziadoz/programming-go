package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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

const idle = time.Second * 15

type client struct {
	conn     net.Conn       // The client connection
	name     string         // The name of the connected client
	outgoing chan string    // An outgoing message channel
	active   chan time.Time // A last active timestamp channe;
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
		conn:     conn,
		name:     conn.RemoteAddr().String(),
		outgoing: make(chan string),
		active:   make(chan time.Time),
	}

	go clientWriter(client)

	client.outgoing <- "You are " + client.name
	messages <- client.name + " has arrived"
	entering <- client

	// Create a timer that once received kicks the client.
	timer := time.NewTimer(idle)
	go func() {
		<-timer.C
		msg := client.name + " kicked for being idle"
		fmt.Fprintln(client.conn, msg)
		client.conn.Close()
		messages <- msg
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- client.name + ": " + input.Text()
		timer.Reset(idle) // Reset the timer once a message is received.
	}
	// NOTE: Ignoring potential errors from input.Err()

	leaving <- client
	messages <- client.name + " has left"
	conn.Close()
}

func clientWriter(client client) {
	for msg := range client.outgoing {
		fmt.Fprintln(client.conn, msg) // NOTE: Ignoring network errors.
	}
}
