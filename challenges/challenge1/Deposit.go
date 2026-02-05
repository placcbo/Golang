package challenge1

type Account struct {
	Owner   string
	Balance int
}

func (a Account) Deposit(amount int) Account {
	return Account{
		Owner:   a.Owner,
		Balance: a.Balance + amount,
	}

}
