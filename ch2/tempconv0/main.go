package main

import (
	"fmt"

	"gopl.io/ch2/tempconv0/tempconv"
)

func main() {
	// Exported
	fmt.Println(tempconv.BoilingC)      // 100ºC
	fmt.Println(tempconv.FreezingC)     // -273.15ºC
	fmt.Println(tempconv.AbsoluteZeroC) // 0ºC

	// Example One
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC) // 100ºC
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC)) // 180ºC

	// Example Two
	c := tempconv.FToC(212.0)
	fmt.Println(c.String()) // 100ºC - Call String() manually
	fmt.Printf("%v\n", c)   // 100ºC - Value in default format
	fmt.Printf("%s\n", c)   // 100ºC - Bytes of string or slice
	fmt.Println(c)          // 100ºC - Call String() automatically
	fmt.Printf("%g\n", c)   // 100 - Floating point precision
	fmt.Println(float64(c)) // 100 - Convert to float64
}
