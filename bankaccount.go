package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance int
}

func (b *BankAccount) deposit(amount int) {
	b.Balance += amount
}

func main() {
	r := BankAccount{
		Owner:   "Thái",
		Balance: 1000,
	}

	r.deposit(500)
	fmt.Println(r.Owner, "has", r.Balance)
}
