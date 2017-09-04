package tempconv

// CToF converts a Celius temperature to Farenheit.
func CToF(c Celcius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

// FToC converts a Farenheit temperature to Celcius.
func FToC(f Farenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
