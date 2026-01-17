package utils

import "errors"

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("denominator cannot be zero")
	}
	return (a / b), nil
}

// 🔥 Exercise 4 — Divide Function with Error

// Task:

// Write a Go function:

// func SafeDivide(a, b int) (int, error)

// Requirements:

// Returns a / b if b != 0

// Returns an error "division by zero" if b == 0
