package main

import "fmt"

type book struct {
	book         string
	author       string
	yearReleased int
	borrowed     bool
}

func main() {
	book1 := book{
		book:         "Noli Me Tangere",
		author:       "Dr.Jose Rizal",
		yearReleased: 1887,
		borrowed:     true,
	}

	book2 := book{
		book:         "Cinderella",
		author:       "Charles Perrault",
		yearReleased: 1697,
		borrowed:     false,
	}

	book3 := book{
		book:         "Superman",
		author:       "James Gunn",
		yearReleased: 2026,
		borrowed:     false,
	}

	fmt.Printf("Book of %s, The author is %s, The Year Released is %d, is this book borrowed? %t\n", book1.book, book1.author, book1.yearReleased, book1.borrowed)
	fmt.Printf("Book of %s, The author is %s, The Year Released is %d, is this book borrowed? %t\n", book2.book, book2.author, book2.yearReleased, book2.borrowed)
	fmt.Printf("Book of %s, The author is %s, The Year Released is %d, is this book borrowed? %t\n", book3.book, book3.author, book3.yearReleased, book3.borrowed)

}

// answer:
// Book of Noli Me Tangere, The author is Dr.Jose Rizal, The Year Released is 1887, is this book borrowed? true
// Book of Cinderella, The author is Charles Perrault, The Year Released is 1697, is this book borrowed? false
// Book of Superman, The author is James Gunn, The Year Released is 2026, is this book borrowed? false
