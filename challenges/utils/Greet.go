package utils

import "errors"

// Write a proper unit test for Greet in greet_test.go.

// The test must handle empty strings as input (return an error if empty).

// Name the test correctly according to Go testing conventions.

// Don’t use table tests — you know the rule from the course.

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty string")
	}
	greeting := "hello " + name

	return greeting, nil
}
