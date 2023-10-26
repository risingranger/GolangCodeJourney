package main

import (
	"fmt"

	"os"
)

func main() {

	fileread, err := os.Open("hello.txt")

	if err != nil {

		return

	}

	fmt.Println("File has been opened successfully")

	defer fileread.Close()

	stat, err := fileread.Stat()

	if err != nil {

		return

	}

	fmt.Println("stat of file ", stat)

	bs := make([]byte, stat.Size())

	data, err := fileread.Read(bs)

	fmt.Println(data)

	if err != nil {

		return

	}

	str := string(bs)

	fmt.Println(str)

}
