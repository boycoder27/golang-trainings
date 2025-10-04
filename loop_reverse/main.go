package main

import (
	"fmt"
)

func main() {
	input := 5

	num := 0

	for input >= num {
		fmt.Println(input)
		input--
	}
}

// answer:
// 5
// 4
// 3
// 2
// 1
// 0
