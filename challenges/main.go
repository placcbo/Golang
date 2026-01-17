package main

import (
	"challenges/utils"
	"fmt"
)

func main() {
	result, err := utils.SafeDivide(10, 4)
	fmt.Println(float64(result), err)

}
