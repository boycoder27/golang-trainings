package main

import "fmt"

func main() {
	firstName := "Juan"
	lastName := "Dela Cruz"

	fullName := getFormatName(firstName, lastName)

	fmt.Println(fullName)
}

func getFormatName(f string, l string) string {
	full := f + " " + l
	return full
}

// answer:
// Juan Dela Cruz
