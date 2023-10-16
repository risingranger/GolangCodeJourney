package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Convert length (inches to centimeters)")
		fmt.Println("2. Convert length (centimeters to inches)")
		fmt.Println("3. Convert weight (kilograms to pounds)")
		fmt.Println("4. Convert weight (pounds to kilograms)")
		fmt.Println("5. Convert temperature (Celsius to Fahrenheit)")
		fmt.Println("6. Convert temperature (Fahrenheit to Celsius)")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 7 {
			fmt.Println("Exiting program.")
			break
		}

		switch choice {
		case 1:
			var inches float64
			fmt.Print("Enter length in inches: ")
			fmt.Scanln(&inches)
			centimeters := inches * 2.54
			fmt.Printf("%.2f inches is %.2f centimeters\n", inches, centimeters)

		case 2:
			var centimeters float64
			fmt.Print("Enter length in centimeters: ")
			fmt.Scanln(&centimeters)
			inches := centimeters / 2.54
			fmt.Printf("%.2f centimeters is %.2f inches\n", centimeters, inches)

		case 3:
			var kilograms float64
			fmt.Print("Enter weight in kilograms: ")
			fmt.Scanln(&kilograms)
			pounds := kilograms * 2.20462
			fmt.Printf("%.2f kilograms is %.2f pounds\n", kilograms, pounds)

		case 4:
			var pounds float64
			fmt.Print("Enter weight in pounds: ")
			fmt.Scanln(&pounds)
			kilograms := pounds / 2.20462
			fmt.Printf("%.2f pounds is %.2f kilograms\n", pounds, kilograms)

		case 5:
			var celsius float64
			fmt.Print("Enter temperature in Celsius: ")
			fmt.Scanln(&celsius)
			fahrenheit := (celsius * 9 / 5) + 32
			fmt.Printf("%.2f degrees Celsius is %.2f degrees Fahrenheit\n", celsius, fahrenheit)

		case 6:
			var fahrenheit float64
			fmt.Print("Enter temperature in Fahrenheit: ")
			fmt.Scanln(&fahrenheit)
			celsius := (fahrenheit - 32) * 5 / 9
			fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celsius\n", fahrenheit, celsius)

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

