package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	// Op is either Add or Sub method expression depending on boolean.
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	path := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	path.TranslateBy(Point{2, 2}, true)
	fmt.Println(path)
	path.TranslateBy(Point{4, 4}, false)
	fmt.Println(path)
}
