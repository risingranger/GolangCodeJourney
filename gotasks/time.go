package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	after2 := now.Add(2)
	p(after2)
	date1 := time.Date(2023, 10, 20, 10, 20, 20, 4, time.UTC)
	p(date1)
}
