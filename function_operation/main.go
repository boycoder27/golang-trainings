package main

import "fmt"

func main() {
	num1 := -3
	result := checkNumber(num1)

	fmt.Println(result)
}

func checkNumber(num1 int) string {
	result := num1

	if result > 0 {
		return "Positive"
	} else {
		return "Negative"
	}
}
