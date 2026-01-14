package utils

import (
	"errors"
	"testing"
)

func TestParseAgeStrict_EmptyInput(t *testing.T) {
	_, err := ParseAgeStrict("")
	if !errors.Is(err, ErrEmptyInput) {
		t.Fatalf("expected ErrEmptyInput, got %v", err)
	}
}

func TestParseAgeStrict_InvalidNumber(t *testing.T) {
	_, err := ParseAgeStrict("abc")
	if !errors.Is(err, ErrInvalidNumber) {
		t.Fatalf("expected ErrInvalidNumber, got %v", err)
	}
}

func TestParseAgeStrict_Negative(t *testing.T) {
	_, err := ParseAgeStrict("-5")
	if !errors.Is(err, ErrInvalidNumber) {
		t.Fatalf("expected ErrInvalidNumber, got %v", err)
	}
}

func TestParseAgeStrict_Valid(t *testing.T) {
	age, err := ParseAgeStrict("18")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if age != 18 {
		t.Fatalf("expected 18, got %d", age)
	}
}
