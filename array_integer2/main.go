package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	totalNum := 0

	for _, number := range numbers {
		totalNum += number
	}

	fmt.Printf("The Total numbers is %d\n", totalNum)
}
