package utils

import (
	"errors"
	"strconv"
)

func ParseAge(input string) (age int, err error) {
	if input == "" {
		return 0, errors.New("empty string")
	}

	age, err = strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	if age < 0 {
		return 0, errors.New("negative age")
	}

	return age, nil

}
