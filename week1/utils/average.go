package utils

func Average(a, b string) (int, error) {
	x, err := ParseAgeStrict(a)

	if err != nil {
		return 0, err
	}

	y, err := ParseAgeStrict(b)
	if err != nil {
		return 0, err
	}

	return (x + y) / 2, nil

}
