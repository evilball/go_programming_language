package convert

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Foot float64
type Meter float64

type Pound float64
type Kilogram float64

func KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k + Kelvin(-273.15))
}

func CelsiusToKelvin(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KelvinToFahrenheit(k Kelvin) Fahrenheit {
	return CelsiusToFahrenheit(KelvinToCelsius(k))
}

func FahrenheitToKelvin(f Fahrenheit) Kelvin {
	return CelsiusToKelvin(FahrenheitToCelsius(f))
}

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FootToMeter(f Foot) Meter {
	return Meter(f * 0.3048)
}

func MeterToFoot(m Meter) Foot {
	return Foot(m * 3.28084)
}

func PoundToKilogram(p Pound) Kilogram {
	return Kilogram(p * 0.453592)
}

func KilogramToPound(k Kilogram) Pound {
	return Pound(k * 2.20462)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}