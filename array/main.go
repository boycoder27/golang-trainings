package main

import "fmt"

func main() {
	names := []string{
		"James",
		"Kobe",
		"Pedro",
	}

	for _, name := range names {
		fmt.Println("Name is ", name)
	}
}
