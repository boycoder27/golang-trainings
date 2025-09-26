package main

import "fmt"

type Menu struct {
	Name  string
	Price int
}

type Order struct {
	Name     string
	Quantity int
}

func main() {
	money := 1000

	menus := []Menu{
		{
			Name:  "Burger",
			Price: 50,
		},

		{
			Name:  "Spaghetti",
			Price: 100,
		},
		{
			Name:  "Coke",
			Price: 25,
		},
	}

	orders := []Order{
		{
			Name:     "Burger",
			Quantity: 3,
		},
		{
			Name:     "Coke",
			Quantity: 4,
		},
	}

	totalCost := 0
	for _, order := range orders {
		fmt.Printf("Item: %s, Quantity: %d\n", order.Name, order.Quantity)

		for _, menu := range menus {
			if order.Name == menu.Name {
				bill := order.Quantity * menu.Price
				fmt.Printf("Bill: %d\n", bill)

				totalCost += bill

			}
		}
	}
	fmt.Printf("Total Cost : %d\n", totalCost)

	fmt.Printf("Bill Amount : %d\n", money)

	change := money - totalCost
	fmt.Printf("Change : %d\n", change)

}
