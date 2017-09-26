package tempconv

// Celcius Conversions
// CToF converts a Celius temperature to Farenheit.
func CToF(c Celcius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

// CToK converts a Celcius temperature to Kelvin.
func CToK(c Celcius) Kelvin {
	return Kelvin(c + 273.15)
}

// Farenheit Conversions
// FToC converts a Farenheit temperature to Celcius.
func FToC(f Farenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

// FToK converts a Farenheit temperature to Kelvin.
func FToK(f Farenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

// Kelvin Conversions
// KToC converts a Kevlin temperature to Celcius.
func KToC(k Kelvin) Celcius {
	return Celcius(k - 273.15)
}

// KToF converts a Kevlin temperature to Farenheit.
func KToF(k Kelvin) Farenheit {
	return Farenheit((k-273.15)*9/5 + 32)
}
