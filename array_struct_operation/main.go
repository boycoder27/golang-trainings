package main

import "fmt"

type Transaction struct {
	Type   string
	Amount int
}

func main() {
	balance := 10000

	transactions := []Transaction{
		{
			Type:   "Debit",
			Amount: 2000,
		},
		{
			Type:   "Credit",
			Amount: 5000,
		},
	}

	fmt.Println("Balance is ", balance)
}
