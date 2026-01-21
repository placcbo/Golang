package utils

import "testing"

//TestGreet_ValidName → check Greet("Kevin") returns "hello Kevin" and nil error

func TestGreet_ValidName(t *testing.T) {
	result, err := Greet("Kevin")

	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	expected := "hello Kevin"
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)

	}

}

//  TestGreet_EmptyName → check Greet("") returns "" and error "Name cannot be empty"

func TestGreet_EmptyName(t *testing.T) {
	result, err := Greet("")
	if err == nil {
		t.Errorf("Name cannot be empty")
	}
	expected := ""

	if result != expected {
		t.Errorf("expected % v, got %v", expected, result)
	}

}
