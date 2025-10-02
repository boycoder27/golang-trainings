package main

import "fmt"

type Student struct {
	Name    string
	English int
	Math    int
	Mapeh   int
}

func main() {
	students := []Student{
		{
			Name:    "Chris",
			English: 84,
			Math:    79,
			Mapeh:   78},
		{
			Name:    "Bronny",
			English: 75,
			Math:    74,
			Mapeh:   76,
		},
	}
	for _, student := range students {

		total := student.English + student.Math + student.Mapeh
		average := total / 3

		if average <= 75 {
			fmt.Printf("Hi %s your average is %d. You Failed!\n", student.Name, average)
		} else {
			fmt.Printf("Hi %s your average is %d. You Passed!\n", student.Name, average)

		}

	}
}

// answer:
// Hi Chris your average is 80. You Passed!
// Hi Bronny your average is 75. You Failed!
