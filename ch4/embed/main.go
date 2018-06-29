package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// This struct literal creates a struct identical to the one below.
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w1)

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 5},
			Radius: 5, // Trailing comma necessary here.
		},
		Spokes: 20, // Trailing comma necessary here.
	}
	fmt.Printf("%#v\n", w2)

	w2.X = 42
	fmt.Printf("%#v\n", w2)
}
