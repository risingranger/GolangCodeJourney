package main

import "fmt"

func main() {

	type Person struct {
		name   string
		salary int
	}

	tarkesh := Person{"Dr Tarkeshwar Barua", 100}
	var ptr *Person
	ptr = &tarkesh
	ptr1 := &tarkesh
	fmt.Println(ptr)
	fmt.Println(ptr.name)
	fmt.Println(ptr.salary)
	fmt.Println(ptr1)
	fmt.Println(ptr1.name)
	fmt.Println(ptr1.salary)
	fmt.Println(tarkesh)
	fmt.Println(tarkesh.name)
	fmt.Println(tarkesh.salary)

}
