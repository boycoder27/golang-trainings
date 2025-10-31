package main

import "fmt"

func main() {
	names := []string{"Josh", "Carl", "James", "Bryan"}
	search := "Josh"

	result := getSearchName(search, names)
	fmt.Println(result)
}

func getSearchName(search string, names []string) bool {
	searchFound := false

	for _, name := range names {
		if name == search {
			searchFound = true
		}
	}

	return searchFound
}

// answer:
// true
