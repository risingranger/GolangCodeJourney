// the pointer variables to store the memory address.
package main

import "fmt"

func main() {
	var name = "Dr Tarkeshwar Barua"
	var ptr *string
	// assign the memory address of name to the pointer
	ptr = &name
	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Address of the variable", &name)
}
