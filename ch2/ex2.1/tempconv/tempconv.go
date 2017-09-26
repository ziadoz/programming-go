// Package tempconv performs Celcius to Farenheit conversions.
package tempconv

import "fmt"

type Celcius float64
type Farenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func (c Celcius) String() string {
	return fmt.Sprintf("%gºC", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
