package main

import "fmt"

func main() {

	num := make(chan int)

	mesg := make(chan string)

	go channelNumber(num)

	go channelMessage(mesg)

	select {

	case firstChannel := <-num:

		fmt.Println("1. Channel Data ", firstChannel)

	case secondChannel := <-mesg:

		fmt.Println("2. Channel Data ", secondChannel)

	}

}

func channelMessage(mesg chan string) {

	mesg <- "Good Evening"

}

func channelNumber(num chan int) {

	num <- 10

}
