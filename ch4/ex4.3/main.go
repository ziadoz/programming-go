package main

import "fmt"

func main() {
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	reverse(&nums)
	fmt.Println(nums)
}

func reverse(nums *[10]int) {
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
