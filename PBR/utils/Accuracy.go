package utils

import "errors"

func WorkerAccuracyAverage(a, b, c int) (float64, error) {
	totalAccuracy := float64(a + b + c)
	if totalAccuracy == 0 {
		return 0, errors.New("total cannot be 0")
	}
	averageAccuracy := totalAccuracy / 3
	return averageAccuracy, nil

}
