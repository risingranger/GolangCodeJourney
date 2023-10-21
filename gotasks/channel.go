package main

import (
	"fmt"

	"time"
)

func main() {

	ch := make(chan int)

	go sayGoodEvening(ch)

	x := <-ch

	fmt.Println("Main Function ", x)

}

func sayGoodEvening(ch chan int) {

	time.Sleep(5 * time.Second)

	fmt.Println("Good Evening")

	ch <- 1234

}
