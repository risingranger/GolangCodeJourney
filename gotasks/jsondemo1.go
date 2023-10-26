package main

import (
	"encoding/json"

	"fmt"
)

type Country1 struct {
	Name string

	Capital string

	Gdp float64
}

func main() {

	var country []Country1

	// JSON Storing object // Storing list of object

	Data := []byte(`

    [

        {"Name":"INDIA", "Capital":"New Delhi"},

        {"Name":"USA", "Capital":"New York"}

    ]

    `)

	err := json.Unmarshal(Data, &country)

	if err != nil {

		fmt.Println("Something Went Wrong", err)

	}

	for obj := range country {

		fmt.Println(country[obj].Name)

		fmt.Println(country[obj].Capital)

	}

}
