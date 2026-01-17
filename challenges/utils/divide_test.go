package utils

import (
	"testing"
)

func TestSafeDivide_ValidDivision(t *testing.T) {
	result, err := SafeDivide(10, 2)

	if err != nil {
		t.Fatalf("expected an error, got nil")
	}

	expected := 5
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

// TestSafeDivide_DivideByZero
func TestSafeDivide_DivideByZero(t *testing.T) {
	result, err := SafeDivide(10, 0)
	if err == nil {
		t.Fatalf("expected error %v but got ", err)
	}
	expected := 0
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}

}
