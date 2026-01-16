package utils

import "errors"

func CheckPositiveInt(number int) (bool, error) {
	if number < 0 {
		return false, errors.New("number is Negative")
	}

	return true, nil
}
