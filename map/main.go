package main

import "fmt"

func main() {
	// player1LastName:="James"
	// player1FirstName:= "Lebron"
	// player1Team := "Lakers"

	// player2LastName := "Jordan"
	// player2FirstName:="Michael"
	// Player2Team :="Bulls"

	player1 := map[string]any{}

	player1["lastName"] = "James"
	player1["firstName"] = "Lebron"
	player1["team"] = "Lakers"
	player1["age"] = 40

	player2 := map[string]any{}

	player2["lastName"] = "Bryant"
	player2["firstName"] = "Kobe"
	player2["team"] = "Lakers"
	player2["age"] = 45

	player3 := map[string]any{}

	player3["lastName"] = "Curry"
	player3["firstName"] = "Stephen"
	player3["team"] = "GSW"
	player3["age"] = 39

	player4 := map[string]any{}

	player4["lastName"] = "Harden"
	player4["firstName"] = "James"
	player4["team"] = "Clippers"
	player4["age"] = 44

	player5 := map[string]any{}

	player5["lastName"] = "Clarkson"
	player5["firstName"] = "Jordan"
	player5["team"] = "Utah"
	player5["age"] = 29

	msg := "Player %d is %s %s from %s\n"

	fmt.Printf(msg, 1, player1["firstName"], player1["lastName"], player1["team"])
	fmt.Printf(msg, 2, player2["firstName"], player2["lastName"], player2["team"])
	fmt.Printf(msg, 3, player3["firstName"], player3["lastName"], player3["team"])
	fmt.Printf(msg, 4, player4["firstName"], player4["lastName"], player4["team"])
	fmt.Printf(msg, 5, player5["firstName"], player5["lastName"], player5["team"])
}
