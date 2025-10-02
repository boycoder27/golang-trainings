package main

import "fmt"

func main() {
	limit := 3.0
	var counter float64
	for counter <= limit {
		//4 > 3.14
		if counter+1 > limit {
			decimalValue := limit - counter
			if decimalValue > 0 {
				fmt.Println("With decimal Value")
			} else {
				fmt.Println("No decimal Value to display.")
			}
		}

		counter++

	}

}
