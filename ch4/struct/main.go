package main

import "fmt"

func main() {
	// An empty struct variable and an empty struct literal.
	var empty struct{}
	fmt.Println(empty)
	fmt.Println(struct{}{})

	// A basic struct with export fields.
	type Point struct {
		X int
		Y int
	}

	// A struct with fields populated by position.
	p1 := Point{5, 10}
	fmt.Println(p1)

	// A struct with fields populated by name.
	p2 := Point{X: 5, Y: 10}
	fmt.Println(p2)

	// A struct with some fields populated by name (remaining are type defaults).
	p3 := Point{X: 5}
	fmt.Println(p3)

	// Structs are comparable.
	pp1 := Point{X: 1, Y: 2}
	pp2 := Point{1, 2}
	fmt.Println(pp1 == pp2)

	// So are there fields.
	fmt.Println(pp1.X == pp2.X)
	fmt.Println(pp1.Y == pp2.Y)

	// Becuase they're comparable they can be used as map keys.
	type FooBar struct {
		Foo string
		Bar string
	}

	data := map[FooBar]int{
		FooBar{"foo", "bar"}: 5,
		FooBar{"baz", "qux"}: 10,
	}
	fmt.Println(data)

	// You can address a struct literal easily.
	pp3 := &Point{X: 5, Y: 10}
	fmt.Println(pp3)

	// Which is the same as:
	pp4 := new(Point)
	*pp4 = Point{X: 5, Y: 10}
	fmt.Println(pp4)

	// They're passed to functions as a copy.
	double := func(p Point) Point {
		p.X = p.X * 2
		p.Y = p.Y * 2
		return p
	}

	pp6 := Point{10, 10}
	pp7 := double(pp6)
	fmt.Println(pp6, pp7)
	fmt.Println(pp6 == pp7)

	// Unless you use a pointer.
	doublep := func(p *Point) {
		p.X = p.X * 2
		p.Y = p.Y * 2
	}

	pp8 := Point{10, 10}
	fmt.Println(pp8)
	doublep(&pp8)
	fmt.Println(pp8)

	// You can use structs as fields in other structs.
	type MyPoint struct {
		X, Y int // Declare multiple fields on one line.
	}

	type MyCircle struct {
		Centre MyPoint
		Radius int
	}

	type MyWheel struct {
		Circle MyCircle
		Spokes int
	}

	// Each type needs to be accessed to populate the struct.
	var w1 MyWheel
	w1.Circle.Centre.X = 10
	w1.Circle.Centre.Y = 10
	w1.Circle.Radius = 10
	w1.Spokes = 20
	fmt.Println(w1)

	// You can also embed structs as fields in one another.
	type AnotherPoint struct {
		X, Y int
	}

	type AnotherCircle struct {
		AnotherPoint // Embedded.
		Radius       int
	}

	type AnotherWheel struct {
		AnotherCircle // Embedded.
		Spokes        int
	}

	// These embedded struct fields can be access as if they exist on the struct.
	var w2 AnotherWheel
	w2.X = 10
	w2.Y = 10
	w2.Radius = 10
	w2.Spokes = 20
	fmt.Println(w2)

	// But you still need to construct things from scratch.
	w3 := AnotherWheel{
		AnotherCircle: AnotherCircle{
			AnotherPoint: AnotherPoint{10, 10},
			Radius:       10,
		},
		Spokes: 20,
	}
	fmt.Println(w3)
}
