package main

import (
	"encoding/json"

	"fmt"
)

type Country struct {
	name string

	capital string

	gdp float64
}

func main() {

	var india Country

	Data := []byte(`

    {

        "name":"Bharat",

        "capital": "New Delhi",

        "gdp":2

    }

    `)

	err := json.Unmarshal(Data, &india)

	if err != nil {

		fmt.Println("Something went wrong")

	}

	fmt.Println(india.capital)

	fmt.Println(india.name)

}
