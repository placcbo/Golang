package main

import (
	"January/week3"
	"fmt"
)

func main() {
	r := week3.Wallet{
		Owner:   "kevin ndirangu",
		Balance: 23000,
	}
	result := r.Deposit(43000)
	fmt.Println(r)
	fmt.Println(result)
}
