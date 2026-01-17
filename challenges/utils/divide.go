package utils

import "errors"

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil

}

// Requirements:

// Returns a / b if b != 0

// Returns an error "division by zero" if b == 0
