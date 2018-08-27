package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColouredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // Method expression. No receiver or parameters.
	fmt.Println(distance(p, q))  // 5 - Call method expression by passing receiver, then regular parameters as normal.
	fmt.Printf("%T\n", distance) // func(Point, Point) float64

	scale := (*Point).ScaleBy // Method expression on a pointer. Again, no receiver or parameters.
	scale(&p, 2)              // Again, receiver is passed first, then regular parameters method would accept.
	fmt.Println(p)            // {2, 4}
	fmt.Printf("%T\n", scale) // func(*Point, float64)
}
