package main

import "fmt"

func main() {
	// Using slices to implement a stack.
	var stack []string

	// Push an item onto the stack.
	stack = append(stack, "hello")
	stack = append(stack, "world")
	fmt.Println(stack)

	// Get the top of the stack (the last element).
	top := stack[len(stack)-1]
	fmt.Println(top)

	// Shrink the stack by one by the popping of the last element.
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	// Removing items from a stack whilst preserving order.
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(s)

	s = remove(s, 2)
	fmt.Println(s) // [5, 6, 8, 9]

	s = remove(s, 0)
	fmt.Println(s) // [6, 8, 9]

	// Removing items from a stack without preserving order.
	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(s)

	s2 = remove2(s2, 2)
	fmt.Println(s2) // [5, 6, 9, 8]

	s2 = remove2(s2, 0)
	fmt.Println(s2) // [8, 6, 9]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:]) // Slide higher numbered elements down by one to fill in gap.
	return slice[:len(slice)-1]  // Return slice minus the last element.
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // Move last element into the gap.
	return slice[:len(slice)-1]    // Return slice minus the last element.
}
