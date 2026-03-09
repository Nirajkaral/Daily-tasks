package main

import "fmt"

type BankAccount struct {
	name    string
	balance float64
}

// Method to deposit money
func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

// Method to check balance
func (b BankAccount) ShowBalance() {
	fmt.Println("Balance:", b.balance)
}

func main() {

	acc := BankAccount{name: "Neeraj", balance: 1000}

	acc.Deposit(500)

	acc.ShowBalance()
}