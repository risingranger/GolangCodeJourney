package main

import "fmt"

func main() {
	num := 55
	fmt.Println("Value of num : ", num)
	fmt.Println("Address of num : ", &num)
	updateValue(&num)
	fmt.Println("Value of num : ", num)
	fmt.Println("Address of num : ", &num)
}
func updateValue(num *int) {
	*num = 50
}
