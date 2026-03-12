package main

import "fmt"

type bankAccount struct {
	name    string
	balance float64
}


func (b *bankAccount) deposit(amount float64) {
	b.balance += amount
}


func (b bankAccount) showBalance() {
	fmt.Println("Balance:", b.balance)
}

func encp() {

	acc := bankAccount{name: "Niraj", balance: 1000}

	acc.deposit(500)

	acc.showBalance()
}