package main

import "fmt"

type Caculator struct {
	num1 int
	num2 int
}

func main() {
	value := Caculator{
		num1: 10,
		num2: 5,
	}

	sum := value.Add()
	diff := value.Substract()
	times := value.Multiply()
	quotient := value.Division()

	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(times)
	fmt.Println(quotient)

}

func (c Caculator) Add() int {
	sum := c.num1 + c.num2
	return sum
}

func (c Caculator) Substract() int {
	diff := c.num1 - c.num2
	return diff
}

func (c Caculator) Multiply() int {
	times := c.num1 * c.num2
	return times

}

func (c Caculator) Division() int {
	quotient := c.num1 / c.num2
	return quotient

}
