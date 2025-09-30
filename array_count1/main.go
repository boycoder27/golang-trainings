package main

import "fmt"

func main() {
	names := []string{"Jose", "wally", "paolo"}

	for count, name := range names {
		fmt.Println(count, " - ", name)
	}
}

//answer:
// 0  -  Jose
// 1  -  wally
// 2  -  paolo

// Note: The count will show how many strings are in the array.
