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
		Owner:   w.Owner,
		Balance: w.Balance - amount,
	}, true

}

func NewWallet(owner string, balance int) Wallet {
	if balance < 0 {
		balance = 0
	}

	return Wallet{
		Owner:   owner,
		Balance: balance,
	}
}

func (w Wallet) Transfer(to Wallet, amount int) (Wallet, Wallet, bool) {
	// Try to take money from sender (w)
	newSender, ok := w.Withdraw(amount)
	if !ok {
		// Not enough money → return originals
		return w, to, false
	}

	// Deposit into receiver (to)
	newReceiver := to.Deposit(amount)

	// Return updated wallets + success
	return newSender, newReceiver, true
}
