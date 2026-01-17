package utils

import "testing"

// testGreet_ValidName
func TestGreet(t *testing.T) {
	name, err := Greet("kevin")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	expected := "hello kevin"

	if name != expected {
		t.Fatalf("expected %v, got %v", expected, name)
	}

}

// TestGreet_EmptyName

func TestGreet_EmptyName(t *testing.T) {
	result, err := Greet("")
	if err == nil {
		t.Fatalf("expected an error, got %v", err)
	}

	if result != "" {
		t.Fatalf("expected empty string, got %v", result)
	}
}

// Also write two test functions:

// TestGreet_ValidName → checks that Greet("Kevin") returns "hello Kevin" and nil error

// TestGreet_EmptyName → checks that Greet("") returns "" and the correct error

// Rules:

// Must check both return value and error

// Test function names must be descriptive

// Use errors.New() for errors
