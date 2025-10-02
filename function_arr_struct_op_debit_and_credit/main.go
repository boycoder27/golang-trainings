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
		{
			Type:   "Debit",
			Amount: 1000,
		},
		{
			Type:   "Credit",
			Amount: 2000,
		},
		{
			Type:   "Debit",
			Amount: 3000,
		},
		{
			Type:   "Credit",
			Amount: 4000,
		},
		{
			Type:   "Debit",
			Amount: 5000,
		},
		{
			Type:   "Credit",
			Amount: 1000,
		},
	}

	for _, transaction := range transactions {
		if transaction.Type == "Debit" {
			balance = debit(balance, transaction.Amount)
		} else if transaction.Type == "Credit" {
			balance = credit(balance, transaction.Amount)
		}
	}

	fmt.Println(balance)

}

//blueprint or Plan
func debit(balance int, Amount int) int {
	diff := balance - Amount

	return diff
}

//blueprint or Plan
func credit(balance int, amount int) int {
	sum := balance + amount

	return sum
}

// answer:
// 11000
