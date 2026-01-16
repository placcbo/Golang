package utils

import "testing"

func TestCheckPositiveInt(t *testing.T) {
	result, err := CheckPositiveInt(10)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	}
	expected := true
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

// test negative number

func TestNegative(t *testing.T) {
	result, err := CheckPositiveInt(-10)
	if err == nil {
		t.Errorf("expected error, %v", err)
	}

	expected := false
	t.Fatalf("expected %v, got %v", expected, result)

}
