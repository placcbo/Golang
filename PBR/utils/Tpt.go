package utils

import "errors"

func Tpt(a, b, c int) (int, error) {
	workerTotalTpt := a + b + c
	averageTpt := workerTotalTpt / 3
	if workerTotalTpt == 0 {
		return 0, errors.New("total cant be 0")
	}
	return averageTpt, nil
}
