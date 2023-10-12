package main

import "fmt"

func abc() {
	fmt.Print("Enter Value of a: ")
	var a int
	fmt.Scanln(&a)
	fmt.Print("Enter Value of b: ")
	var b int
	fmt.Scanln(&b)
	var c int
	c = a + b
	fmt.Print(c)
}
