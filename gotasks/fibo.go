package main

import "fmt"

func generateFibonacci(n int) []int {
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	var N int
	fmt.Print("Enter the number of Fibonacci numbers to generate: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	fibonacci := generateFibonacci(N)
	fmt.Println("Fibonacci Sequence:")
	fmt.Println(fibonacci)
}
