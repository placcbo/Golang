package utils

import "testing"

// TestSafeDivide_ValidDivision → check SafeDivide(10, 2) returns 5 and nil error

func TestSafeDivide_ValidDivision(t *testing.T) {
	result, err := SafeDivide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := 5

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

// TestSafeDivide_ValidDivision → check SafeDivide(10, 2) returns 5 and nil error
func TestSafeDivide_DivisionByZero(t *testing.T) {
	result, err := SafeDivide(10, 0)
	if err == nil {
		t.Fatalf("expected an error, found nil")
	}
	if result != 0 {
		t.Fatalf("expected 0, got %v", result)
	}
}
