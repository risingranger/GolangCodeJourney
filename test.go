package main

import "fmt"

func main() {
	fmt.Print("Enter Value of a: ")
	var a int
	fmt.Scanln(&a)
	fmt.Print("Enter Value of b: ")
	var b int
	fmt.Scanln(&b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After Swapping", a, b)
}
