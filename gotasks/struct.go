package main

import "fmt"

// create struct
type Weather struct {
	city        string
	temperature int
}

func main() {
	// create struct variable
	weather := Weather{"New Delhi", 35}
	fmt.Println("Initial Weather:", weather)
	// create struct type pointer
	ptr := &weather
	// change value of temperature to 25
	ptr.temperature = 37
	fmt.Println("Updated Weather:", weather)
}
