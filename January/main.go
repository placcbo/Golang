package main

import (
	week3 "January/week3"
	"fmt"
)

func main() {
	r := week3.Wallet{
		Owner:   "kevin ndirangu",
		Balance: 23000,
	}
	result, err := r.Withdraw(500)
	fmt.Println(r)
	fmt.Println(result, err)
}
