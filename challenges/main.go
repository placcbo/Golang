package main

import (
	"challenges/utils"
	"fmt"
)

func main() {
	result, err := utils.ParseAge("abc")
	fmt.Println(result, err)

}
