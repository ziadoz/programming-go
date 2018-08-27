package main

import (
	"fmt"
	"time"
)

type Greeter struct {
	Name string
}

func (g Greeter) Greet(greeting string) string {
	return fmt.Sprintf("%s, %s!", greeting, g.Name)
}

type Rocket struct{}

func (r *Rocket) Launch() {
	fmt.Println("Rocket Launched!")
}

func main() {
	g := Greeter{Name: "World"}
	greet := g.Greet // Method value - Receiver is bound, parameters to be passed in.
	fmt.Println(greet("Hello"))

	r := &Rocket{}
	time.AfterFunc(2*time.Second, func() { r.Launch() }) // Call method inside an anonymous func.
	time.AfterFunc(2*time.Second, r.Launch)              // Or use a method value instead.

	time.Sleep(4 * time.Second)
}
