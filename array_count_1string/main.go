package main

import "fmt"

func main() {
	names := []string{"Jose", "wally", "paolo"}

	for count, name := range names {
		fmt.Println(count, " - ", name)
	}
}
