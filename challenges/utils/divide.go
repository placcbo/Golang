package utils

import "errors"

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	division := a / b
	return float64(division), nil
}
