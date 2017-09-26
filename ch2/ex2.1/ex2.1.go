package main

import (
	"fmt"
	"gopl.io/ch2/ex2.1/tempconv"
)

func main() {
	// Celcius to Kelvin
	c := tempconv.Celcius(0)
	fmt.Println(tempconv.CToK(c)) // 273.15K

	c = tempconv.Celcius(20)
	fmt.Println(tempconv.CToK(c)) // 293.15K

	c = tempconv.Celcius(100)
	fmt.Println(tempconv.CToK(c)) // 373.15K

	// Farenheit to Kelvin
	f := tempconv.Farenheit(0)
	fmt.Println(tempconv.FToK(f)) // 255.3722222222222K

	f = tempconv.Farenheit(20)
	fmt.Println(tempconv.FToK(f)) // 266.4833333333333K

	f = tempconv.Farenheit(100)
	fmt.Println(tempconv.FToK(f)) // 310.92777777777775K

	// Kelvin to Celcius and Farenheit
	k := tempconv.Kelvin(0)
	fmt.Println(k) // 0K

	k = tempconv.Kelvin(50)
	fmt.Println(tempconv.KToC(k)) // -223.14999999999998ºC
	fmt.Println(tempconv.KToF(k)) // -369.66999999999996ºF

	k = tempconv.Kelvin(1000)
	fmt.Println(tempconv.KToC(k)) // 726.85ºC
	fmt.Println(tempconv.KToF(k)) // 1340.3300000000002ºF
}
