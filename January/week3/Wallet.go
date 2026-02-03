package week3

type Wallet struct {
	Owner   string
	Balance int
}

func (w Wallet) Deposit(amount int) Wallet {
	return Wallet{
		Owner:   w.Owner,
		Balance: w.Balance + amount,
	}
}

// You’re building a tiny Wallet system.

// Requirements

// Define a struct called Wallet with:

// Owner (string)

// Balance (int)

// Add a method called Deposit that:

// Takes an amount (int)

// Returns a new Wallet

// Increases the balance by amount

// ❗️Do not modify the original Wallet
