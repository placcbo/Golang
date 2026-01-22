package week2

import "errors"

func Grade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", errors.New("invalid score")
	}
	if a := score; a >= 90 && a <= 100 {
		return "A", nil
	}
	if score >= 70 && score <= 89 {
		return "B", nil
	}

	if score >= 50 && score <= 69 {
		return "C", nil
	}

	return "F", nil

}

// func Grade(score int) (string, error)
// Rules
// If score < 0 or score > 100 → return "" and an error

// Otherwise:

// 90–100 → "A"

// 70–89 → "B"

// 50–69 → "C"

// < 50 → "F"

// Hard requirement
// You must use an if with a short statement at least once.

// Example usage (for your own thinking):

// grade, err := Grade(85)
// // grade = "B", err = nil
