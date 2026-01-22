package week2

func Grade(score int) string {
	if score < 0 || score > 100 {
		return "invalid"
	}

	if a := score; a >= 90 && a <= 100 {
		return "A"
	}

	if b := score; b >= 70 && b <= 89 {
		return "B"
	}

	if e := score; e >= 50 && e <= 69 {
		return "C"
	} else {
		return "F"
	}

}

// func Grade(score int) string
// Rules
// If score < 0 or score > 100 → return "invalid"

// 90–100 → "A"

// 70–89 → "B"

// 50–69 → "C"

// Below 50 → "F"

// Requirements

// Use if / else

// Use at least ONE short statement inside an if

// Do NOT use switch
