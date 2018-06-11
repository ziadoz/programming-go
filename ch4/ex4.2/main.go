package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var allowedAlgos = [...]string{}

var algo string

func init() {
	flag.StringVar(&algo, "algo", "sha256", "The hashing algorithm to use (sha256, sha384 or sha512)")
	flag.Parse()
}

func main() {
	switch algo {
	case "sha256":
	case "sha384":
	case "sha512":
		// Do nothing.
	default:
		fmt.Printf("Invalid hash algorithm: '%s'\n", algo)
		os.Exit(1)
	}

	fmt.Println("Enter some text to hash using SHA256 (or Ctrl+C to exit): ")

	var arg string
	for true {
		fmt.Scan(&arg)

		switch algo {
		case "sha256":
			fmt.Printf("%x\n", sha256.Sum256([]byte(arg)))
		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		}
	}
}
