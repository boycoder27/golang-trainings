package main

import "fmt"

type Customer struct {
	Name   string
	Gender string
}

func main() {

	customers := []Customer{
		{
			Name:   "Marie",
			Gender: "Female",
		},
		{
			Name:   "Richard",
			Gender: "Male",
		},
		{
			Name:   "Dingdong",
			Gender: "Male",
		},
		{
			Name:   "Rose",
			Gender: "Female",
		},
		{
			Name:   "Carla",
			Gender: "Female",
		},
	}
	for _, customer := range customers {
		if customer.Gender == "Male" {
			fmt.Printf("Hi Sir %s welcome\n", customer.Name)
		} else {
			fmt.Printf("Hi Maam %s welcome\n", customer.Name)
		}

	}
}

// answer:
// Hi Maam Marie welcome
// Hi Sir Richard welcome
// Hi Sir Dingdong welcome
// Hi Maam Rose welcome
// Hi Maam Carla welcome
