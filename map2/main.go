package main

import "fmt"

func main() {

	student1 := map[string]any{}

	student1["lastName"] = "Manalo"
	student1["firstName"] = "Jose"
	student1["age"] = 20
	student1["Grade level"] = 12
	student1["Section"] = "Mars"

	student2 := map[string]any{}

	student2["lastName"] = "Bayola"
	student2["firstName"] = "Wally"
	student2["age"] = 18
	student2["Grade level"] = 11
	student2["Section"] = "Pluto"

	student3 := map[string]any{}

	student3["lastName"] = "Ballesteros"
	student3["firstName"] = "Paolo"
	student3["age"] = 15
	student3["Grade level"] = 9
	student3["Section"] = "Venus"

	student4 := map[string]any{}

	student4["lastName"] = "Mendoza"
	student4["firstName"] = "Maine"
	student4["age"] = 13
	student4["Grade level"] = 6
	student4["Section"] = "Jupiter"

	msg := "Student %d is %s %s at age of %d he/she is a Grade %d student from %s section.\n"

	fmt.Printf(msg, 1, student1["firstName"], student1["lastName"], student1["age"], student1["Grade level"], student1["Section"])
	fmt.Printf(msg, 2, student2["firstName"], student2["lastName"], student2["age"], student2["Grade level"], student2["Section"])
	fmt.Printf(msg, 3, student3["firstName"], student3["lastName"], student3["age"], student3["Grade level"], student3["Section"])
	fmt.Printf(msg, 4, student4["firstName"], student4["lastName"], student4["age"], student4["Grade level"], student4["Section"])

}
