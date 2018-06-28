package main

import (
	"fmt"
	"math/rand"
	"sort"

	"gopl.io/ch4/treesort/treesort"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Println("Unsorted Integers: ", data)

	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		fmt.Println("The integers could not be sorted")
	} else {
		fmt.Println("Sorted Integers", data)
	}
}
