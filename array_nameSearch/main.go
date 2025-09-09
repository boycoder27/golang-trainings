package main

import "fmt"

func main() {
	names := []string{"Jose", "Vic", "Wally", "Joey", "Maine"}
	search := "boy"
	searchFound := false

	for _, name := range names {
		if name == search {
			searchFound = true
		}
	}

	if searchFound {
		fmt.Printf("%s was found.\n", search)
	} else {
		fmt.Printf("%s was not found.\n", search)
	}
}
