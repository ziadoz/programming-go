package main

import "fmt"

func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	addEdge("a", "d")
	addEdge("d", "a")
	fmt.Println(graph)
	fmt.Println(hasEdge("a", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("a", "d"))
	fmt.Println(hasEdge("d", "a"))
	fmt.Println(hasEdge("x", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("x", "d"))
	fmt.Println(hasEdge("d", "x"))
}

// A map of string keys with a map of string boolean values.
var graph = make(map[string]map[string]bool)

// Which would look something like this in a dynamic language:
// {
// 	"a" => {
// 		"b" => true,
// 		"c" => true,
// 	},
// }

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool) // Lazily populate the map.
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
