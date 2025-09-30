package main

import "fmt"

func main() {
	firstNames := []string{"Kobe", "Lebron", "Stephen", "Jordan", "Klay"}
	lastNames := []string{"Bryant", "James", "Curry", "Clarkson", "Thompson"}

	for count := range firstNames {
		fmt.Println(firstNames[count] + " " + lastNames[count])
	}

}
