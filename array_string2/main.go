package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f"}
	letterStorage := ""
	first := false

	for _, letter := range letters {
		if first {
			letterStorage = letter + letterStorage
			first = false
		} else {
			letterStorage = letterStorage + letter
			first = true
		}

	}

	fmt.Println(letterStorage)
}
