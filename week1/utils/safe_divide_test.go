package utils

import "testing"

func TestSafeDivide_ValidDivision(t *testing.T) {
	result, err := SafeDivide(10, 5)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if result != 2 {
		t.Errorf("expected 5. got %d", result)
	}

}

func TestSafeDivide_ZeroNumerator(t *testing.T) {
	result, err := SafeDivide(10, 0)

	if err != nil {
		t.Error("unexpected result %v", err)
	}
	if result != 0 {
		t.Fatalf("expected 0, got %d", result)
	}

}

func TestSafeDivide_NegativeNumerator(t *testing.T) {
	result, err := SafeDivide(-10, 2)

	if err == nil {
		t.Fatalf("unexpected error %v", err)
	}

	if result != -5 {
		t.Fatalf("expected -5, got %d", result)
	}
}
