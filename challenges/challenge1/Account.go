package challenge1

type Account struct {
	Owner   string
	Balance int
}

// deposit

func (a *Account) Deposit(amount int) {

	a.Balance += amount
}
func (a *Account) Withdraw(amount int) {
	if a.Balance >= 0 && a.Balance > amount {
		return
	}

	a.Balance -= amount

}

// validate methods

func (a Account) canWithdraw(amount int) bool {
	return a.Balance > amount
}

func (a *Account) PayRent(amount int) {
	a.Withdraw(amount)
}

func (a *Account) PaySubscription(amount int) {
	a.Withdraw(amount)
}
