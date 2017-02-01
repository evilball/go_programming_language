//Exercise 2.1: Add types, constants, and functions to tempconv for processing
//temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of
//1K has the same magnitude as 1°C.

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KtoF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func FtoK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
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