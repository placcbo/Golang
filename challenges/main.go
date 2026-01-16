package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func ParseInt(input string) (int, error) {
	result, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	result, err := ParseInt("100H")
	fmt.Println(result, err)

	results, err := Greet("kevin")
	fmt.Println(results, err)

}

// Write a function with named return values that takes a name string and returns a greeting string and an error. The error should be returned if the name is empty. Then write one test case for the error scenario.

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}
	greetings := "hello " + name

	return greetings, nil
}
func TestGreet_emptyString(t *testing.T) {
	result, err := Greet("")
	if err == nil {
		t.Fatalf("expected an error for empty name, got nil")
	}

	expected := ""

	if result != expected {
		t.Fatalf("expected result to be empty string, got '%s'", result)
	}
}
