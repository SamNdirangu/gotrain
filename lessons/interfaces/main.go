package interfaces

import "fmt"

type IBankAccount interface {
	GetAccountName() string
	GetBalance() float64
	Deposit(amount float64)
	Withdraw(amount float64) error
}

func Run() {
	myAccounts := []IBankAccount{
		NewBitcoinAccount(),
		NewStanChartAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(20)
		account.Withdraw(10)
		accountName := account.GetAccountName()
		balance := account.GetBalance()
		fmt.Println(accountName, " account balance is ", balance)
	}
}
