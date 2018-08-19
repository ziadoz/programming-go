// prereqs maps computer science courses to their prerequisites.
package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organisation",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organisation"},
	"programming languages": {"data structures", "computer organisation"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) // Must declare the function variable separately so it can recurse.

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item]) // Recursive.
				order = append(order, item)
			}
		}
	}

	for key := range m {
		visitAll([]string{key})
	}

	return order
}
