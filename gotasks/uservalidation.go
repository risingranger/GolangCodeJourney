package main

import (
	"strings"

	"testing"
)

func TestUsername(t *testing.T) {

	t.Run("withoutDot", func(t *testing.T) {

		username := ExactUsername("tbarua1@gmail.com")

		if username != "r" {

			t.Fatalf("Got : %v", username)

		}

	})

	t.Run("withDot", func(t *testing.T) {

		username := ExactUsername("tbarua1@rediffmail.com")

		if username != "tbarua1" {

			t.Fatalf("Got : %v", username)

		}

	})

}

func ExactUsername(s string) string {

	at := strings.Index(s, "@")

	return s[:at] // tbarua1

}
