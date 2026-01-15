package utils

import "testing"

func TestSafeDivide_ValidDivision(t *testing.T) {
	result, err := SafeDivide(10, 5)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if result != 5 {
		t.Fatalf("expected 5, got %d", result)
	}

}

func TestSafeDivide_DivisionWithRemainder(t *testing.T) {
	result, err := SafeDivide(7, 2)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if result != 3 {
		t.Fatalf("expected 3 from 7/2, got %d", result)
	}
}
func TestSafeDivide_ZeroNumerator(t *testing.T) {
	result, err := SafeDivide(0, 10)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if result != 0 {
		t.Fatalf("unexpected 0 from 0/10 %d", result)

	}
}

func TestSafeDivide_NegativeNumbers(t *testing.T) {
	result, err := SafeDivide(10, -2)
	if err != nil {
		t.Fatalf("unexoected error %v", err)
	}
	if result != -5 {
		t.Fatalf("unexpected -5 from 10/-2 %d", result)

	}
}

func TestSafeDivide_DivideByZero(t *testing.T) {
	result, err := SafeDivide(10, 0)

	if err == nil {
		t.Fatal("expected an error when dividing by zero, got nil")
	}

	if result != 0 {
		t.Fatalf("expected 0 from 10/0, got %d", result)
	}

}
