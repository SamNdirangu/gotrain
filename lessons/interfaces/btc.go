package interfaces

import "errors"

type BitcoinAccount struct {
	accountName string
	balance     float64
	fee         float64
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		accountName: "BTC",
		balance:     0,
		fee:         3,
	}
}

func (btc *BitcoinAccount) GetAccountName() string {
	return btc.accountName
}

func (btc *BitcoinAccount) GetBalance() float64 {
	return btc.balance
}

func (btc *BitcoinAccount) Deposit(amount float64) {
	btc.balance += amount
}

func (btc *BitcoinAccount) Withdraw(amount float64) error {
	newBalance := btc.balance - amount - btc.fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	btc.balance = newBalance
	return nil
}
