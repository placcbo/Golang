package utils

import "errors"

// score takes the result of tpt and accuracy and return a score

// score 1, tpt <55 || accuracy <85
// score 2, tpt >= 55 && tpt <= 67 && accuracy >=85 && accuracy <=90
// score 3, tpt >= 67 && tpt <= 100 && accuracy >=95 && accuracy <=100

func Score(Tpt int, Accuracy float64) (int, error) {

	if Accuracy < 0 || Accuracy > 100 {
		return 0, errors.New("Accuracy must be between 0 - 100")
	}

	if Tpt < 0 {
		return 0, errors.New("tpt cannot be negative")

	}

	var finalScore int

	if Tpt < 55 || Accuracy < 85 {
		finalScore = 1
	} else if (Tpt >= 55 && Tpt <= 67 && Accuracy >= 85) || (Tpt >= 55 && Accuracy >= 85 && Accuracy < 90) {
		finalScore = 2
	} else if Tpt >= 67 && Tpt <= 100 && Accuracy >= 95 && Accuracy <= 100 {
		finalScore = 3

	} else {
		finalScore = 1
	}

	return finalScore, nil

}
