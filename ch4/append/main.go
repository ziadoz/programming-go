package main

import (
	"fmt"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	y = appendInt(y, 99, 98, 97, 96)
	fmt.Printf("%d  cap=%d\t%v\n", 10, cap(y), y)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*cap(x) {
			zcap = 2 * cap(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Built-in function.
	}
	copy(z[len(x):], y)
	return z
}
