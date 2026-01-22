package week2

func SumEven(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum = sum + num
		}
	}

	return sum
}
