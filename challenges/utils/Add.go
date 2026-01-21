package utils

import "errors"

func Add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("int cannot be negative")
	}

	return a + b, nil
}
