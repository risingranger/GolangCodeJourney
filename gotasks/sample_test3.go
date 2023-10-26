package main

import "testing"

func IntMin(num1, num2 int) int {

	if num1 > num2 {

		return num1

	} else {

		return num2

	}

}

func TestIntMin(t *testing.T) {

	result := IntMin(2, -1)

	expected := -1

	if result != expected {

		t.Errorf("IntMin(2,-1) PASS expected %d but result %d", expected, result)

	}

}
