package week2

import "errors"

func SafeDivide(a, b int) (int, error) {
	if num := b; num == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil

}

// Rules

// Return the result of a / b.

// If b == 0, return 0 and an error "division by zero".

// Use an if with a short statement at least once (Week 2 concept).

// Do not panic.

// Examples:
