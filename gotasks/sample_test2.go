package main

import "testing"

func TestFooer(t *testing.T) {

	result := Fooer(2)

	if result != "more" {

		t.Error("Found More not less")

	}

}
