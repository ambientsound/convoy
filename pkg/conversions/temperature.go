package conversions

const celsiusToFahrenheitFraction = float64(9) / 5
const kelvinCelsiusZeroTemperature = 273.15
const unitCelsius = "°C"
const unitFahrenheit = "°F"
const unitKelvin = "°K"

var Temperatures = []Converter{
	CelsiusFahrenheit,
	CelsiusKelvin,
	FahrenheitCelsius,
	FahrenheitKelvin,
	KelvinCelsius,
	KelvinFahrenheit,
}

func CelsiusFahrenheit(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitCelsius,
		},
		Output: Amount{
			Value: (in.Value * celsiusToFahrenheitFraction) + 32,
			Unit:  unitFahrenheit,
		},
	}
}

func CelsiusKelvin(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitCelsius,
		},
		Output: Amount{
			Value: in.Value + kelvinCelsiusZeroTemperature,
			Unit:  unitKelvin,
		},
	}
}

func FahrenheitCelsius(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitFahrenheit,
		},
		Output: Amount{
			Value: (in.Value - 32) / celsiusToFahrenheitFraction,
			Unit:  unitCelsius,
		},
	}
}

func FahrenheitKelvin(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitFahrenheit,
		},
		Output: Amount{
			Value: ((in.Value - 32) / celsiusToFahrenheitFraction) + kelvinCelsiusZeroTemperature,
			Unit:  unitKelvin,
		},
	}
}

func KelvinCelsius(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitKelvin,
		},
		Output: Amount{
			Value: kelvinCelsiusZeroTemperature - in.Value,
			Unit:  unitCelsius,
		},
	}
}

func KelvinFahrenheit(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitKelvin,
		},
		Output: Amount{
			Value: ((in.Value - kelvinCelsiusZeroTemperature) * celsiusToFahrenheitFraction) + 32,
			Unit:  unitFahrenheit,
		},
	}
}
