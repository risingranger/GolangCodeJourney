package main

import (
	"conversion/length"
	"conversion/temperature"
	"conversion/weight"
	"fmt"
)

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Length Conversion")
		fmt.Println("2. Weight Conversion")
		fmt.Println("3. Temperature Conversion")
		fmt.Println("4. Quit")
		var choice int
		fmt.Scanln(&choice)

		if choice == 4 {
			fmt.Println("Goodbye!")
			break
		}

		switch choice {
		case 1:
			lengthConversion()
		case 2:
			weightConversion()
		case 3:
			temperatureConversion()
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func lengthConversion() {
	fmt.Println("Choose a length conversion:")
	fmt.Println("1. Inches to Centimeters")
	fmt.Println("2. Centimeters to Inches")
	var subChoice int
	fmt.Scanln(&subChoice)

	switch subChoice {
	case 1:
		var inches float64
		fmt.Print("Enter length in inches: ")
		fmt.Scanln(&inches)
		centimeters := length.InchesToCentimeters(inches)
		fmt.Printf("%.2f inches is %.2f centimeters\n", inches, centimeters)
	case 2:
		var centimeters float64
		fmt.Print("Enter length in centimeters: ")
		fmt.Scanln(&centimeters)
		inches := length.CentimetersToInches(centimeters)
		fmt.Printf("%.2f centimeters is %.2f inches\n", centimeters, inches)
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}

func weightConversion() {
	fmt.Println("Choose a weight conversion:")
	fmt.Println("1. Kilogram to Pound")
	fmt.Println("2. Pound to Kilogram")
	var subChoice int
	fmt.Scanln(&subChoice)

	switch subChoice {
	case 1:
		var kilograms float64
		fmt.Print("Enter weight in kilograms: ")
		fmt.Scanln(&kilograms)
		pounds := weight.KilogramsToPounds(kilograms)
		fmt.Printf("%.2f kilograms is %.2f pounds\n", kilograms, pounds)
	case 2:
		var pounds float64
		fmt.Print("Enter weight in pounds: ")
		fmt.Scanln(&pounds)
		kilograms := weight.PoundsToKilograms(pounds)
		fmt.Printf("%.2f pounds is %.2f kilograms\n", pounds, kilograms)
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}

func temperatureConversion() {
	fmt.Println("Choose a temperature conversion:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	var subChoice int
	fmt.Scanln(&subChoice)

	switch subChoice {
	case 1:
		var celsius float64
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scanln(&celsius)
		fahrenheit := temperature.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f degrees Celsius is %.2f degrees Fahrenheit\n", celsius, fahrenheit)
	case 2:
		var fahrenheit float64
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scanln(&fahrenheit)
		celsius := temperature.FahrenheitToCelsius(fahrenheit)
		fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celsius\n", fahrenheit, celsius)
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}
