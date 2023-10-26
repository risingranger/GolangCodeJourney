package main

import "fmt"

func main() {

	age := -3
	error := fmt.Errorf("Invalid age %d", age)
	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("You are ", age, " years old")
	}

}
