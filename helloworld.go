package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(os.Getenv("USER"), "Lets be Friends!") //Read linux $USER Environment
}