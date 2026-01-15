package utils

import "testing"

func TestTpt_validTpt(t *testing.T) {
	result, err := Tpt(10, 10, 20)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := 40

	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
