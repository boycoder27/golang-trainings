package main

import "fmt"

type Math struct {
	FirstNumber  int
	SecondNumber int
	Operation    string
}

func NewMath(firstNumber int, secondNumber int, operation string) Math {
	math := Math{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
		Operation:    operation,
	}
	return math
}

func main() {
	addition := NewMath(10, 5, "Add")
	sum := addition.Compute()
	fmt.Println("Sum is:", sum)
}

func (m Math) Compute() int {
	var answer int
	if m.Operation == "Add" {
		answer = m.FirstNumber + m.SecondNumber
	}

	return answer
}
