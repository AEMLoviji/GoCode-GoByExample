package main

import "fmt"

// Assignment: Using 'map of maps' implement the following functionality:
// 1. addPlayer("firstName", "lastName")
//    to add a new player to our map
//    'first name' must be unique (for simplicity and onlyt in this program)
// 2. hasPlayer("firstName", "lastName") - returns bool
//    to check if a pair of firstName/lastName exist in our map.

var players = make(map[string]map[string]bool)

func main() {
	addPlayer("Leo", "Messi")
	addPlayer("Roger", "Federer")
	addPlayer("Michael", "Jordan")

	fmt.Println()
	fmt.Println(hasPlayer("Roger", "Federer"))
	fmt.Println(hasPlayer("Michael", "Jordan"))
	fmt.Println(hasPlayer("Michael", "Messi"))
	fmt.Println(hasPlayer("Leo", "Jordan"))

	fmt.Println()
	fmt.Println(players)

	fmt.Println()
	for k1, v1 := range players {
		for k2, v2 := range v1 {
			fmt.Println(k1, k2, v2)
		}
	}
}

func addPlayer(fName, lName string) {
	n := players[fName]
	if n == nil {
		n = make(map[string]bool)
		players[fName] = n
	}
	n[lName] = true
}

func hasPlayer(fName, lName string) bool {
	return players[fName][lName]
}
