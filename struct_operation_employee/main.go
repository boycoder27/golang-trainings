package main

import "fmt"

type employee struct {
	Name     string
	Position string
	Salary   int
}

func main() {
	employee1 := employee{
		Name:     "Pedro",
		Position: "HR",
		Salary:   20000,
	}

	employee2 := employee{
		Name:     "Juan",
		Position: "IT Manager",
		Salary:   55000,
	}

	employee3 := employee{
		Name:     "Hanzel",
		Position: "Marketing",
		Salary:   35000,
	}

	fmt.Printf("The Employee name is %s, The position is %s and the salary is %d\n", employee1.Name, employee1.Position, employee1.Salary)
	fmt.Printf("The Employee name is %s, The position is %s and the salary is %d\n", employee2.Name, employee2.Position, employee2.Salary)
	fmt.Printf("The Employee name is %s, The position is %s and the salary is %d\n", employee3.Name, employee3.Position, employee3.Salary)

	totalSalary := employee1.Salary + employee2.Salary + employee3.Salary
	fmt.Printf("The Total Salary is %d\n", totalSalary)
}

// answer:
// The Employee name is Pedro, The position is HR and the salary is 20000
// The Employee name is Juan, The position is IT Manager and the salary is 55000
// The Employee name is Hanzel, The position is Marketing and the salary is 35000
// The Total Salary is 110000
