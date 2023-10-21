package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan string = make(chan string)

	go pinger(c)

	go pinger1(c)

	go pinger(c)

	var input string

	fmt.Scanln((&input))

}

func pinger(c chan string) {

	for i := 0; true; i++ {

		c <- "Ping"

	}

}

func pinger1(c chan string) {

	for {

		c <- "Ping"

		fmt.Println(c)

		time.Sleep(time.Second * 5)

	}

}
