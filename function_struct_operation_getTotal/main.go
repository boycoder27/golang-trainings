package main

import "fmt"

type Numbers struct {
	values []int
}

func main() {
	numbers := Numbers{
		values: []int{2, 4, 6, 8, 10},
	}

	numbers.getTotal()
}

func (n Numbers) getTotal() {
	totalNum := 0

	for _, number := range n.values {
		totalNum += number
	}

	fmt.Println(totalNum)
}
