package main

import "fmt"

func main() {
	fmt.Println(min(22, 5, 2, 55, 7, 2, 64))
	fmt.Println(max(22, 5, 2, 55, 7, 2, 64))
}

func min(num int, nums ...int) int {
	nums = append(nums, num)
	min := num
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(num int, nums ...int) int {
	nums = append(nums, num)
	min := num
	for _, num := range nums {
		if num > min {
			min = num
		}
	}
	return min
}
