package main

import "fmt"

func main() {
	input := 5

	// for num := range input {
	// 	fmt.Println(num)
	// }

	num := 0
	for num < input {
		fmt.Println(num)
		num++
	}

}
