// Usage: go run main.go is:open json decoder
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// Field formatting:
		// Literal # Symbol
		// 5 Digits, Padded Right
		// 9 String Character Width, 9 Precision
		// Default Character Width, 55 Precision
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
