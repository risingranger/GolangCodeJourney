package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		return
	}
	fmt.Println(bs)
	fmt.Println(string(bs))

}
