package utils

import "testing"

func TestAverage_Valid(t *testing.T) {
	avg, err := Average("10", "20")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if avg != 15 {
		t.Fatalf("expected 15, got %d", avg)
	}
}

func TestAverage_InvalidFirst(t *testing.T) {
	_, err := Average("x", "10")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestAverage_InvalidSecond(t *testing.T) {
	_, err := Average("10", "-1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
