package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(i)

}
