// Usage: go run main.go is:open json decoder
package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Sort slices by time descending.
	sort.Slice(result.Items, func(i, j int) bool {
		a := result.Items[i]
		b := result.Items[j]
		return a.CreatedAt.After(b.CreatedAt)
	})

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// Field formatting:
		// - Literal # Symbol
		// - 5 Digits, Padded Right
		// - String
		// - 9 String Character Width, 9 Precision
		// - Default Character Width, 55 Precision
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}
}
