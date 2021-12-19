package main

import "fmt"

type player struct {
	string
	int
}

func main() {
	player1 := player{"Muhammad Ali", 99}
	fmt.Println("Player 1:", player1)
	fmt.Printf("player1.int=%d player1.string=%s\n", player1.int, player1.string)

	player2 := player{
		int:    36,
		string: "Federer",
	}
	fmt.Println("Player 2:", player2)
}
