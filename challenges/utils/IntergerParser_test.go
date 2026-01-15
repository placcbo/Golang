package utils

import (
	"testing"
)

func TestParseStrictInt_ValidInteger(t *testing.T) {
	result, err := IntergerParser("10")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if result != 10 {
		t.Fatalf("expected 10, got %d", result)
	}

}
func TestParseStrictInt_Zero(t *testing.T) {

	result, err := IntergerParser("0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 0 {
		t.Fatalf("expected 0, got %d", result)
	}

}

func TestParseStrictInt_NegativeInteger(t *testing.T) {
	result, err := IntergerParser("-10")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != -10 {
		t.Errorf("expected -10, got %d", result)
	}
}

func TestParseStrictInt_InvalidString(t *testing.T) {
	_, err := IntergerParser("abc")

	if err == nil {
		t.Fatalf("expected error for invalid input, got nil")
	}

}
