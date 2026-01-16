package main

import (
	"challenges/utils"
	"fmt"
)

func main() {
	greeting, err := utils.Greet("kevin")
	fmt.Println(greeting, err)

}
