package main

import "fmt"

func main() {
	// new method will create variable to store pointer
	xPtr := new(int) // xPtr is an ordinary variable to store value
	one(xPtr)
	fmt.Println(xPtr)
	fmt.Println(&xPtr)
	fmt.Println(*xPtr)
}
func one(xPtr *int) {
	*xPtr = 10
}
