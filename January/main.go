package main

import (
	"fmt"

	week3 "January/week3"
)

func main() {
	// Create two wallets
	alice := week3.NewWallet("Alice", 100)
	bob := week3.NewWallet("Bob", 50)

	fmt.Println("Original wallets:")
	fmt.Println("Alice:", alice.Balance) // 100
	fmt.Println("Bob  :", bob.Balance)   // 50
	fmt.Println("-----")

	// Successful transfer
	alice2, bob2, ok := alice.Transfer(bob, 30)
	fmt.Println("After transferring $30 from Alice to Bob:")
	fmt.Println("Alice:", alice2.Balance) // 70
	fmt.Println("Bob  :", bob2.Balance)   // 80
	fmt.Println("Success?", ok)           // true
	fmt.Println("Original wallets remain unchanged:")
	fmt.Println("Alice:", alice.Balance) // 100
	fmt.Println("Bob  :", bob.Balance)   // 50
	fmt.Println("-----")

	// Failed transfer
	alice3, bob3, ok := alice.Transfer(bob, 200)
	fmt.Println("Attempting to transfer $200 from Alice to Bob:")
	fmt.Println("Alice:", alice3.Balance) // 100
	fmt.Println("Bob  :", bob3.Balance)   // 50
	fmt.Println("Success?", ok)           // false
	fmt.Println("-----")

	// Print history for Alice2 and Bob2
	fmt.Println("Alice2 History:")
	for _, h := range alice2.History {
		fmt.Println(h)
	}
	fmt.Println("Bob2 History:")
	for _, h := range bob2.History {
		fmt.Println(h)
	}
}
