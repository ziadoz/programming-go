// Converts numbers from one unit of measurement to another (e.g. Celcius to Farenheit, Feet to Metres etc.)
package conv

import "fmt"

type Celcius float64
type Farenheit float64

type Feet float64
type Metre float64

type Pound float64
type Kilogram float64

// Pretty formatting of each type.
func (c Celcius) String() string   { return fmt.Sprintf("%gºC", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%gºF", f) }
func (f Feet) String() string      { return fmt.Sprintf("%g feet", f) }
func (m Metre) String() string     { return fmt.Sprintf("%g metres", m) }
func (p Pound) String() string     { return fmt.Sprintf("%g lbs", p) }
func (k Kilogram) String() string  { return fmt.Sprintf("%g kg", k) }

// Methods to convert to/from applicable types.
func CToF(c Celcius) Farenheit { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func FToM(f Feet) Metre { return Metre(f / 3.28) }
func MToF(m Metre) Feet { return Feet(m * 0.3048) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }
