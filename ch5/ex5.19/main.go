package main

import "fmt"

func main() {
	fmt.Println(die())
}

func die() (num int) {
	defer func() {
		num = recover().(int) + 10
	}()

	panic(5)
}
