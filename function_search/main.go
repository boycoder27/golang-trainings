package main

import "fmt"

func main() {
	names := []string{"Josh", "Carl", "James", "Bryan"}
	search := "Josh"

	result := searchName(search, names)
	fmt.Println(result)
}

func searchName(search string, names []string) bool {
	searchFound := false

	for _, name := range names {
		if name == search {
			searchFound = true
		}
	}

	return searchFound
}
