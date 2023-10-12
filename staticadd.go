package main // creating package
import "fmt" // format text library
func mnc() { // Gateway to start your application
	fmt.Println("Please Input value of A ") // printing statement
	//var a int =10
	//var b int =20
	a := 30

	b := 10

	var c int

	//fmt.Scanln(&a)

	fmt.Println("Please Input value of B ")

	//fmt.Scanln(&b)

	c = a + b

	fmt.Println("Addition of ", a, " and ", b, " is ", c)

}
