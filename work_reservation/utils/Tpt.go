package utils

import "errors"

func Tpt(a, b, c int) (int, error) {
	workerTotalTpt := a + b + c
	if workerTotalTpt == 0 {
		return 0, errors.New("total cant be 0")
	}
	return workerTotalTpt, nil
}
