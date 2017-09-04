// Package tempconv performs Celcius and Farenheit temperature computations.
package tempconv

import "fmt"

type Celcius float64
type Farenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func CToF(c Celcius) Farenheit   { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celcius   { return Celcius((f - 32) * 5 / 9) }
func (c Celcius) String() string { return fmt.Sprintf("%gºC", c) }
