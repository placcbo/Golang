package utils

import "testing"

func TestParseAge_valid(t *testing.T) {
	age, err := ParseAge("25")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if age != 25 {
		t.Fatalf("expected 25, got %d", age)
	}
}

func TestParseAge_Zero(t *testing.T) {
	age, err := ParseAge("0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)

	}

	if age != 0 {
		t.Fatalf("expected 0, got %d", age)

	}
}

func TestParseAge_EmptyInput(t *testing.T) {
	_, err := ParseAge("")
	if err != nil {
		t.Fatal("expected error, got nil")
	}
}
