package utils

import "errors"

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}
	return "hello " + name, nil

}

// Write a Go function Greet(name string) (string, error) that:

// Returns "hello <name>" if name is not empty

// Returns an error "Name cannot be empty" if name is empty
