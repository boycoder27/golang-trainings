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
	sum := addition.getCompute()
	fmt.Println("Sum is:", sum)

	substract := NewMath(25, 5, "Sub")
	diff := substract.getCompute()
	fmt.Println("Diff is:", diff)

	multiplication := NewMath(3, 3, "Multiply")
	product := multiplication.getCompute()
	fmt.Println("Multiply is:", product)

	quotient := NewMath(100, 10, "Division")
	division := quotient.getCompute()
	fmt.Println("Division is:", division)
}

func (m Math) getCompute() int {
	var answer int
	if m.Operation == "Add" {
		answer = m.FirstNumber + m.SecondNumber
	} else if m.Operation == "Sub" {
		answer = m.FirstNumber - m.SecondNumber
	} else if m.Operation == "Multiply" {
		answer = m.FirstNumber * m.SecondNumber
	} else if m.Operation == "Division" {
		answer = m.FirstNumber / m.SecondNumber
	}
	return answer
}

// answer:
// Sum is: 15
// Diff is: 20
// Multiply is: 9
// Division is: 10
