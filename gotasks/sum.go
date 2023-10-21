package main

import "fmt"

func main() {
	var N, sum int
	fmt.Print("Enter a positive integer N: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	for i := 2; i <= N; i += 2 {
		sum += i
	}
	fmt.Printf("The sum of even numbers from 1 to %d is %d\n", N, sum)
}
