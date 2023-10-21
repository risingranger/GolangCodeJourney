package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "sayHello":
		fmt.Println("Hello, World!")
	case "concat":
		if len(os.Args) != 4 {
			fmt.Println("Usage: go run main.go concat <num1> <num2>")
			return
		}
		num1 := os.Args[2]
		num2 := os.Args[3]
		sum := num1 + num2
		fmt.Printf("concat: %s\n", sum)
	default:
		fmt.Println("Unknown command:", command)
	}
}
