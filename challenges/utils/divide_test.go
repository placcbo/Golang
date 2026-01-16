package utils

import "testing"

//valid division

func TestDivide_validDivision(t *testing.T) {
	result, err := Divide(10, 4)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := 2.5

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)

	}

}

// division be Zero

func TestDivide_DivisionByZero(t *testing.T) {
	_, err := Divide(27, 0)
	if err == nil {
		t.Fatalf("expected error but did not get any error, got %v", err)
	}

}
