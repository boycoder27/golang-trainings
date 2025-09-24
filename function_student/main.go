package main

import "fmt"

type student struct {
	Name    string
	Math    int
	English int
	Science int
}

func main() {
	student1 := student{
		Name:    "Juan",
		Math:    93,
		English: 90,
		Science: 89,
	}
	Average1 := getAverageGrade(student1.Math, student1.English, student1.Science)
	fmt.Println(Average1)

	student2 := student{
		Name:    "Pedro",
		Math:    87,
		English: 86,
		Science: 88,
	}
	Average2 := getAverageGrade(student2.Math, student2.English, student2.Science)
	fmt.Println(Average2)

	student3 := student{
		Name:    "Mars",
		Math:    88,
		English: 90,
		Science: 79,
	}
	Average3 := getAverageGrade(student3.Math, student3.English, student3.Science)

	fmt.Println(Average3)

}

func getAverageGrade(Math int, English int, Science int) int {
	total := (Math + English + Science) / 3

	return total
}
