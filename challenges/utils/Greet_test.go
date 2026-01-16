package utils

import "testing"

// Write a proper unit test for Greet in greet_test.go.

// The test must handle empty strings as input (return an error if empty).

// Name the test correctly according to Go testing conventions.

// Don’t use table tests — you know the rule from the course.

// func Greet(name string) string {
//     return "Hello " + name
// }

func TestGreet_EmptyInput(t *testing.T) {
	greeting, err := Greet("")
	if err == nil {
		t.Errorf("expected error but got nil")
	}

	if greeting != "" {
		t.Fatalf("expected empty string but got %v", greeting)
	}

}
