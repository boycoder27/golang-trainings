package main

import "fmt"

func main() {
	firstName := "Juan"
	lastName := "Dela Cruz"

	fullName := formatName(firstName, lastName)

	fmt.Println(fullName)
}

func formatName(f string, l string) string {
	full := f + " " + l
	return full
}

// answer:
// Juan Dela Cruz
