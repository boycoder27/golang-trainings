package main

import "fmt"

type Numbers struct {
	values []int
}

func main() {
	numbers := Numbers{
		values: []int{2, 4, 6, 8, 10},
	}

	fmt.Printf("The total Number is %d\n", numbers.getTotal())
	fmt.Printf("The Average Number is %d\n", numbers.getAverage())
}

func (n Numbers) getTotal() int {
	var totalNum = 0

	for _, number := range n.values {
		totalNum += number
	}

	return totalNum

}

func (n Numbers) getAverage() int {
	var totalNum = 0
	var count = 0

	for _, number := range n.values {
		totalNum += number
		count++
	}

	average := totalNum / count

	return average

}
