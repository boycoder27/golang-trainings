package main

import "fmt"

func main() {
	input := 5
	num := 0

	// for num := range input {
	// 	fmt.Println(num)
	// }

	for num < input {
		fmt.Println(num)
		num++
	}

}

// answer:
// 0
// 1
// 2
// 3
// 4
