package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Dr Tarkeshwar Barua"
	salary := 100
	paid := false
	room := 3.5
	displayValue(name)
	displayValue(salary)
	displayValue(paid)
	displayValue(room)
}

func displayValue(i interface{}) {
	fmt.Println(reflect.TypeOf(i), "Data Type", i, " Value")
}
