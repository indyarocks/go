package tempconv

func FtoC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * (5.0 / 9))
}

func CtoF(c Celcius) Fahrenheit {
	return Fahrenheit(32 + (c * 9 / 5))
}
