package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celciusFlag satisfies the flag.Value interface.
// String() is "inherited" from Celcius.
type celciusFlag struct{ Celsius }

func (f *celciusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // No error check needed, as no switch case will match.
	switch unit {
	case "C", "ºC":
		f.Celsius = Celsius(value)
		return nil
	case "F", "ºF":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelciusFlag defines a Celcius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit (e.g. 10ºC).
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celciusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
