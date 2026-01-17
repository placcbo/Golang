package utils

import (
	"errors"
	"strconv"
)

func ParseAge(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")

	}
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return age, nil

}

// Requirements:

// Converts input to an integer

// Returns an error "input cannot be empty" if the input is ""

// Returns an error if input cannot be converted to int

// Returns the parsed int if successful
