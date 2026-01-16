package utils

import (
	"testing"
)

// Normal tes
func TestScore_normalValues(t *testing.T) {
	result, err := Score(67, 90.0)
	if err != nil {
		t.Fatalf("unexpexteded error, %v", err)
	}
	expexted := 2
	if result != expexted {
		t.Errorf("got %d, expexted %d", result, expexted)
	}
}
