package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil || name == "" {
		name = "World"
	}
	fmt.Printf("Hello, %s!\n", name)
}
