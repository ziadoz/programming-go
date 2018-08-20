package main

import "fmt"

func main() {
	fmt.Println(sum())  // 0
	fmt.Println(sum(3)) // 3

	// This:
	fmt.Println(sum(1, 2, 3, 4)) // 3

	// Is essentially the same as this:
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	// Functions with slices are not the same as those with variadic parameters.
	f := func(...int) {}
	g := func([]int) {}

	fmt.Printf("%T\n", f) // func(...int)
	fmt.Printf("%T\n", g) // func([]int)
}

// Variadic function.
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
