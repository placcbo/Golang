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

func (w Wallet) Withdraw(amount int) (Wallet, bool) {

	// Cannot withdraw, return original wallet
	if amount > w.Balance {
		return w, false
	}
	// Withdrawal possible, return new wallet
	return Wallet{
		Owner:   "kevin ndirangu",
		Balance: w.Balance - amount,
	}, true

}
