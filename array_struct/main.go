package main

import "fmt"

type Book struct {
	Name          string
	Author        string
	YearPublished int
}

func main() {
	books := []Book{
		{
			Name:          "Abc",
			Author:        "Pedro",
			YearPublished: 1995,
		},
		{
			Name:          "Xyz",
			Author:        "Juan",
			YearPublished: 2000,
		},
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
