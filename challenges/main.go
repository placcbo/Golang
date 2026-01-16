package main

import (
	"challenges/utils"
	"fmt"
)

func main() {
	greeting, err := utils.Greet("kevin")
	fmt.Println(greeting, err)

	result, err := utils.Divide(19, 0)

	fmt.Println(result, err)

	sum, product := utils.Stats(3, 4)
	fmt.Printf("sum: %v, product: %v\n", sum, product)

	value1 := Price("kamau")
	fmt.Println(value1)
}

func Price(name string) string {
	return "hello " + name

}
