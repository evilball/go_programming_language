//Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that
//reads numbers from its command-line arguments or from the standard input if there
//are no arguments, and converts each number into units like temperature in Celsius
//and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the
//like.

package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/evilball/go_programming_language/ch2/ex/ex2.2/convert"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			makeConversions(arg)
		}
	} else {
		showInitMessage()
	}

	exit := false
	for exit != true {
		var arg string
		_, err := fmt.Scanln(&arg)
		if err != nil {
			if (err.Error() == "unexpected newline") {
				continue
			} else {
				fmt.Fprintf(os.Stderr, "ex2.1: %v\n", err)
				os.Exit(1)
			}
		}
		if arg != "exit" {
			makeConversions(arg)
		} else {
			exit = true
		}
	}
}
func makeConversions(arg string) {
	number, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Println("Please input correct number or type 'exit' for exit")
		return
	}

	fmt.Println("======================")

	fahrenheit := convert.Fahrenheit(number)
	celsius := convert.Celsius(number)
	kelvin := convert.Kelvin(number)
	fmt.Println("Temperature conversion")
	fmt.Println("----------------------")
	fmt.Printf("%s = %s\n", celsius, convert.CelsiusToFahrenheit(celsius))
	fmt.Printf("%s = %s\n", celsius, convert.CelsiusToKelvin(celsius))
	fmt.Printf("%s = %s\n", fahrenheit, convert.FahrenheitToCelsius(fahrenheit))
	fmt.Printf("%s = %s\n", fahrenheit, convert.FahrenheitToKelvin(fahrenheit))
	fmt.Printf("%s = %s\n", kelvin, convert.KelvinToFahrenheit(kelvin))
	fmt.Printf("%s = %s\n", kelvin, convert.KelvinToCelsius(kelvin))

	foot := convert.Foot(number)
	meter := convert.Meter(number)
	fmt.Println()
	fmt.Println("Length conversion")
	fmt.Println("----------------------")
	fmt.Printf("%s = %s\n", meter, convert.MeterToFoot(meter))
	fmt.Printf("%s = %s\n", foot, convert.FootToMeter(foot))

	pound := convert.Pound(number)
	kilogram := convert.Kilogram(number)
	fmt.Println()
	fmt.Println("Weight conversion")
	fmt.Println("----------------------")
	fmt.Printf("%s = %s\n", pound, convert.PoundToKilogram(pound))
	fmt.Printf("%s = %s\n", kilogram, convert.KilogramToPound(kilogram))
	fmt.Println()
	fmt.Println("======================")
	showInitMessage()
}
func showInitMessage() {
	fmt.Println("Input any number or type 'exit' for exit")
}

