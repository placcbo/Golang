package week3

import "fmt"

type Wallet struct {
	Owner   string
	Balance int
	History []string
}

func (w Wallet) Deposit(amount int) Wallet {
	return Wallet{
		Owner:   w.Owner,
		Balance: w.Balance + amount,
		History: append(w.History, fmt.Sprintf("Deposited $%0d", amount)),
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
		History: append(w.History, fmt.Sprintf("Withdrew $%d", amount)),
	}, true

}

func NewWallet(owner string, balance int) Wallet {
	if balance < 0 {
		balance = 0
	}

	return Wallet{
		Owner:   owner,
		Balance: balance,
		History: []string{},
	}
}

func (w Wallet) Transfer(to Wallet, amount int) (Wallet, Wallet, bool) {
	newSender, ok := w.Withdraw(amount)
	if !ok {
		return w, to, false
	}
	newReceiver := to.Deposit(amount)

	// Add transfer history
	newSender.History = append(newSender.History, fmt.Sprintf("Transferred $%d to %s", amount, to.Owner))
	newReceiver.History = append(newReceiver.History, fmt.Sprintf("Received $%d from %s", amount, w.Owner))

	return newSender, newReceiver, true
}
