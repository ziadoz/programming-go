package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var hello string = "Hello, World!"

func main() {
	log.SetFlags(0)
	log.SetPrefix("webserver: ")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		local := "Hello, Local!"

		fmt.Fprintf(os.Stdout, "%v\n", &local)
		fmt.Fprintf(os.Stdout, "%v\n", &hello)

		writer.Write([]byte(hello))
		writer.Write([]byte("\n"))
		writer.Write([]byte(local))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
