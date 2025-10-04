package main

import "fmt"

func main() {
	names := []string{
		"James",
		"Kobe",
		"Pedro",
	}

	for _, name := range names {
		fmt.Println("Name is", name)
	}
}

// answer:
// Name is James
// Name is Kobe
// Name is Pedro
