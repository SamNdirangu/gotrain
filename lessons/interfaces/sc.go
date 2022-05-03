package interfaces

import "errors"

type StanChartAccount struct {
	accountName string
	balance     float64
}

func NewStanChartAccount() *StanChartAccount {
	return &StanChartAccount{
		balance:     0,
		accountName: "StanChart",
	}
}
func (sc *StanChartAccount) GetAccountName() string {
	return sc.accountName
}

func (sc *StanChartAccount) GetBalance() float64 {
	return sc.balance
}

func (sc *StanChartAccount) Deposit(amount float64) {
	sc.balance += amount
}

func (sc *StanChartAccount) Withdraw(amount float64) error {
	if amount > sc.balance {
		return errors.New("insufficient funds")
	}
	sc.balance -= amount
	return nil
}
