// Usage: go run main.go is:open json decoder
//        go run main.go is:open laravel eloquent
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	categories := map[string]time.Time{
		"less than a month old": time.Date(now.Year(), now.Month(), 01, 0, 0, 0, 0, time.UTC),
		"less than a year old":  time.Date(now.Year()-1, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
		"older than a year":     time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
	}

	before := now
	fmt.Printf("%d issues:\n\n", result.TotalCount)
	for category, after := range categories {
		issues := result.FilterBetweenTimes(after, before)
		issues.SortByCreatedAtDescending()
		before = after

		if len(issues.Items) == 0 {
			continue
		}

		fmt.Println(category)
		for _, item := range issues.Items {
			fmt.Printf("#%-5d %s %9.9s %.55s\n",
				item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
		}
		fmt.Println()
	}
}

// Field formatting:
// - Literal # Symbol
// - 5 Digits, Padded Right
// - String
// - 9 String Character Width, 9 Precision
// - Default Character Width, 55 Precision
