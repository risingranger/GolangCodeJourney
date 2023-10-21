package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	if isPalindrome(input) {
		fmt.Println("It's a palindrome.")
	} else {
		fmt.Println("It's not a palindrome.")
	}
}
