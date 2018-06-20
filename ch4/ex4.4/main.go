package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5}
	fmt.Println(i) // [1 2 3 4 5]

	rotate(i, 2)
	fmt.Println(i) // [3 4 5 1 2]

	rotate(i, 3)
	fmt.Println(i) // [1 2 3 4 5]

	rotate(i, 4)
	fmt.Println(i) // [5 1 2 3 4]
}

func rotate(nums []int, by int) {
	// Create a temporary slice large enough to contain the number of elements we're rotating by.
	chunk := make([]int, by)

	// Copy the number of elements we're rotating by off of the original slice and into the temporary slice.
	copy(chunk, nums[:by])

	// Copy the remaining elements into the original slice.
	copy(nums, nums[by:])

	// Then copy the temporary slice onto the end.
	copy(nums[len(nums)-by:], chunk)
}
