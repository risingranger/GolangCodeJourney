package main

import (
	"fmt"

	"log"

	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "Good Evening")

	//"Good Evening"

}

func main() {

	http.HandleFunc("/greeting", HomePage)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
