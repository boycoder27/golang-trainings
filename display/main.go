package main

import "fmt"

func main() {
	name := "Juan"
	age := 25
	address := "Quezon City"

	msg := "My name is %s, I am %d and I live in %s\n"

	fmt.Printf(msg, name, age, address)
}
