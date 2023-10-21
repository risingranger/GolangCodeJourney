package main

import "fmt"

func main() {
	name := "Dr Tarkeshwar Barua"
	message := greeting(name)
	fmt.Println(message)
}
func greeting(name string) *string {
	msg := name + "Good Evening"
	return &msg
}
