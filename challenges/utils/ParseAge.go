package utils

import (
	"errors"
	"strconv"
)

func ParseAge(input string) (int, error) {

	age, err := strconv.Atoi(input)
	if input == "" {
		return 0, errors.New("age is required")
	}

	if err != nil {
		return 0, errors.New("age must be a number")
	}
	//validate age

	if age > 120 {
		return 0, errors.New("age must be less than or equal to 120")
	}
	return age, nil

}

// Create a function:

// func ParseAge(input string) (int, error)

// Rules

// input is a string (example: "23")

// Convert it to an integer

// Valid age range: 1 to 120

// Behavior
// Input	Expected Result
// "25"	25, nil
// "0"	error
// "121"	error
// "-5"	error
// "abc"	error
// ""	error
