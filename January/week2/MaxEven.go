package week2

import "errors"

func MaxEven(nums []int) (int, error) {

	var largest int
	found := false

	for _, num := range nums {
		if num%2 == 0 {
			if !found || num > largest {
				largest = num
				found = true
			}
		}
	}

	if !found {
		return 0, errors.New("no even numbers found")
	}
	return largest, nil

}

// Challenge

// Write a function:

// func MaxEven(nums []int) (int, error)

// Rules

// Return the largest even number in the slice.

// If the slice is empty, or has no even numbers, return 0 and an error "no even numbers found".

// You must use if/else and can optionally use a short statement in one of the ifs.

// Examples:

// MaxEven([]int{1, 3, 4, 2}) // returns 4, nil
// MaxEven([]int{1, 3, 5})    // returns 0, error
// MaxEven([]int{})            // returns 0, error
