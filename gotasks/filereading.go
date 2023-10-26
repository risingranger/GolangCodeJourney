package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Something went wrong ", err)
		return
	}
	fmt.Println(string(data))
}
