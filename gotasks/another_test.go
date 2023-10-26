package greeting

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Tarkeshwar"
	want := regexp.MustCompile("\b" + name + "\b")
	if name != "" {
		t.Fatalf("Test case failed")
	}
}
