package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Print("Enter the start of the range: ")
	_, err := fmt.Scanf("%d", &start)
	if err != nil {
		fmt.Println("Invalid input for the start of the range.")
		return
	}

	fmt.Print("Enter the end of the range: ")
	_, err = fmt.Scanf("%d", &end)
	if err != nil {
		fmt.Println("Invalid input for the end of the range.")
		return
	}

	fmt.Printf("Even numbers in the range from %d to %d:\n", start, end)

	for num := start; num <= end; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
