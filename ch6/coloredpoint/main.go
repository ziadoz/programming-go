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
	// ColoredPoint embeds Point. Fields can be accessed either way.
	var cp ColouredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // 1
	cp.Point.Y = 2
	fmt.Println(cp.Y) // 2

	// We can call methods on ColoredPoint that only exist on Point due to embedding.
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColouredPoint{Point{1, 1}, red}
	var q = ColouredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // 5 - Need to call on the embedded Point.
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // 10 - Need to call on the embedded Point.
}
