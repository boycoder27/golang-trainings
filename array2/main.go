package main

import "fmt"

func main() {
	grades := []int{85, 74, 72, 88, 90}

	for _, grade := range grades {

		if grade >= 75 {
			fmt.Printf("The Grade is %d Passed\n", grade)
		} else {
			fmt.Printf("The Grade is %d Failed\n", grade)
		}
	}

}
