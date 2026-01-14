package utils

import (
	"errors"
	"strconv"
)

// Sentinel Errors

var ErrEmptyInput = errors.New("empty input")
var ErrInvalidNumber = errors.New("invalid number")

func ParseAgeStrict(input string) (int, error) {
	if input == "" {
		return 0, ErrEmptyInput
	}

	age, err := strconv.Atoi(input)
	if err != nil || age < 0 {
		return 0, ErrInvalidNumber
	}

	return age, nil
}
