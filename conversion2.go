package main

import (
	"fmt"
	"length"
	"weight"
	"temperature"
)

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Length Conversion")
		fmt.Println("2. Weight Conversion")
		fmt.Println("3. Temperature Conversion")
		fmt.Println("4. Quit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			lengthConversion()
		case 2:
			weightConversion()
		case 3:
			temperatureConversion()
		case 4:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

func lengthConversion() {
	fmt.Println("Length Conversion:")
	fmt.Println("1. Inches to Centimeters")
	fmt.Println("2. Centimeters to Inches")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter inches: ")
		var inches float64
		fmt.Scan(&inches)
		centimeters := length.InchesToCentimeters(inches)
		fmt.Printf("%.2f inches is %.2f centimeters\n", inches, centimeters)
	case 2:
		fmt.Print("Enter centimeters: ")
		var centimeters float64
		fmt.Scan(&centimeters)
		inches := length.CentimetersToInches(centimeters)
		fmt.Printf("%.2f centimeters is %.2f inches\n", centimeters, inches)
	default:
		fmt.Println("Invalid choice. Please enter a valid option.")
	}
}

func weightConversion() {
	fmt.Println("Weight Conversion:")
	fmt.Println("1. Kilograms to Pounds")
	fmt.Println("2. Pounds to Kilograms")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter kilograms: ")
		var kilograms float64
		fmt.Scan(&kilograms)
		pounds := weight.KilogramsToPounds(kilograms)
		fmt.Printf("%.2f kilograms is %.2f pounds\n", kilograms, pounds)
	case 2:
		fmt.Print("Enter pounds: ")
		var pounds float64
		fmt.Scan(&pounds)
		kilograms := weight.PoundsToKilograms(pounds)
		fmt.Printf("%.2f pounds is %.2f kilograms\n", pounds, kilograms)
	default:
		fmt.Println("Invalid choice. Please enter a valid option.")
	}
}

func temperatureConversion() {
	fmt.Println("Temperature Conversion:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter degrees Celsius: ")
		var celsius float64
		fmt.Scan(&celsius)
		fahrenheit := temperature.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f degrees Celsius is %.2f degrees Fahrenheit\n", celsius, fahrenheit)
	case 2:
		fmt.Print("Enter degrees Fahrenheit: ")
		var fahrenheit float64
		fmt.Scan(&fahrenheit)
		celsius := temperature.FahrenheitToCelsius(fahrenheit)
		fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celsius\n", fahrenheit, celsius)
	default:
		fmt.Println("Invalid choice. Please enter a valid option.")
	}
}
