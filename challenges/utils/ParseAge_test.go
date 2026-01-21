package utils

import "testing"

// TestParseAge_EmptyInput → check "" returns 0 and correct error
func TestParseAge_EmptyInput(t *testing.T) {
	result, err := ParseAge("")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expected := 0
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

// TestParseAge_InvalidInput → check "abc" returns 0 and correct error

func TestParseAge_InvalidInput(t *testing.T) {
	result, err := ParseAge("abc")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expected := 0

	if result != expected {
		t.Fatalf("got %v, expected %v", result, expected)
	}
}
