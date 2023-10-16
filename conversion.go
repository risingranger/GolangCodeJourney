package main

import (
	"fmt"
	"length"
	"weight"
	"temperature"
)

func main() {
	// Length conversion
	fmt.Println("Length Conversion:")
	inches := 10.0
	centimeters := length.InchesToCentimeters(inches)
	fmt.Printf("%.2f inches is %.2f centimeters\n", inches, centimeters)

	centimeters = 25.4
	inches = length.CentimetersToInches(centimeters)
	fmt.Printf("%.2f centimeters is %.2f inches\n", centimeters, inches)

	// Weight conversion
	fmt.Println("\nWeight Conversion:")
	kilograms := 5.0
	pounds := weight.KilogramsToPounds(kilograms)
	fmt.Printf("%.2f kilograms is %.2f pounds\n", kilograms, pounds)

	pounds = 11.02
	kilograms = weight.PoundsToKilograms(pounds)
	fmt.Printf("%.2f pounds is %.2f kilograms\n", pounds, kilograms)

	// Temperature conversion
	fmt.Println("\nTemperature Conversion:")
	celsius := 20.0
	fahrenheit := temperature.CelsiusToFahrenheit(celsius)
	fmt.Printf("%.2f degrees Celsius is %.2f degrees Fahrenheit\n", celsius, fahrenheit)

	fahrenheit = 68.0
	celsius = temperature.FahrenheitToCelsius(fahrenheit)
	fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celsius\n", fahrenheit, celsius)
}
