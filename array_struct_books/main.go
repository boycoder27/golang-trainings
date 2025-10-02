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
		fmt.Printf("The book name is %s, The author is %s , The Year Published is %d\n", book.Name, book.Author, book.YearPublished)
	}
}

// answer:
// The book name is Abc, The author is Pedro , The Year Published is 1995
// The book name is Xyz, The author is Juan , The Year Published is 2000
