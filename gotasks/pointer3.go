package main

import "fmt"

func main() {

	y := 434

	var pointers *int

	pointers = &y

	//pointers := &y

	fmt.Println("Value of y ", y)

	fmt.Println("Address of  ", pointers)
}
