package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(273.15 + c) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func FToK(f Fahrenheit) Kelvin { return Kelvin(273.15 + FToC(f)) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
