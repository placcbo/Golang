package main

import (
	"January/week1/utils"
	"fmt"
)

func main() {

	result, err := utils.SafeDivide(10, -1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("result:", result)
}
