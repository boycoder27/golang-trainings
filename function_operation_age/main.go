package main

import "fmt"

func main() {
	age := 18

	result := getAge(age)

	fmt.Println(result)
}

func getAge(age int) string {
	ageResult := age

	if ageResult < 3 {
		return "Baby"
	} else if ageResult < 8 {
		return "Toddler"
	} else if ageResult < 12 {
		return "Child"
	} else if ageResult < 17 {
		return "Teen"
	} else {
		return "Adult"
	}

}
