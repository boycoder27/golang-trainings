package main

import "fmt"

func main() {
	input := 6
	num := 0

	for num <= input {

		if num+1 > input {
			quotient := float64(num) / 2

			counter := 0
			for float64(counter) <= quotient {
				if counter+1 > int(quotient) {

					identifier := quotient - float64(counter)
					if identifier > 0 {
						fmt.Println("ODD")
					} else {
						fmt.Println("EVEN")
					}

				}
				counter++
			}

		}
		num++

	}
}
