package utils

import (
	"strconv"
)

func IntergerParser(input string) (int, error) {
	result, err := strconv.Atoi("151")

	if err != nil {
		return 0, nil
	}

	return result, nil
}
