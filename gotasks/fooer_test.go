package main

import "testing"

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(2)
	t.Logf("Your input is %d", input)
	if result != "more" {
		t.Error("Found More not less  and expected ")
	}
	t.Fatalf("Stop the test now, we have seen enough")
	t.Error(" Nothing to display any more ")
}
