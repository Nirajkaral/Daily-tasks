package main

import "fmt"

type BankAccount struct {
    Owner        string
    balance      float64  // Private
    transactions []string // Private
}


func NewBankAccount(owner string, balance float64) *BankAccount {
    return &BankAccount{Owner: owner, balance: balance}
}


func (b *BankAccount) deposit(amount float64) {
    if amount <= 0 {
        fmt.Println("Deposit amount must be positive!")
        return
    }
    b.balance += amount
    b.transactions = append(b.transactions, fmt.Sprintf("Deposited: +%.2f", amount))
    fmt.Printf(" Deposited %.2f. Balance: %.2f\n", amount, b.balance)
}


func (b *BankAccount) Withdraw(amount float64) {
    if amount > b.balance {
        fmt.Println(" Insufficient funds!")
        return
    }
    b.balance -= amount
    b.transactions = append(b.transactions, fmt.Sprintf("Withdrew: -%.2f", amount))
    fmt.Printf(" Withdrew %.2f. Balance: %.2f\n", amount, b.balance)
}


func (b BankAccount) GetBalance() float64 {
    return b.balance
}

func (b BankAccount) GetTransactions() {
    fmt.Println("\n Transaction History:")
    for _, t := range b.transactions {
        fmt.Println(" ", t)
    }
}

func structpointer() {
    acc := NewBankAccount("Ravi", 5000)
    acc.deposit(2000)
    acc.Withdraw(1000)
    acc.Withdraw(9000)   
    fmt.Println("Balance:", acc.GetBalance())
    acc.GetTransactions()
}