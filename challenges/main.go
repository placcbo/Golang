package main

import (
	"fmt"
	"strconv"
)

func ParseInt(input string) (int, error) {
	result, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	result, err := ParseInt("100H")
	fmt.Println(result, err)
}
