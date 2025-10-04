package main

import "fmt"

type Student struct {
	Name    string
	English int
	Math    int
	Science int
	Average int
}

func main() {

	students := []Student{
		{
			Name:    "Bea Alonzo",
			English: 89,
			Math:    91,
			Science: 89,
		},
		{
			Name:    "Lyza Soberano",
			English: 94,
			Math:    90,
			Science: 93,
		},
		{
			Name:    "Kim Chui",
			English: 88,
			Math:    91,
			Science: 87,
		},
	}

	for num, student := range students {
		grades := []int{student.English, student.Math, student.Science}
		ave := getAverage(grades)
		fmt.Println(student.Name, ave)
		students[num].Average = ave
	}

	studentAverage := []int{students[0].Average, students[1].Average, students[2].Average}
	result := getAverage(studentAverage)

	fmt.Println("Overall Average is", result)

}

// blueprint
func getAverage(items []int) int {
	var total int
	var totalNumber int

	for _, item := range items {
		total += item
		totalNumber++
	}

	result := total / totalNumber

	return result

}

// answer:
// Bea Alonzo 89
// Lyza Soberano 92
// Kim Chui 88
// Overall Average is 89
