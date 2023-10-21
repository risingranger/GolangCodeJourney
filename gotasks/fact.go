package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Print("Enter a non-negative integer: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	if n < 0 {
		fmt.Println("Please enter a non-negative integer.")
		return
	}

	result := factorial(n)
	fmt.Printf("The factorial of %d is %d\n", n, result)
}
