package main

import "fmt"

func main() {
	age := 18

	if age <= 3 {
		fmt.Println("Baby")
	} else if age < 6 {
		fmt.Println("Toddler")
	} else if age < 12 {
		fmt.Println("Child")
	} else if age < 17 {
		fmt.Println("Teen")
	} else {
		fmt.Println("Adult")
	}
}

// answer:
// Adult
